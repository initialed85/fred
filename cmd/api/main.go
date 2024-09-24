package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/initialed85/djangolang/pkg/config"
	"github.com/initialed85/djangolang/pkg/server"
	"github.com/initialed85/fred/pkg/api"
	"github.com/initialed85/fred/pkg/job_executor"
)

var log = api.ThisLogger()

type ClaimRequest struct {
	ClaimDurationSeconds float64 `json:"claim_duration_seconds"`
}

func RunServeWithEnvironment(
	httpMiddlewares []server.HTTPMiddleware,
	objectMiddlewares []server.ObjectMiddleware,
	addCustomHandlers func(chi.Router) error,
) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	port := config.Port()

	db, err := config.GetDBFromEnvironment(ctx)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer func() {
		db.Close()
	}()

	redisPool, err := config.GetRedisFromEnvironment()
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer func() {
		_ = redisPool.Close()
	}()

	actualAddCustomHandlers := func(r chi.Router) error {
		claimVideoForObjectDetectorHandler, err := api.GetHTTPHandler(
			http.MethodPatch,
			"/custom/claim-trigger-for-job-executor",
			http.StatusOK,
			func(
				ctx context.Context,
				pathParams server.EmptyPathParams,
				queryParams server.EmptyQueryParams,
				req ClaimRequest,
				rawReq any,
			) (*api.Trigger, error) {
				now := time.Now().UTC()

				claimUntil := now.Add(time.Second * time.Duration(req.ClaimDurationSeconds))

				if claimUntil.Sub(now) <= 0 {
					return nil, fmt.Errorf("claim_duration_seconds too short; must result in a claim that expires in the future")
				}

				tx, err := db.Begin(ctx)
				if err != nil {
					return nil, err
				}

				defer func() {
					_ = tx.Rollback(ctx)
				}()

				trigger, err := job_executor.ClaimTriggerForJobExecutor(ctx, tx, now.Sub(claimUntil))
				if err != nil {
					return nil, err
				}

				err = tx.Commit(ctx)
				if err != nil {
					return nil, err
				}

				return trigger, nil
			},
		)
		if err != nil {
			return err
		}

		r.Patch(claimVideoForObjectDetectorHandler.PathWithinRouter, claimVideoForObjectDetectorHandler.ServeHTTP)

		if addCustomHandlers != nil {
			err = addCustomHandlers(r)
			if err != nil {
				return err
			}
		}

		return nil
	}

	api.RunServeWithArguments(ctx, cancel, port, db, redisPool, httpMiddlewares, objectMiddlewares, actualAddCustomHandlers)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("first argument must be command (one of 'dump-config', 'dump-openapi-json', 'dump-openapi-yaml' or 'serve')")
	}

	command := strings.TrimSpace(strings.ToLower(os.Args[1]))

	switch command {

	case "dump-config":
		config.DumpConfig()

	case "dump-openapi-json":
		api.RunDumpOpenAPIJSON()

	case "dump-openapi-yaml":
		api.RunDumpOpenAPIYAML()

	case "serve":
		RunServeWithEnvironment(nil, nil, nil)
	}
}
