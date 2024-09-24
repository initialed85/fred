package internal

import (
	"context"
	_log "log"
	"time"

	_config "github.com/initialed85/djangolang/pkg/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Run(
	log *_log.Logger,
	work func(ctx context.Context, db *pgxpool.Pool) error,
) error {
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

	log.Printf("running...")

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-t.C:
			err = work(ctx, db)
			if err != nil {
				return err
			}
		}
	}

	// never returns via this path
	// return nil
}

func RunWithTx(
	log *_log.Logger,
	work func(ctx context.Context, db *pgxpool.Pool, tx pgx.Tx) error,
) error {
	return Run(
		log,
		func(ctx context.Context, db *pgxpool.Pool) error {
			tx, err := db.Begin(ctx)
			if err != nil {
				return err
			}

			defer func() {
				_ = tx.Rollback(ctx)
			}()

			err = work(ctx, db, tx)
			if err != nil {
				return err
			}

			err = tx.Commit(ctx)
			if err != nil {
				return err
			}

			return nil
		},
	)
}
