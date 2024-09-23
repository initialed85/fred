package main

import (
	"os"
	"strings"

	"github.com/initialed85/djangolang/pkg/config"
	"github.com/initialed85/fred/pkg/api"
)

var log = api.ThisLogger()

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
		api.RunServeWithEnvironment(nil, nil, nil)
	}
}
