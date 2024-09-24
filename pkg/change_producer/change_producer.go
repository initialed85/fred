package change_producer

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
	"github.com/initialed85/djangolang/pkg/helpers"
	"github.com/initialed85/djangolang/pkg/query"
	"github.com/initialed85/fred/internal"
	"github.com/initialed85/fred/pkg/api"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const syncInterval = time.Second * 30

var log = helpers.GetLogger("change_producer")

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

	return internal.RunWithTx(
		log,
		func(ctx context.Context, db *pgxpool.Pool, tx pgx.Tx) error {
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
							SingleBranch:      false,
							Depth:             1,
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

				remotes, err := gitRepository.Remotes()
				if err != nil {
					return fmt.Errorf("get remotes for %s failed: %s", repository.URL, err.Error())
				}

				if len(remotes) == 0 {
					return fmt.Errorf("get remotes for %s failed: %s", repository.URL, "no remotes")
				}

				remote := remotes[0]

				allRefs, err := remote.ListContext(ctx, &git.ListOptions{})
				if err != nil {
					return fmt.Errorf("list remote for %s failed: %s", repository.URL, "no remotes")
				}

				branchRefs := make([]*plumbing.Reference, 0)

				for _, ref := range allRefs {
					refName := ref.Name().String()

					if !strings.HasPrefix(refName, "refs/heads/") {
						continue
					}

					branchRefs = append(branchRefs, ref)
				}

				refSpecs := make([]config.RefSpec, 0)
				for _, branchRef := range branchRefs {
					refName := branchRef.Name().String() // e.g. refs/heads/refactor-1
					parts := strings.Split(refName, "/")
					branchName := parts[len(parts)-1] // e.g. refactor-1
					remoteBranchName := fmt.Sprintf("remotes/origin/%s", branchName)
					remoteRefName := fmt.Sprintf("refs/%s", remoteBranchName)

					refSpec := fmt.Sprintf("+%s:%s", refName, remoteRefName)
					refSpecs = append(refSpecs, config.RefSpec(refSpec))
				}

				log.Printf("fetching %s:%#+v...", repositoryPath, refSpecs)

				err = gitRepository.FetchContext(ctx, &git.FetchOptions{
					RefSpecs: refSpecs,
				})
				if err != nil {
					if !errors.Is(err, git.NoErrAlreadyUpToDate) {

						return fmt.Errorf("fetch for %s:%#+v failed: %s", repository.URL, refSpecs, err.Error())
					}
				}

				for _, branchRef := range branchRefs {
					refName := branchRef.Name().String() // e.g. refs/heads/refactor-1
					parts := strings.Split(refName, "/")
					branchName := parts[len(parts)-1] // e.g. refactor-1
					remoteBranchName := fmt.Sprintf("remotes/origin/%s", branchName)
					remoteRefName := fmt.Sprintf("refs/%s", remoteBranchName)

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

					err = workTree.Clean(&git.CleanOptions{
						Dir: true,
					})
					if err != nil {
						return fmt.Errorf("clean for %s failed: %s", repository.URL, err.Error())
					}

					log.Printf("checking out %s:%s", repository.URL, refName)

					err = workTree.Checkout(&git.CheckoutOptions{
						Branch: plumbing.ReferenceName(remoteRefName),
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
							"%s = $$?? AND %s = $$?? AND %s = $$??",
							api.ChangeTableBranchNameColumn,
							api.ChangeTableCommitHashColumn,
							api.ChangeTableRepositoryIDColumn,
						),
						nil,
						nil,
						nil,
						branchName,
						commitObject.Hash.String(),
						repository.ID,
					)
					if err != nil {
						return err
					}

					if len(existingChanges) > 0 {
						log.Printf("already know about %s:%s:%s", repository.URL, refName, commitObject.Hash.String())
						continue
					}

					change := &api.Change{
						BranchName:   branchName,
						CommitHash:   commitObject.Hash.String(),
						Message:      commitObject.Message,
						AuthoredBy:   commitObject.Author.Email,
						AuthoredAt:   commitObject.Author.When,
						CommittedBy:  commitObject.Committer.Email,
						CommittedAt:  commitObject.Committer.When,
						RepositoryID: repository.ID,
					}

					ctx = query.WithMaxDepth(ctx, helpers.Ptr(2))

					err = change.Insert(ctx, tx, false, false)
					if err != nil {
						return fmt.Errorf("failed to produce %s: %s", internal.GetChangeSummary(change), err)
					}

					log.Printf("produced %s", internal.GetChangeSummary(change))
				}

				repository.LastSynced = time.Now().UTC()
				err = repository.Update(ctx, tx, false)
				if err != nil {
					return err
				}
			}

			return nil
		},
	)
}
