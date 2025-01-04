package job_executor

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/docker/docker/client"
	_config "github.com/initialed85/djangolang/pkg/config"
	"github.com/initialed85/djangolang/pkg/helpers"
	"github.com/initialed85/djangolang/pkg/query"
	"github.com/initialed85/fred/pkg/api"
	"github.com/initialed85/fred/pkg/change_producer"
	"github.com/initialed85/fred/pkg/trigger_producer"
	"github.com/stretchr/testify/require"
)

func TestJobExecutor(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := _config.GetDBFromEnvironment(ctx)
	require.NoError(t, err)
	defer func() {
		db.Close()
	}()

	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	require.NoError(t, err)

	defer func() {
		_ = apiClient.Close()
	}()

	apiClient.NegotiateAPIVersion(ctx)

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

	err = trigger_producer.HandleChange(ctx, tx, change)
	require.NoError(t, err)

	execution, err := api.JobExecutorClaimExecution(
		getExecutionCtx(ctx),
		tx,
		time.Now().UTC().Add(claimDuration),
		time.Second*2,
		fmt.Sprintf(
			"%s = $$??",
			api.ExecutionTableChangeIDColumn,
		),
		change.ID,
	)
	require.NoError(t, err)
	require.NotNil(t, execution)

	err = tx.Commit(ctx)
	require.NoError(t, err)

	wd, err := os.Getwd()
	require.NoError(t, err)

	_, last := filepath.Split(t.TempDir())
	tempPath := filepath.Join(wd, "tmp", last)

	err = HandleExecution(ctx, db, execution, apiClient, tempPath)
	require.NoError(t, err)

	// then
}
