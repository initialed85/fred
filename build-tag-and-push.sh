#!/bin/bash

set -e

# this block ensures we can invoke this script from anywhere and have it automatically change to this folder first
pushd "$(dirname -- "${BASH_SOURCE[0]}")" >/dev/null 2>&1
function teardown() {
    popd >/dev/null 2>&1 || true
}
trap teardown exit

# we need docker to build the images
if ! command -v docker >/dev/null 2>&1; then
    echo "error: docker not found"
    exit 1
fi

# TODO: hack workaround for intermittent builds on darwin aarch64
mkdir -p ./tmp
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./tmp/api -trimpath ./cmd/api
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./tmp/repository-syncer -trimpath ./cmd/repository_syncer
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./tmp/job-coordinator -trimpath ./cmd/job_coordinator
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./tmp/job-executor -trimpath ./cmd/job_executor

docker build --platform=linux/amd64 -t kube-registry:5000/fred-api:latest -f ./docker/api/Dockerfile .
docker build --platform=linux/amd64 -t kube-registry:5000/fred-repository-syncer:latest -f ./docker/repository-syncer/Dockerfile .
docker build --platform=linux/amd64 -t kube-registry:5000/fred-job-coordinator:latest -f ./docker/job-coordinator/Dockerfile .
docker build --platform=linux/amd64 -t kube-registry:5000/fred-job-executor:latest -f ./docker/job-executor/Dockerfile .
docker build --platform=linux/amd64 -t kube-registry:5000/fred-frontend:latest -f ./docker/frontend/Dockerfile .

docker image push kube-registry:5000/fred-api:latest
docker image push kube-registry:5000/fred-repository-syncer:latest
docker image push kube-registry:5000/fred-job-coordinator:latest
docker image push kube-registry:5000/fred-job-executor:latest
docker image push kube-registry:5000/fred-frontend:latest
