#!/bin/bash

set -e

echo -e "\nenv vars:\n"
env

if [[ "${REPOSITORY_URL}" == "" ]]; then
    echo "error: REPOSITORY_URL empty or unset"
    exit 1
fi

if [[ "${REPOSITORY_BRANCH_NAME}" == "" ]]; then
    echo "error: REPOSITORY_BRANCH_NAME empty or unset"
    exit 1
fi

if [[ "${REPOSITORY_FOLDER_NAME}" == "" ]]; then
    echo "error: REPOSITORY_FOLDER_NAME empty or unset"
    exit 1
fi

if [[ "${REPOSITORY_COMMIT_HASH}" == "" ]]; then
    echo "error: REPOSITORY_COMMIT_HASH empty or unset"
    exit 1
fi

if ! test -e "./task-entrypoint.sh"; then
    echo "error: task-entrypoint.sh does not exist"
    exit 1
fi

echo -e "\ngit clone and git checkout:\n"

if ! test -e "${REPOSITORY_FOLDER_NAME}"; then
    git clone --depth=1 --branch "${REPOSITORY_BRANCH_NAME}" "${REPOSITORY_URL}"
else
    git reset --hard

    git clean -d -f -x .
fi

cd "${REPOSITORY_FOLDER_NAME}"

git checkout "${REPOSITORY_COMMIT_HASH}"

git log -n 1

echo -e "\ntask-entrypoint.sh:\n"

if ! ../task-entrypoint.sh; then
    echo -e "error: task-entrypoint.sh failed\n"
    echo -e "done."
    exit 1
else
    echo -e "info: task-entrypoint.sh succeeded\n"
    echo -e "done."
    exit 0
fi
