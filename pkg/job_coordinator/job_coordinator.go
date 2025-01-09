package job_coordinator

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/initialed85/djangolang/pkg/helpers"
	"github.com/initialed85/djangolang/pkg/query"
	"github.com/initialed85/fred/internal"
	"github.com/initialed85/fred/pkg/api"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var log = helpers.GetLogger("job_coordinator")

const claimDuration = time.Minute * 5

func HandleChange(ctx context.Context, tx pgx.Tx, change *api.Change) error {
	log.Printf("handling %s", internal.GetChangeSummary(change))

	jobs, _, _, _, _, err := api.SelectJobs(
		query.WithLoad(ctx, fmt.Sprintf("referenced_by_%s", api.TaskTable)),
		tx,
		fmt.Sprintf(
			"%s = $$??",
			api.JobTableRepositoryIDColumn,
		),
		nil,
		nil,
		nil,
		change.RepositoryID,
	)
	if err != nil {
		return err
	}

	handledAllJobs := true

	if len(jobs) != 0 {
		for _, job := range jobs {
			if job.Branches != nil {
				branches := *job.Branches

				expr, err := regexp.Compile(branches)
				if err != nil {
					return fmt.Errorf("attempt to treat branch expression for %#+v failed: %s", job, err)
				}

				if !expr.MatchString(change.Branch) {
					log.Printf("%s did not match %s by branch; skipping...", internal.GetJobSummary(job), internal.GetChangeSummary(change))
					continue
				}
			}

			if change.Tag != nil && job.Tags != nil {
				tags := *job.Tags

				expr, err := regexp.Compile(tags)
				if err != nil {
					return fmt.Errorf("attempt to treat branch expression for %#+v failed: %s", job, err)
				}

				if !expr.MatchString(*change.Tag) {
					log.Printf("%s did not match %s by tag; skipping...", internal.GetJobSummary(job), internal.GetChangeSummary(change))
					continue
				}
			}

			existingExecutions, _, _, _, _, err := api.SelectExecutions(
				query.WithLoad(ctx, api.JobTable),
				tx,
				fmt.Sprintf(
					"%s = $$?? AND %s = $$??",
					api.ExecutionTableJobIDColumn,
					api.ExecutionTableChangeIDColumn,
				),
				nil,
				nil,
				nil,
				job.ID,
				change.ID,
			)
			if err != nil {
				return err
			}

			if len(existingExecutions) > 0 {
				log.Printf("%s already had %d existing executions; skipping...", internal.GetJobSummary(job), len(existingExecutions))
				continue
			}

			dependsOns, _, _, _, _, err := api.SelectDependsOns(
				query.WithLoad(ctx, api.JobTable),
				tx,
				fmt.Sprintf(
					"%s = $$??",
					api.DependsOnTableSinkJobIDColumn,
				),
				nil,
				nil,
				nil,
				job.ID,
			)
			if err != nil {
				return err
			}

			skip := false

			for _, dependsOn := range dependsOns {
				sourceExecutions, _, _, _, _, err := api.SelectExecutions(
					query.WithLoad(ctx, api.JobTable),
					tx,
					fmt.Sprintf(
						"%s = $$?? AND %s = $$??",
						api.ExecutionTableJobIDColumn,
						api.ExecutionTableChangeIDColumn,
					),
					nil,
					nil,
					nil,
					dependsOn.SourceJobID,
					change.ID,
				)
				if err != nil {
					return err
				}

				if len(sourceExecutions) == 0 {
					log.Printf(
						"%s will be skipped because depends-on %s had no executions",
						internal.GetJobSummary(job),
						internal.GetJobSummary(dependsOn.SourceJobIDObject),
					)

					skip = true

					continue
				}

				for _, sourceExecution := range sourceExecutions {
					if sourceExecution.Status != internal.ExecutionOrTaskStatusSucceeded {
						log.Printf(
							"%s will be skipped because depends-on %s had %s in status other than %s",
							internal.GetJobSummary(job),
							internal.GetJobSummary(dependsOn.SourceJobIDObject),
							internal.GetExecutionSummary(sourceExecution),
							internal.ExecutionOrTaskStatusSucceeded,
						)
						skip = true
					}
				}
			}

			if skip {
				handledAllJobs = false
				continue
			}

			execution := &api.Execution{
				Status:   internal.ExecutionOrTaskStatusPending,
				ChangeID: change.ID,
				JobID:    job.ID,
			}

			err = execution.Insert(query.WithLoad(ctx, api.JobTable), tx, false, false)
			if err != nil {
				return fmt.Errorf("attempt to insert %#+v failed: %s", execution, err)
			}

			for _, task := range job.ReferencedByTaskJobIDObjects {
				logObj := &api.Log{}

				err = logObj.Insert(ctx, tx, false, false)
				if err != nil {
					return fmt.Errorf("attempt to insert %#+v failed: %s", logObj, err)
				}

				output := &api.Output{
					Status:      internal.ExecutionOrTaskStatusPending,
					ExecutionID: execution.ID,
					TaskID:      task.ID,
					LogID:       logObj.ID,
				}

				err = output.Insert(ctx, tx, false, false)
				if err != nil {
					return fmt.Errorf("attempt to insert %#+v failed: %s", output, err)
				}

				logObj.OutputID = output.ID
				err = logObj.Update(ctx, tx, false)
				if err != nil {
					return fmt.Errorf("attempt to update %#+v failed: %s", output, err)
				}
			}

			log.Printf("produced execution %s", internal.GetExecutionSummary(execution))
		}
	}

	if handledAllJobs {
		change.HandledAt = helpers.Ptr(time.Now().UTC())
	}

	err = change.Update(ctx, tx, false)
	if err != nil {
		return fmt.Errorf("failed to produce trigger for %s: %s", internal.GetChangeSummary(change), err)
	}

	return nil
}

func Run() error {
	lastRunFoundNothing := false

	return internal.RunWithTx(
		log,
		func(ctx context.Context, db *pgxpool.Pool, tx pgx.Tx) error {
			change, err := api.JobCoordinatorClaimChange(
				query.WithLoad(ctx, api.RepositoryTable),
				tx,
				time.Now().UTC().Add(claimDuration),
				time.Second*10,
				fmt.Sprintf(
					"%s IS null",
					api.ChangeTableHandledAtColumn,
				),
			)
			if err != nil {
				return fmt.Errorf("attempt to claim a change failed: %s", err.Error())
			}

			if change == nil {
				if !lastRunFoundNothing {
					lastRunFoundNothing = true
					log.Printf("no changes to claim")
				}

				return nil
			}

			lastRunFoundNothing = false

			defer func() {
				if change != nil {
					change.JobCoordinatorClaimedUntil = time.Now().UTC()
					err = change.Update(ctx, tx, false)
					if err != nil {
						err = fmt.Errorf("attempt to release claim on %s:%s@%s failed: %s", change.RepositoryIDObject.URL, change.Branch, change.CommitHash, err.Error())
						log.Printf("warning: %s", err.Error())
					}
				}
			}()

			err = HandleChange(ctx, tx, change)
			if err != nil {
				return fmt.Errorf("attempt to handle %s failed: %s", internal.GetChangeSummary(change), err.Error())
			}

			return nil
		},
	)
}
