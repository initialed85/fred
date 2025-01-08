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
	"github.com/initialed85/fred/internal"
	"github.com/initialed85/fred/pkg/api"
	"github.com/initialed85/fred/pkg/job_coordinator"
	"github.com/initialed85/fred/pkg/repository_syncer"
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

	job2 := &api.Job{
		Name:         "test",
		Branches:     helpers.Ptr("main"),
		RepositoryID: repository.ID,
	}

	err = job2.Insert(ctx, tx, false, false)
	require.NoError(t, err)

	dependsOn := &api.DependsOn{
		SourceJobID: job1.ID,
		SinkJobID:   job2.ID,
	}

	err = dependsOn.Insert(ctx, tx, false, false)
	require.NoError(t, err)

	//
	// given tasks
	//

	job1Task1 := &api.Task{
		Name:     "lint",
		Index:    0,
		Platform: "linux/amd64",
		Image:    "initialed85/the-last-ci-image-you-will-ever-need:latest",
		Script:   "./lint.sh",
		JobID:    job1.ID,
	}

	err = job1Task1.Insert(ctx, tx, false, false)
	require.NoError(t, err)

	job1Task2 := &api.Task{
		Name:     "build",
		Index:    1,
		Platform: "linux/amd64",
		Image:    "initialed85/the-last-ci-image-you-will-ever-need:latest",
		Script:   "go build -o bin/cmd ./cmd/",
		JobID:    job1.ID,
	}

	err = job1Task2.Insert(ctx, tx, false, false)
	require.NoError(t, err)

	job2Task1 := &api.Task{
		Name:     "build",
		Index:    0,
		Platform: "linux/amd64",
		Image:    "initialed85/the-last-ci-image-you-will-ever-need:latest",
		Script:   "./test.sh",
		JobID:    job2.ID,
	}

	err = job2Task1.Insert(ctx, tx, false, false)
	require.NoError(t, err)

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

	err = job_coordinator.HandleChange(ctx, tx, change)
	require.NoError(t, err)

	execution1, err := api.JobExecutorClaimExecution(
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
	require.NotNil(t, execution1)

	require.NotNil(t, execution1.JobIDObject)
	require.Equal(t, "build", execution1.JobIDObject.Name)

	execution2, err := api.JobExecutorClaimExecution(
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
	require.Nil(t, execution2)

	err = tx.Commit(ctx)
	require.NoError(t, err)

	wd, err := os.Getwd()
	require.NoError(t, err)

	tempPath := filepath.Join(wd, "tmp")

	err = HandleExecution(ctx, db, execution1, apiClient, tempPath)
	require.NoError(t, err)

	// then

	func() {
		tx, err := db.Begin(ctx)
		require.NoError(t, err)
		defer func() {
			_ = tx.Rollback(ctx)
		}()

		err = execution1.Reload(ctx, tx)
		require.NoError(t, err)
		require.Equal(t, internal.ExecutionOrTaskStatusSucceeded, execution1.Status)

		outputs, _, _, _, _, err := api.SelectOutputs(
			query.WithLoad(ctx, api.LogTable),
			tx,
			fmt.Sprintf("%s = $$??", api.OutputTableExecutionIDColumn),
			nil,
			nil,
			nil,
			execution1.ID,
		)
		require.NoError(t, err)

		require.Len(t, outputs, 2)
		require.Equal(t, int64(0), outputs[0].ExitStatus)
		require.Equal(t, int64(0), outputs[1].ExitStatus)

		_ = tx.Commit(ctx)
	}()

	// when

	func() {
		tx, err := db.Begin(ctx)
		require.NoError(t, err)
		defer func() {
			_ = tx.Rollback(ctx)
		}()
		err = job_coordinator.HandleChange(ctx, tx, change)
		require.NoError(t, err)

		execution2, err = api.JobExecutorClaimExecution(
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
		require.NotNil(t, execution2)

		err = tx.Commit(ctx)
		require.NoError(t, err)

		err = HandleExecution(ctx, db, execution2, apiClient, tempPath)
		require.NoError(t, err)
		_ = tx.Commit(ctx)
	}()

	// then

	func() {
		tx, err := db.Begin(ctx)
		require.NoError(t, err)
		defer func() {
			_ = tx.Rollback(ctx)
		}()

		err = execution2.Reload(ctx, tx)
		require.NoError(t, err)
		require.Equal(t, internal.ExecutionOrTaskStatusSucceeded, execution2.Status)

		outputs, _, _, _, _, err := api.SelectOutputs(
			query.WithLoad(ctx, api.LogTable),
			tx,
			fmt.Sprintf("%s = $$??", api.OutputTableExecutionIDColumn),
			nil,
			nil,
			nil,
			execution2.ID,
		)
		require.NoError(t, err)

		require.Len(t, outputs, 1)
		require.Equal(t, int64(0), outputs[0].ExitStatus)

		_ = tx.Commit(ctx)
	}()
}
