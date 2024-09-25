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

docker build --platform=linux/amd64 -t kube-registry:5000/fred-api:latest -f ./docker/api/Dockerfile .
docker build --platform=linux/amd64 -t kube-registry:5000/fred-change-producer:latest -f ./docker/change-producer/Dockerfile .
docker build --platform=linux/amd64 -t kube-registry:5000/fred-job-executor:latest -f ./docker/job-executor/Dockerfile .
docker build --platform=linux/amd64 -t kube-registry:5000/fred-trigger-producer:latest -f ./docker/trigger-producer/Dockerfile .
docker build --platform=linux/amd64 -t kube-registry:5000/fred-frontend:latest -f ./docker/frontend/Dockerfile .

docker image push kube-registry:5000/fred-api:latest
docker image push kube-registry:5000/fred-change-producer:latest
docker image push kube-registry:5000/fred-job-executor:latest
docker image push kube-registry:5000/fred-trigger-producer:latest
docker image push kube-registry:5000/fred-frontend:latest
