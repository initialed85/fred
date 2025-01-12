package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/initialed85/djangolang/pkg/config"
	"github.com/initialed85/djangolang/pkg/query"
	"github.com/initialed85/djangolang/pkg/server"
	"github.com/initialed85/fred/pkg/api"
	"github.com/initialed85/fred/pkg/job_coordinator"
	"github.com/jackc/pgx/v5/pgxpool"
)

var log = api.ThisLogger()

type RunJobPathParams struct {
	JobName    string `json:"job_name"`
	CommitHash string `json:"commit_hash"`
}

func addBaseCustomHandlers(db *pgxpool.Pool, r chi.Router) error {
	postHandler, err := server.GetHTTPHandler(
		http.MethodPut,
		"/trigger-job/{job_name}/{commit_hash}",
		http.StatusCreated,
		func(
			ctx context.Context,
			pathParams RunJobPathParams,
			queryParams server.EmptyQueryParams,
			req server.EmptyRequest,
			rawReq any,
		) (server.Response[api.Execution], error) {
			tx, err := db.Begin(ctx)
			if err != nil {
				return server.Response[api.Execution]{}, fmt.Errorf("failed to begin DB transaction; %v", err)
			}

			defer func() {
				_ = tx.Rollback(ctx)
			}()

			job, _, _, _, _, err := api.SelectJob(
				query.WithLoad(ctx, fmt.Sprintf("referenced_by_%s", api.TaskTable)),
				tx,
				fmt.Sprintf("%s = $$??", api.JobTableNameColumn),
				pathParams.JobName,
			)
			if err != nil {
				return server.Response[api.Execution]{}, fmt.Errorf(
					"failed to get job for job name %#+v; %v",
					pathParams.JobName, err,
				)
			}

			change, _, _, _, _, err := api.SelectChange(
				ctx,
				tx,
				fmt.Sprintf("%s = $$??", api.ChangeTableCommitHashColumn),
				pathParams.CommitHash,
			)
			if err != nil {
				return server.Response[api.Execution]{}, fmt.Errorf(
					"failed to get change for commit hash %#+v; %v",
					pathParams.CommitHash, err,
				)
			}

			execution, err := job_coordinator.CreateExecution(
				ctx,
				tx,
				change,
				job,
			)
			if err != nil {
				return server.Response[api.Execution]{}, fmt.Errorf(
					"failed to create execution for job name %#+v and commit hash %#+v; %v",
					pathParams.JobName, pathParams.CommitHash, err,
				)
			}

			err = tx.Commit(ctx)
			if err != nil {
				return server.Response[api.Execution]{}, fmt.Errorf(
					"failed to commit DB transaction; %v", err,
				)
			}

			return server.Response[api.Execution]{
				Status:     http.StatusOK,
				Success:    true,
				Error:      nil,
				Objects:    []*api.Execution{execution},
				Count:      1,
				TotalCount: 1,
				Limit:      1,
				Offset:     0,
			}, nil
		},
	)
	if err != nil {
		return err
	}
	r.Post(postHandler.FullPath, postHandler.ServeHTTP)

	return nil
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
		err := addBaseCustomHandlers(db, r)
		if err != nil {
			return err
		}

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
