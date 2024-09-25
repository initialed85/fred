#!/bin/bash

set -e -m

# this block ensures we can invoke this script from anywhere and have it automatically change to this folder first
pushd "$(dirname -- "${BASH_SOURCE[0]}")" >/dev/null 2>&1
function teardown() {
    popd >/dev/null 2>&1 || true
}
trap teardown exit

# # we need oapi-codegen to generate the client for use by Go code
if ! command -v oapi-codegen >/dev/null 2>&1; then
    go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@main
fi

# we need npm to generate the client for use by the frontend
if ! command -v npm >/dev/null 2>&1; then
    echo "error: can't find npm command- you likely need to install node / npm"
    exit 1
fi

# we don't need natscli for tooling, but it's a handy debug tool
if ! command -v nats >/dev/null 2>&1; then
    go install github.com/nats-io/natscli/nats@latest
fi

# ensure we've got a djangolang executable available (required for templating)
if [[ "${FORCE_UPDATE_DJANGOLANG}" == "1" ]] || ! command -v djangolang >/dev/null 2>&1; then
    GOPRIVATE="${GOPRIVATE:-}" go install github.com/initialed85/djangolang@latest
    GOPRIVATE="${GOPRIVATE:-}" go get -u github.com/initialed85/djangolang@latest
fi

# ensure the docker compose environment is already running
if [[ "${SKIP_DOCKER_COMPOSE_CHECK}" != "1" ]]; then
    echo -e "\nwaiting for healthy docker compose environment..."
    while ! docker compose ps -a | grep post-migrate | grep 'Exited (0)' >/dev/null 2>&1; do
        sleep 0.1
    done
fi

# introspect the database and generate the Djangolang API
# note: the environment variables are coupled to the environment described in docker-compose.yaml
echo -e "\ngenerating the api..."
DJANGOLANG_API_ROOT=/api DJANGOLANG_PACKAGE_NAME=api POSTGRES_DB=fred POSTGRES_PASSWORD=NoCI!11 djangolang template

# TODO: you don't want to do this once you put some custom stuff in your api entrypoint
# if test -e ./cmd/api; then
#     rm -frv ./cmd/api
# fi
# cp -frv ./pkg/api/cmd ./cmd/api

# dump out the OpenAPI v3 schema for the Djangolang API
mkdir -p ./schema
go build -o ./cmd/api/api ./cmd/api

# TODO: we need this because custom endpoints don't come out in a plain old dump-openapi-json call
REDIS_URL=redis://localhost:6379 DJANGOLANG_API_ROOT=/api POSTGRES_DB=fred POSTGRES_PASSWORD=NoCI!11 ./cmd/api/api serve &
pid=$!
while ! curl http://localhost:7070/api/openapi.json >./schema/openapi.json; do
    sleep 0.1
done
kill -15 ${pid} || true >/dev/null 2>&1

# generate the client for use by the frontend
echo -e "\ngenerating typescript client..."
cd frontend
if [[ "${SKIP_UPDATE_FRONTEND}" != "1" ]]; then
    npm ci
fi
npm run openapi-typescript
npm run prettier
cd ..

# generate the client for use by Go code
if test -e pkg/api_client; then
    rm -frv pkg/api_client
fi
mkdir -p pkg/api_client
openapi-generator-cli generate -i schema/openapi.json -g go -o pkg/api_client --strict-spec true --package-name api_client
