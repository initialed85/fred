#!/bin/bash

set -e

REDIS_URL=redis://localhost:6379
DJANGOLANG_API_ROOT=/api/
POSTGRES_DB=fred
POSTGRES_PASSWORD=NoCI\!11

export REDIS_URL
export DJANGOLANG_API_ROOT
export POSTGRES_DB
export POSTGRES_PASSWORD

redis-cli flushall

# shellcheck disable=SC2086
go test -v -count=1 -failfast ${1:-./...}
