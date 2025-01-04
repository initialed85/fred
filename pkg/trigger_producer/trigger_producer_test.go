package trigger_producer

import (
	"context"
	"fmt"
	"testing"
	"time"

	_config "github.com/initialed85/djangolang/pkg/config"
	"github.com/initialed85/djangolang/pkg/helpers"
	"github.com/initialed85/djangolang/pkg/query"
	"github.com/initialed85/fred/internal"
	"github.com/initialed85/fred/pkg/api"
	"github.com/initialed85/fred/pkg/change_producer"
	"github.com/stretchr/testify/require"
)

func TestTriggerProducer(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := _config.GetDBFromEnvironment(ctx)
	require.NoError(t, err)
	defer func() {
		db.Close()
	}()

	//
	// given repository
	//

	repository := &api.Repository{
		URL: "https://github.com/initialed85/djangolang",
	}

	tx, err := db.Begin(ctx)
	require.NoError(t, err)
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	err = repository.Insert(ctx, tx, false, false)
	require.NoError(t, err)
	defer func() {
		tx, err := db.Begin(ctx)
		require.NoError(t, err)
		defer func() {
			_ = tx.Rollback(ctx)
		}()

		_ = repository.Delete(ctx, tx)
		_ = tx.Commit(ctx)
	}()

	//
	// given rule
	//

	rule := &api.Rule{
		BranchName:   helpers.Ptr("main"),
		RepositoryID: repository.ID,
	}

	err = rule.Insert(ctx, tx, false, false)
	require.NoError(t, err)
	defer func() {
		tx, err := db.Begin(ctx)
		require.NoError(t, err)
		defer func() {
			_ = tx.Rollback(ctx)
		}()

		_ = rule.Delete(ctx, tx)
		_ = tx.Commit(ctx)
	}()

	//
	// given job
	//

	job := &api.Job{
		Name: "lint-and-build",
	}

	err = job.Insert(ctx, tx, false, false)
	require.NoError(t, err)
	defer func() {
		tx, err := db.Begin(ctx)
		require.NoError(t, err)
		defer func() {
			_ = tx.Rollback(ctx)
		}()

		_ = job.Delete(ctx, tx)
		_ = tx.Commit(ctx)
	}()

	//
	// given tasks
	//

	task1 := &api.Task{
		Name:     "lint",
		Index:    0,
		Platform: "linux/amd64",
		Image:    "initialed85/the-last-ci-image-you-will-ever-need:latest",
		Script:   "go vet ./...",
		JobID:    job.ID,
	}

	err = task1.Insert(ctx, tx, false, false)
	require.NoError(t, err)
	defer func() {
		tx, err := db.Begin(ctx)
		require.NoError(t, err)
		defer func() {
			_ = tx.Rollback(ctx)
		}()

		_ = task1.Delete(ctx, tx)
		_ = tx.Commit(ctx)
	}()

	task2 := &api.Task{
		Name:     "build",
		Index:    0,
		Platform: "linux/amd64",
		Image:    "initialed85/the-last-ci-image-you-will-ever-need:latest",
		Script:   "go build -o bin/cmd ./cmd/",
		JobID:    job.ID,
	}

	err = task2.Insert(ctx, tx, false, false)
	require.NoError(t, err)
	defer func() {
		tx, err := db.Begin(ctx)
		require.NoError(t, err)
		defer func() {
			_ = tx.Rollback(ctx)
		}()

		_ = task2.Delete(ctx, tx)
		_ = tx.Commit(ctx)
	}()

	//
	// given trigger
	//

	trigger := &api.Trigger{
		RuleID: rule.ID,
		JobID:  job.ID,
	}

	err = trigger.Insert(ctx, tx, false, false)
	require.NoError(t, err)
	defer func() {
		tx, err := db.Begin(ctx)
		require.NoError(t, err)
		defer func() {
			_ = tx.Rollback(ctx)
		}()

		_ = trigger.Delete(ctx, tx)
		_ = tx.Commit(ctx)
	}()

	// when

	repositoriesPath := t.TempDir()
	err = change_producer.HandleRepository(ctx, tx, repository, repositoriesPath)
	require.NoError(t, err)

	changes, _, _, _, _, err := api.SelectChanges(
		ctx,
		tx,
		fmt.Sprintf(
			"%s = $$?? AND %s = $$??",
			api.ChangeTableRepositoryIDColumn,
			api.ChangeTableBranchNameColumn,
		),
		helpers.Ptr(fmt.Sprintf("%s DESC", api.ChangeTableCreatedAtColumn)),
		helpers.Ptr(1),
		nil,
		repository.ID,
		"main",
	)
	require.NoError(t, err)
	require.NotEmpty(t, changes)

	change := changes[0]
	err = change.Reload(query.WithLoad(ctx, api.RepositoryTable), tx)
	require.NoError(t, err)

	err = HandleChange(ctx, tx, change)
	require.NoError(t, err)

	// then

	err = change.Reload(ctx, tx)
	require.NoError(t, err)

	require.NotNil(t, change.TriggersProducedAt)
	require.Greater(t, *change.TriggersProducedAt, time.Time{})

	selectCtx := ctx
	selectCtx = query.WithLoad(selectCtx, api.JobTable)
	selectCtx = query.WithLoad(selectCtx, fmt.Sprintf("referenced_by_%s", api.OutputTable))

	executions, _, _, _, _, err := api.SelectExecutions(
		selectCtx,
		tx,
		fmt.Sprintf(
			"%s = $$??",
			api.ExecutionTableChangeIDColumn,
		),
		helpers.Ptr(fmt.Sprintf("%s DESC", api.ExecutionTableCreatedAtColumn)),
		helpers.Ptr(1),
		nil,
		change.ID,
	)
	require.NoError(t, err)
	require.NotEmpty(t, executions)

	execution := executions[0]
	require.Equal(t, internal.ExecutionOrTaskStatusPending, execution.Status)
	require.NotEmpty(t, execution.ReferencedByOutputExecutionIDObjects)
	require.Len(t, execution.ReferencedByOutputExecutionIDObjects, 2)

	output1 := execution.ReferencedByOutputExecutionIDObjects[0]
	require.Equal(t, internal.ExecutionOrTaskStatusPending, output1.Status)

	output2 := execution.ReferencedByOutputExecutionIDObjects[0]
	require.Equal(t, internal.ExecutionOrTaskStatusPending, output2.Status)

	err = tx.Commit(ctx)
	require.NoError(t, err)
}
