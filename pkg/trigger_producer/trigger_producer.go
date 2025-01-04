package trigger_producer

import (
	"context"
	"fmt"
	"time"

	"github.com/initialed85/djangolang/pkg/helpers"
	"github.com/initialed85/djangolang/pkg/query"
	"github.com/initialed85/fred/internal"
	"github.com/initialed85/fred/pkg/api"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var log = helpers.GetLogger("trigger_producer")

const claimDuration = time.Minute * 5

func HandleChange(ctx context.Context, tx pgx.Tx, change *api.Change) error {
	log.Printf("handling %s", internal.GetChangeSummary(change))

	rules, _, _, _, _, err := api.SelectRules(
		query.WithLoad(ctx, fmt.Sprintf("referenced_by_%s", api.TriggerTable)),
		tx,
		fmt.Sprintf(
			"%s = $$??",
			api.RuleTableRepositoryIDColumn,
		),
		nil,
		nil,
		nil,
		change.RepositoryID,
	)
	if err != nil {
		return err
	}

	if len(rules) != 0 {
		for _, rule := range rules {
			if rule.BranchName == nil {
				continue
			}

			possibleBranchName := *rule.BranchName

			if !(possibleBranchName == "*" || possibleBranchName == change.BranchName) {
				continue
			}

			if rule.ReferencedByTriggerRuleIDObjects == nil {
				return fmt.Errorf("assertion failed: rule.ReferencedByTriggerRuleIDObjects was unexpectedly nil")
			}

			for _, trigger := range rule.ReferencedByTriggerRuleIDObjects {
				err = trigger.Reload(query.WithLoad(ctx, api.JobTable), tx)
				if err != nil {
					return fmt.Errorf("attempt to reload %#+v failed: %s", trigger, err)
				}

				if trigger.JobIDObject == nil {
					return fmt.Errorf("assertion failed: trigger.JobIDObject was unexpectedly nil")
				}

				job := trigger.JobIDObject

				err = job.Reload(query.WithLoad(ctx, fmt.Sprintf("referenced_by_%s", api.TaskTable)), tx)
				if err != nil {
					return fmt.Errorf("attempt to reload %#+v failed: %s", job, err)
				}

				execution := &api.Execution{
					Status:    internal.ExecutionOrTaskStatusPending,
					ChangeID:  change.ID,
					JobID:     job.ID,
					TriggerID: trigger.ID,
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
	}

	change.TriggersProducedAt = helpers.Ptr(time.Now().UTC())
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
			change, err := api.TriggerProducerClaimChange(
				query.WithLoad(ctx, api.RepositoryTable),
				tx,
				time.Now().UTC().Add(claimDuration),
				time.Second*10,
				fmt.Sprintf(
					"%s IS null",
					api.ChangeTableTriggersProducedAtColumn,
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
					change.TriggerProducerClaimedUntil = time.Now().UTC()
					err = change.Update(ctx, tx, false)
					if err != nil {
						err = fmt.Errorf("attempt to release claim on %s:%s@%s failed: %s", change.RepositoryIDObject.URL, change.BranchName, change.CommitHash, err.Error())
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
