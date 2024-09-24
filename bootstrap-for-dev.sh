#!/bin/bash

set -e

object_id=""
function do_request() {
    if ! response=$(curl -s -X POST "http://localhost:7070/api/${1}" -d "${2}" 2>&1); then
        echo "error: ${response}"
        return 1
    fi

    status=$(echo "${response}" | jq -r '.status')
    if [[ "${status}" != "200" && "${status}" != "201" ]]; then
        echo "error: $(echo "${response}" | jq)"
        return 1
    fi

    echo "info: $(echo "${response}" | jq)"

    # shellcheck disable=SC2002
    object_id=$(echo "${response}" | jq -r '.objects[0].id')
    echo "object_id: ${object_id}"

    echo ""
    return 0
}

echo ""

#
# repo
#

do_request "repositories" '[{"url": "https://github.com/initialed85/djangolang"}]'
repository_1_id="${object_id}"

#
# rules
#

do_request "rules" "[{\"branch_name\": \"main\", \"repository_id\": \"${repository_1_id}\"}]"
rule_1_id="${object_id}"

do_request "rules" "[{\"branch_name\": \"pr-*\", \"repository_id\": \"${repository_1_id}\"}]"
rule_2_id="${object_id}"

#
# build
#

raw_build_script=$(
    cat <<-EOM
#!/bin/bash

set -e

./build.sh
EOM
)
build_script=$(python3 -c "import json; print(json.dumps('''${raw_build_script}'''))")

do_request "tasks" "[{\"name\": \"build\", \"platform\": \"linux/amd64\", \"image\": \"initialed85/the-last-ci-image-you-will-ever-need:latest\", \"script\": ${build_script}}]"
task_1_id="${object_id}"

#
# test
#

raw_test_script=$(
    cat <<-EOM
#!/bin/bash

set -e

./test.sh
EOM
)
test_script=$(python3 -c "import json; print(json.dumps('''${raw_test_script}'''))")

do_request "tasks" "[{\"name\": \"test\", \"platform\": \"linux/amd64\", \"image\": \"initialed85/the-last-ci-image-you-will-ever-need:latest\", \"script\": ${test_script}}]"
task_2_id="${object_id}"

#
# job
#

do_request "jobs" "[{\"name\": \"djangolang-build-and-test-main\", \"rule_id\": \"${rule_1_id}\", \"build_task_id\": \"${task_1_id}\", \"test_task_id\": \"${task_2_id}\"}]"

do_request "jobs" "[{\"name\": \"djangolang-build-and-test-pr\", \"rule_id\": \"${rule_2_id}\", \"build_task_id\": \"${task_1_id}\", \"test_task_id\": \"${task_2_id}\"}]"

echo 'done.'
