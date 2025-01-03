#!/bin/bash

set -e

./migrate.sh down -all || ./migrate.sh force 1

djangolang schema ./database/schema.yaml >./database/migrations/00001_initial.up.sql

echo "" >./database/migrations/00001_initial.down.sql

# shellcheck disable=SC1073
# shellcheck disable=SC2002
for table in $(cat "./database/migrations/00001_initial.up.sql" | grep -A 1 'CREATE TABLE' | grep -v 'CREATE TABLE' | grep -oE 'public\.\w+' | xargs); do
    echo "DROP TABLE IF EXISTS ${table} CASCADE;" >>./database/migrations/00001_initial.down.sql
done

./migrate.sh up

./build.sh

redis-cli flushall

REDIS_URL=redis://localhost:6379 DJANGOLANG_API_ROOT=/api/ POSTGRES_DB=fred POSTGRES_PASSWORD=NoCI\!11 go run ./cmd/api/ serve
