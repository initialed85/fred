package repository_syncer

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/initialed85/djangolang/pkg/config"
	"github.com/initialed85/djangolang/pkg/helpers"
	"github.com/initialed85/fred/pkg/api"
)

const syncInterval = time.Second * 60

func Run() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	repositoriesPath := filepath.Join(wd, "repositories")

	err = os.MkdirAll(repositoriesPath, 0o777)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := config.GetDBFromEnvironment(ctx)
	if err != nil {
		return err
	}
	defer func() {
		db.Close()
	}()

	t := time.NewTicker(time.Second * 1)

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-t.C:
		}

		err := func() error {
			tx, err := db.Begin(ctx)
			if err != nil {
				return err
			}

			defer func() {
				_ = tx.Rollback(ctx)
			}()

			repositories, _, _, _, _, err := api.SelectRepositories(
				ctx,
				tx,
				fmt.Sprintf(
					"now() - %s > interval '%f'",
					api.RepositoryTableLastSyncedColumn,
					syncInterval.Seconds(),
				),
				helpers.Ptr(fmt.Sprintf(
					"%s DESC",
					api.RepositoryTableLastSyncedColumn,
				)),
				nil,
				nil,
			)
			if err != nil {
				return err
			}

			if len(repositories) == 0 {
				log.Printf("no repositories...")
				return nil
			}

			for _, repository := range repositories {
				parts := strings.Split(strings.Trim(repository.URL, "/"), "/")
				repositoryName := parts[len(parts)-1]
				repositoryPath := filepath.Join(repositoriesPath, repositoryName)

				log.Printf("syncing %s", repositoryPath)

				exists := true
				_, err = os.Stat(repositoryPath)
				if err != nil {
					if !os.IsNotExist(err) {
						return err
					}

					exists = false
				}

				var gitRepository *git.Repository

				if !exists {
					log.Printf("%s doesn't exist, cloning...", repositoryPath)
					gitRepository, err = git.PlainCloneContext(
						ctx,
						repositoryPath,
						false,
						&git.CloneOptions{
							URL:               repository.URL,
							RecurseSubmodules: git.DefaultSubmoduleRecursionDepth, // TODO: should probably be configurable
							Progress:          os.Stdout,
						},
					)
					if err != nil {
						return fmt.Errorf("clone %s failed: %s", repositoryPath, err.Error())
					}
				} else {
					log.Printf("%s already exists, opening...", repositoryPath)
					gitRepository, err = git.PlainOpen(repositoryPath)
					if err != nil {
						return fmt.Errorf("open %s failed: %s", repositoryPath, err.Error())
					}
				}

				workTree, err := gitRepository.Worktree()
				if err != nil {
					return err
				}

				log.Printf("resetting %s...", repositoryPath)
				err = workTree.Reset(&git.ResetOptions{
					Mode: git.HardReset,
				})
				if err != nil {
					return fmt.Errorf("reset %s failed: %s", repositoryPath, err.Error())
				}

				log.Printf("fetching %s...", repositoryPath)
				err = gitRepository.FetchContext(ctx, &git.FetchOptions{
					Progress: os.Stdout,
				})
				if err != nil {
					if !errors.Is(err, git.NoErrAlreadyUpToDate) {

						return fmt.Errorf("fetch %s failed: %s", repositoryPath, err.Error())
					}
				}

				repository.LastSynced = time.Now().UTC()
				err = repository.Update(ctx, tx, false)
				if err != nil {
					return fmt.Errorf("failed to update DB for %s: %s", repositoryPath, err.Error())
				}

				err = tx.Commit(ctx)
				if err != nil {
					return fmt.Errorf("failed to commit DB for %s: %s", repositoryPath, err.Error())
				}
			}

			return nil
		}()
		if err != nil {
			return err
		}
	}
}
