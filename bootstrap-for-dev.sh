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

do_request "repositories" '[{"url": "https://github.com/initialed85/camry"}]'
repository_2_id="${object_id}"

do_request "repositories" '[{"url": "https://github.com/initialed85/eds-game-for-ftp-game-jam-2022"}]'
repository_3_id="${object_id}"

do_request "repositories" '[{"url": "https://github.com/initialed85/quake-websocket-proxy"}]'
repository_4_id="${object_id}"

do_request "repositories" '[{"url": "https://github.com/initialed85/mqtt_things"}]'
repository_5_id="${object_id}"

do_request "repositories" '[{"url": "https://github.com/initialed85/game-of-life"}]'
repository_6_id="${object_id}"

do_request "repositories" '[{"url": "https://github.com/initialed85/the-last-ci-image-you-will-ever-need"}]'
repository_7_id="${object_id}"

#
# rules
#

do_request "rules" "[{\"branch_name\": \"main\", \"repository_id\": \"${repository_1_id}\"}]"
rule_1_id="${object_id}"

do_request "rules" "[{\"branch_name\": \"main\", \"repository_id\": \"${repository_2_id}\"}]"
rule_2_id="${object_id}"

do_request "rules" "[{\"branch_name\": \"master\", \"repository_id\": \"${repository_3_id}\"}]"
rule_3_id="${object_id}"

do_request "rules" "[{\"branch_name\": \"master\", \"repository_id\": \"${repository_4_id}\"}]"
rule_4_id="${object_id}"

do_request "rules" "[{\"branch_name\": \"master\", \"repository_id\": \"${repository_5_id}\"}]"
rule_5_id="${object_id}"

do_request "rules" "[{\"branch_name\": \"main\", \"repository_id\": \"${repository_6_id}\"}]"
rule_6_id="${object_id}"

do_request "rules" "[{\"branch_name\": \"master\", \"repository_id\": \"${repository_7_id}\"}]"
rule_7_id="${object_id}"

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
# publish
#

raw_publish_script=$(
    cat <<-EOM
#!/bin/bash

set -e

./build-tag-and-push.sh
EOM
)
publish_script=$(python3 -c "import json; print(json.dumps('''${raw_publish_script}'''))")

do_request "tasks" "[{\"name\": \"publish\", \"platform\": \"linux/amd64\", \"image\": \"initialed85/the-last-ci-image-you-will-ever-need:latest\", \"script\": ${publish_script}}]"
task_3_id="${object_id}"

#
# job
#

# https://github.com/initialed85/djangolang
# https://github.com/initialed85/camry
# https://github.com/initialed85/eds-game-for-ftp-game-jam-2022
# https://github.com/initialed85/quake-websocket-proxy
# https://github.com/initialed85/mqtt_things
# https://github.com/initialed85/game-of-life
# https://github.com/initialed85/the-last-ci-image-you-will-ever-need

do_request "jobs" "[{\"name\": \"djangolang-main\", \"rule_id\": \"${rule_1_id}\", \"build_task_id\": \"${task_1_id}\", \"test_task_id\": \"${task_2_id}\"}]"

do_request "jobs" "[{\"name\": \"camry-main\", \"rule_id\": \"${rule_2_id}\", \"publish_task_id\": \"${task_3_id}\"}]"

do_request "jobs" "[{\"name\": \"eds-game-for-ftp-game-jam-2022-main\", \"rule_id\": \"${rule_3_id}\", \"publish_task_id\": \"${task_3_id}\"}]"

do_request "jobs" "[{\"name\": \"quake-websocket-proxy-main\", \"rule_id\": \"${rule_4_id}\", \"publish_task_id\": \"${task_3_id}\"}]"

do_request "jobs" "[{\"name\": \"mqtt-things-main\", \"rule_id\": \"${rule_5_id}\", \"publish_task_id\": \"${task_3_id}\"}]"

do_request "jobs" "[{\"name\": \"game-of-life-main\", \"rule_id\": \"${rule_6_id}\", \"publish_task_id\": \"${task_3_id}\"}]"

do_request "jobs" "[{\"name\": \"the-last-ci-image-you-will-ever-need-main\", \"rule_id\": \"${rule_7_id}\", \"publish_task_id\": \"${task_3_id}\"}]"

echo 'done.'
