package job_coordinator

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
	"github.com/initialed85/fred/pkg/repository_syncer"
	"github.com/stretchr/testify/require"
)

const (
	repositoryURL = "https://github.com/initialed85/djangolang"
)

func TestJobCoordinator(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := _config.GetDBFromEnvironment(ctx)
	require.NoError(t, err)
	defer func() {
		db.Close()
	}()

	t.Run("WithoutDependsOn", func(t *testing.T) {
		//
		// given repository
		//

		repository := &api.Repository{
			URL: repositoryURL,
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
		// given job
		//

		job := &api.Job{
			Name:         "build",
			Branches:     helpers.Ptr("main"),
			RepositoryID: repository.ID,
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
			Platform: nil,
			Image:    "golang:1.23",
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
			Platform: nil,
			Image:    "golang:1.23",
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

		// when

		repositoriesPath := t.TempDir()
		err = repository_syncer.HandleRepository(ctx, tx, repository, repositoriesPath)
		require.NoError(t, err)

		changes, _, _, _, _, err := api.SelectChanges(
			ctx,
			tx,
			fmt.Sprintf(
				"%s = $$?? AND %s = $$??",
				api.ChangeTableRepositoryIDColumn,
				api.ChangeTableBranchColumn,
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

		require.NotNil(t, change.HandledAt)
		require.Greater(t, *change.HandledAt, time.Time{})

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
	})

	t.Run("WithDependsOn", func(t *testing.T) {
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
		// given jobs
		//

		job1 := &api.Job{
			Name:         "build",
			Branches:     helpers.Ptr("main"),
			RepositoryID: repository.ID,
		}

		err = job1.Insert(ctx, tx, false, false)
		require.NoError(t, err)
		defer func() {
			tx, err := db.Begin(ctx)
			require.NoError(t, err)
			defer func() {
				_ = tx.Rollback(ctx)
			}()

			_ = job1.Delete(ctx, tx)
			_ = tx.Commit(ctx)
		}()

		job2 := &api.Job{
			Name:         "test",
			Branches:     helpers.Ptr("main"),
			RepositoryID: repository.ID,
		}

		err = job2.Insert(ctx, tx, false, false)
		require.NoError(t, err)
		defer func() {
			tx, err := db.Begin(ctx)
			require.NoError(t, err)
			defer func() {
				_ = tx.Rollback(ctx)
			}()

			_ = job2.Delete(ctx, tx)
			_ = tx.Commit(ctx)
		}()

		dependsOn := &api.DependsOn{
			SourceJobID: job1.ID,
			SinkJobID:   job2.ID,
		}

		err = dependsOn.Insert(ctx, tx, false, false)
		require.NoError(t, err)
		defer func() {
			tx, err := db.Begin(ctx)
			require.NoError(t, err)
			defer func() {
				_ = tx.Rollback(ctx)
			}()

			_ = dependsOn.Delete(ctx, tx)
			_ = tx.Commit(ctx)
		}()

		//
		// given tasks
		//

		job1Task1 := &api.Task{
			Name:     "lint",
			Index:    0,
			Platform: nil,
			Image:    "golang:1.23",
			Script:   "./lint.sh",
			JobID:    job1.ID,
		}

		err = job1Task1.Insert(ctx, tx, false, false)
		require.NoError(t, err)
		defer func() {
			tx, err := db.Begin(ctx)
			require.NoError(t, err)
			defer func() {
				_ = tx.Rollback(ctx)
			}()

			_ = job1Task1.Delete(ctx, tx)
			_ = tx.Commit(ctx)
		}()

		job1Task2 := &api.Task{
			Name:     "build",
			Index:    0,
			Platform: nil,
			Image:    "golang:1.23",
			Script:   "go build -o bin/cmd ./cmd/",
			JobID:    job1.ID,
		}

		err = job1Task2.Insert(ctx, tx, false, false)
		require.NoError(t, err)
		defer func() {
			tx, err := db.Begin(ctx)
			require.NoError(t, err)
			defer func() {
				_ = tx.Rollback(ctx)
			}()

			_ = job1Task2.Delete(ctx, tx)
			_ = tx.Commit(ctx)
		}()

		job2Task1 := &api.Task{
			Name:     "build",
			Index:    0,
			Platform: nil,
			Image:    "golang:1.23",
			Script:   "./test.sh",
			JobID:    job2.ID,
		}

		err = job2Task1.Insert(ctx, tx, false, false)
		require.NoError(t, err)
		defer func() {
			tx, err := db.Begin(ctx)
			require.NoError(t, err)
			defer func() {
				_ = tx.Rollback(ctx)
			}()

			_ = job1Task2.Delete(ctx, tx)
			_ = tx.Commit(ctx)
		}()

		// when

		repositoriesPath := t.TempDir()
		err = repository_syncer.HandleRepository(ctx, tx, repository, repositoriesPath)
		require.NoError(t, err)

		changes, _, _, _, _, err := api.SelectChanges(
			ctx,
			tx,
			fmt.Sprintf(
				"%s = $$?? AND %s = $$??",
				api.ChangeTableRepositoryIDColumn,
				api.ChangeTableBranchColumn,
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

		require.Nil(t, change.HandledAt)

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
			helpers.Ptr(2),
			nil,
			change.ID,
		)
		require.NoError(t, err)
		require.NotEmpty(t, executions)
		require.Len(t, executions, 1)

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
	})
}
