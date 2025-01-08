package repository_syncer

import (
	"context"
	"fmt"
	"testing"
	"time"

	_config "github.com/initialed85/djangolang/pkg/config"
	"github.com/initialed85/djangolang/pkg/helpers"
	"github.com/initialed85/fred/pkg/api"
	"github.com/stretchr/testify/require"
)

func TestRepositorySyncer(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := _config.GetDBFromEnvironment(ctx)
	require.NoError(t, err)
	defer func() {
		db.Close()
	}()

	repository := &api.Repository{
		URL: "https://github.com/initialed85/djangolang",
	}

	tx, err := db.Begin(ctx)
	require.NoError(t, err)
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	//
	// given repository
	//

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
	// when
	//

	repositoriesPath := t.TempDir()
	err = HandleRepository(ctx, tx, repository, repositoriesPath)
	require.NoError(t, err)

	//
	// then
	//

	err = repository.Reload(ctx, tx)
	require.NoError(t, err)

	require.Greater(t, repository.HandledAt, time.Time{})
	require.NotNil(t, repository.Name)
	require.Equal(t, "djangolang", *repository.Name)

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
	require.NotNil(t, change)
	require.Equal(t, "main", change.Branch)
	require.Equal(t, "initialed85@gmail.com", change.AuthoredBy)

	err = tx.Commit(ctx)
	require.NoError(t, err)
}
