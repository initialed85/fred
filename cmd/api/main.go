package main

import (
	"context"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/initialed85/djangolang/pkg/config"
	"github.com/initialed85/djangolang/pkg/server"
	"github.com/initialed85/fred/pkg/api"
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
		// TODO: customer handlers here

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
