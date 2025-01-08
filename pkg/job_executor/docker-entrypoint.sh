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

env >"${env_vars_before}"
echo -e "\nenv vars:\n"
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

if ! command -v "git"; then
    echo "error: git command not found in FRED_TASK_IMAGE=${FRED_TASK_IMAGE}; cannot continue"
    exit 1
fi

if ! command -v "ssh"; then
    echo "error: ssh command not found in FRED_TASK_IMAGE=${FRED_TASK_IMAGE}; cannot continue"
    exit 1
fi

if test -e "/root/.ssh/id_rsa"; then
    chmod_id_rsa=$(stat -c '%a' /root/.ssh/id_rsa)
    if [[ "${chmod_id_rsa}" != "400" ]] && [[ "${chmod_id_rsa}" != "600" ]]; then
        echo "error: chmod for /root/.ssh/id_rsa is ${chmod_id_rsa} not 400 or 600; cannot continue"
        exit 1
    fi
fi

sleep 5

echo -e "\ngit clone and git checkout:\n"

git config --global pack.threads "1"

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

set +e
bash -c "../${task_entrypoint}"
retval=${?}
set -e

popd
env >"${env_vars_after}"

if [[ "${retval}" != "0" ]]; then
    echo -e "error: ${task_entrypoint} failed\n"
    echo -e "done."
    sleep 5
    exit "${retval}"
fi

echo -e "info: ${task_entrypoint} succeeded\n"
echo -e "done."
sleep 5
