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

if [[ "${ONLY}" == "api" ]] || [[ "${ONLY}" == "" ]]; then
    docker build --platform=linux/amd64 -t kube-registry:5000/fred-api:latest -f ./docker/api/Dockerfile .
fi

if [[ "${ONLY}" == "repository-syncer" ]] || [[ "${ONLY}" == "" ]]; then
    docker build --platform=linux/amd64 -t kube-registry:5000/fred-repository-syncer:latest -f ./docker/repository-syncer/Dockerfile .
fi

if [[ "${ONLY}" == "job-coordinator" ]] || [[ "${ONLY}" == "" ]]; then
    docker build --platform=linux/amd64 -t kube-registry:5000/fred-job-coordinator:latest -f ./docker/job-coordinator/Dockerfile .
fi

if [[ "${ONLY}" == "job-executor" ]] || [[ "${ONLY}" == "" ]]; then
    docker build --platform=linux/amd64 -t kube-registry:5000/fred-job-executor:latest -f ./docker/job-executor/Dockerfile .
fi

if [[ "${ONLY}" == "frontend" ]] || [[ "${ONLY}" == "" ]]; then
    docker build --platform=linux/amd64 -t kube-registry:5000/fred-frontend:latest -f ./docker/frontend/Dockerfile .
fi

if [[ "${ONLY}" == "api" ]] || [[ "${ONLY}" == "" ]]; then
    docker image push kube-registry:5000/fred-api:latest
fi

if [[ "${ONLY}" == "repository-syncer" ]] || [[ "${ONLY}" == "" ]]; then
    docker image push kube-registry:5000/fred-repository-syncer:latest
fi

if [[ "${ONLY}" == "job-coordinator" ]] || [[ "${ONLY}" == "" ]]; then
    docker image push kube-registry:5000/fred-job-coordinator:latest
fi

if [[ "${ONLY}" == "job-executor" ]] || [[ "${ONLY}" == "" ]]; then
    docker image push kube-registry:5000/fred-job-executor:latest
fi

if [[ "${ONLY}" == "frontend" ]] || [[ "${ONLY}" == "" ]]; then
    docker image push kube-registry:5000/fred-frontend:latest
fi
