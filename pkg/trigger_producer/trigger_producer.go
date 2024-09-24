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

func Run() error {
	return internal.RunWithTx(
		log,
		func(ctx context.Context, db *pgxpool.Pool, tx pgx.Tx) error {
			// there's probably no need to run multiple trigger producer replicas, but if by chance you do, this should ensure consistency
			err := (&api.Trigger{}).AdvisoryLockWithRetries(ctx, tx, internal.TriggerProducerAdvisoryLockID, time.Second*10, time.Second*1)
			if err != nil {
				return err
			}

			ctx = query.WithMaxDepth(ctx, helpers.Ptr(2))

			changes, _, _, _, _, err := api.SelectChanges(
				ctx,
				tx,
				fmt.Sprintf(
					"%s IS null",
					api.ChangeTableTriggerProducedAtColumn,
				),
				helpers.Ptr(fmt.Sprintf(
					"%s ASC",
					api.ChangeTableCreatedAtColumn,
				)),
				nil,
				nil,
			)
			if err != nil {
				return err
			}

			for _, change := range changes {
				changeSummary := internal.GetChangeSummary(change)

				log.Printf("handling %s", changeSummary)

				rules, _, _, _, _, err := api.SelectRules(
					ctx,
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

						trigger := &api.Trigger{
							RuleID:   rule.ID,
							ChangeID: change.ID,
						}

						err = trigger.Insert(ctx, tx, false, false)
						if err != nil {
							return fmt.Errorf("failed to produce trigger for %s: %s", internal.GetChangeSummary(change), err)
						}

						log.Printf("produced %s", internal.GetTriggerSummary(trigger))
					}
				}

				change.TriggerProducedAt = helpers.Ptr(time.Now().UTC())
				err = change.Update(ctx, tx, false)
				if err != nil {
					return fmt.Errorf("failed to produce trigger for %s: %s", internal.GetChangeSummary(change), err)
				}
			}

			return nil
		},
	)
}
