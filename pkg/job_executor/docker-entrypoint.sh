#!/bin/bash

set -e

pushd "$(pwd)"

if [[ "${FRED_TASK_NAME}" == "" ]]; then
    echo "error: FRED_TASK_NAME empty or unset"
    exit 1
fi

env_vars_before="${FRED_TASK_NAME}-env-vars-before.txt"
env_vars_after="${FRED_TASK_NAME}-env-vars-after.txt"
task_entrypoint="${FRED_TASK_NAME}-task-entrypoint.sh"

function teardown() {
    popd

    env >"${env_vars_after}"
    echo -e "\nenv vars (after execution):\n"
    cat "${env_vars_after}"
}

trap teardown EXIT

env >"${env_vars_before}"
echo -e "\nenv vars (before execution):\n"
cat "${env_vars_before}"

if [[ "${FRED_REPOSITORY_URL}" == "" ]]; then
    echo "error: FRED_REPOSITORY_URL empty or unset"
    exit 1
fi

if [[ "${FRED_REPOSITORY_BRANCH_NAME}" == "" ]]; then
    echo "error: FRED_REPOSITORY_BRANCH_NAME empty or unset"
    exit 1
fi

if [[ "${FRED_REPOSITORY_FOLDER_NAME}" == "" ]]; then
    echo "error: FRED_REPOSITORY_FOLDER_NAME empty or unset"
    exit 1
fi

if [[ "${FRED_REPOSITORY_COMMIT_HASH}" == "" ]]; then
    echo "error: FRED_REPOSITORY_COMMIT_HASH empty or unset"
    exit 1
fi

if ! test -e "./${task_entrypoint}"; then
    echo "error: ${task_entrypoint} does not exist"
    exit 1
fi

echo -e "\ngit clone and git checkout:\n"

if ! test -e "${FRED_REPOSITORY_FOLDER_NAME}"; then
    git clone --depth=1 --branch "${FRED_REPOSITORY_BRANCH_NAME}" "${FRED_REPOSITORY_URL}"
    cd "${FRED_REPOSITORY_FOLDER_NAME}"
else
    cd "${FRED_REPOSITORY_FOLDER_NAME}"
    git reset --hard
    git clean -d -f -x .
fi

git checkout "${FRED_REPOSITORY_COMMIT_HASH}"

git log -n 1

echo -e "\n${task_entrypoint}:\n"

if ! ../${task_entrypoint}; then
    echo -e "error: ${task_entrypoint} failed\n"
    echo -e "done."
    exit 1
else
    echo -e "info: ${task_entrypoint} succeeded\n"
    echo -e "done."
fi
