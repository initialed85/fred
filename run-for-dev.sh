#!/bin/bash

set -e

./migrate.sh down -all || ./migrate.sh force 1

./migrate.sh up

./build.sh

redis-cli flushall

REDIS_URL=redis://localhost:6379 DJANGOLANG_API_ROOT=/api/ POSTGRES_DB=fred POSTGRES_PASSWORD=NoCI\!11 go run ./cmd/api/ serve
