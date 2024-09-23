package repository_syncer

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	_config "github.com/initialed85/djangolang/pkg/config"
	"github.com/initialed85/djangolang/pkg/helpers"
	"github.com/initialed85/fred/pkg/api"
)

const syncInterval = time.Second * 30

var log = helpers.GetLogger("repository_syncer")

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

	db, err := _config.GetDBFromEnvironment(ctx)
	if err != nil {
		return err
	}
	defer func() {
		db.Close()
	}()

	t := time.NewTicker(time.Second * 1)

	log.Printf("repository syncer running...")

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
				return nil
			}

			for _, repository := range repositories {
				parts := strings.Split(strings.Trim(repository.URL, "/"), "/")
				repositoryName := parts[len(parts)-1]
				repositoryPath := filepath.Join(repositoriesPath, repositoryName)

				log.Printf("syncing %s", repositoryPath)

				rules, _, _, _, _, err := api.SelectRules(
					ctx,
					tx,
					fmt.Sprintf(
						"%s = '%s'",
						api.RuleTableRepositoryIDColumn,
						repository.ID.String(),
					),
					nil,
					nil,
					nil,
				)
				if err != nil {
					return err
				}

				if len(rules) == 0 {
					log.Printf("no rules for %s; skipping sync", repositoryPath)
					return nil
				}

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
					log.Printf("%s doesn't exist, cloning...", repository.URL)
					gitRepository, err = git.PlainCloneContext(
						ctx,
						repositoryPath,
						false,
						&git.CloneOptions{
							URL:               repository.URL,
							RecurseSubmodules: git.DefaultSubmoduleRecursionDepth, // TODO: should probably be configurable
						},
					)
					if err != nil {
						return fmt.Errorf("clone for %s failed: %s", repository.URL, err.Error())
					}
				} else {
					log.Printf("%s already exists, opening...", repository.URL)
					gitRepository, err = git.PlainOpen(repositoryPath)
					if err != nil {
						return fmt.Errorf("open for %s failed: %s", repository.URL, err.Error())
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
					return fmt.Errorf("reset for %s failed: %s", repository.URL, err.Error())
				}

				refSpecs := make([]config.RefSpec, 0)
				for _, rule := range rules {
					if rule.BranchName == nil {
						continue
					}

					refSpec := fmt.Sprintf("+refs/heads/%s:refs/remotes/origin/%s", *rule.BranchName, *rule.BranchName)

					refSpecs = append(refSpecs, config.RefSpec(refSpec))
				}

				log.Printf("fetching %s...", repositoryPath)
				err = gitRepository.FetchContext(ctx, &git.FetchOptions{
					RefSpecs: refSpecs,
				})
				if err != nil {
					if !errors.Is(err, git.NoErrAlreadyUpToDate) {

						return fmt.Errorf("fetch for %s:%#+v failed: %s", repository.URL, refSpecs, err.Error())
					}
				}

				for _, rule := range rules {
					if rule.BranchName == nil {
						continue
					}

					refName := fmt.Sprintf("refs/remotes/origin/%s", *rule.BranchName)

					log.Printf("checking out %s:%s", repository.URL, refName)

					err = workTree.Checkout(&git.CheckoutOptions{
						Branch: plumbing.ReferenceName(refName),
					})
					if err != nil {
						return fmt.Errorf("checkout for %s:%s failed: %s", repository.URL, refName, err.Error())
					}

					reference, err := gitRepository.Head()
					if err != nil {
						return fmt.Errorf("get head for %s:%s failed: %s", repository.URL, refName, err.Error())
					}

					commitObject, err := gitRepository.CommitObject(reference.Hash())
					if err != nil {
						return fmt.Errorf("get commit for %s:%s failed: %s", repository.URL, refName, err.Error())
					}

					existingChanges, _, _, _, _, err := api.SelectChanges(
						ctx,
						tx,
						fmt.Sprintf(
							"%s = '%s' AND %s = '%s' AND %s = '%s'",
							api.ChangeTableBranchNameColumn,
							*rule.BranchName,
							api.ChangeTableCommitHashColumn,
							commitObject.Hash.String(),
							api.ChangeTableRepositoryIDColumn,
							repository.ID.String(),
						),
						nil,
						nil,
						nil,
					)
					if err != nil {
						return err
					}

					if len(existingChanges) > 0 {
						log.Printf("already know about %s:%s:%s", repository.URL, refName, commitObject.Hash.String())
						continue
					}

					log.Printf("recording change for %s:%s:%s", repository.URL, refName, commitObject.Hash.String())

					change := &api.Change{
						BranchName:   *rule.BranchName,
						CommitHash:   commitObject.Hash.String(),
						Message:      commitObject.Message,
						AuthoredBy:   commitObject.Author.Email,
						AuthoredAt:   commitObject.Author.When,
						CommittedBy:  commitObject.Committer.Email,
						CommittedAt:  commitObject.Committer.When,
						RepositoryID: repository.ID,
					}

					err = change.Insert(ctx, tx, false, false)
					if err != nil {
						if !strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
							return err
						}
					}
				}

				repository.LastSynced = time.Now().UTC()
				err = repository.Update(ctx, tx, false)
				if err != nil {
					return err
				}

				err = tx.Commit(ctx)
				if err != nil {
					return err
				}
			}

			return nil
		}()
		if err != nil {
			return err
		}
	}
}
