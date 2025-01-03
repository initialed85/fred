#!/bin/bash

set -e -x

object_id=""
function do_request() {
    if ! response=$(curl -s -X POST "https://fred.initialed85.cc/api/${1}" -d "${2}" 2>&1); then
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
# ----
#

do_request "repositories" '[{"url": "https://github.com/initialed85/djangolang"}]'
repository_id="${object_id}"

do_request "rules" "[{\"branch_name\": \"main\", \"repository_id\": \"${repository_id}\"}]"
rule_id="${object_id}"

do_request "jobs" "[{\"name\": \"djangolang-main\", \"rule_id\": \"${rule_id}\"}]"
job_id="${object_id}"

raw_build_script=$(
    cat <<-EOM
#!/bin/bash

set -e

./build.sh
EOM
)
build_script=$(python3 -c "import json; print(json.dumps('''${raw_build_script}'''))")

do_request "tasks" "[{\"index\": 0, \"name\": \"build\", \"platform\": \"linux/amd64\", \"image\": \"initialed85/the-last-ci-image-you-will-ever-need:latest\", \"script\": ${build_script}, \"job_id\": \"${job_id}\"}]"
_="${object_id}"

raw_test_script=$(
    cat <<-EOM
#!/bin/bash

set -e

./test.sh
EOM
)
test_script=$(python3 -c "import json; print(json.dumps('''${raw_test_script}'''))")

do_request "tasks" "[{\"index\": 1, \"name\": \"test\", \"platform\": \"linux/amd64\", \"image\": \"initialed85/the-last-ci-image-you-will-ever-need:latest\", \"script\": ${test_script}, \"job_id\": \"${job_id}\"}]"
_="${object_id}"

echo 'done.'

#
# ----
#

# do_request "repositories" '[{"url": "https://github.com/initialed85/camry"}]'
# repository_id="${object_id}"

# do_request "rules" "[{\"branch_name\": \"main\", \"repository_id\": \"${repository_id}\"}]"
# rule_id="${object_id}"

# do_request "jobs" "[{\"name\": \"camry-main\", \"rule_id\": \"${rule_id}\"}]"
# job_id="${object_id}"

# raw_publish_script=$(
#     cat <<-EOM
# #!/bin/bash

# set -e

# ./build-tag-and-push.sh
# EOM
# )
# publish_script=$(python3 -c "import json; print(json.dumps('''${raw_publish_script}'''))")

# do_request "tasks" "[{\"index\": 0, \"name\": \"publish\", \"platform\": \"linux/amd64\", \"image\": \"initialed85/the-last-ci-image-you-will-ever-need:latest\", \"script\": ${publish_script}, \"job_id\": \"${job_id}\"}]"
# _="${object_id}"

# echo 'done.'

#
# ----
#

do_request "repositories" '[{"url": "https://github.com/initialed85/game-of-life"}]'
repository_id="${object_id}"

do_request "rules" "[{\"branch_name\": \"main\", \"repository_id\": \"${repository_id}\"}]"
rule_id="${object_id}"

do_request "jobs" "[{\"name\": \"game-of-life-main\", \"rule_id\": \"${rule_id}\"}]"
job_id="${object_id}"

raw_publish_script=$(
    cat <<-EOM
#!/bin/bash

set -e

./build-tag-and-push.sh
EOM
)
publish_script=$(python3 -c "import json; print(json.dumps('''${raw_publish_script}'''))")

do_request "tasks" "[{\"index\": 0, \"name\": \"publish\", \"platform\": \"linux/amd64\", \"image\": \"initialed85/the-last-ci-image-you-will-ever-need:latest\", \"script\": ${publish_script}, \"job_id\": \"${job_id}\"}]"
_="${object_id}"

echo 'done.'

#
# ----
#

do_request "repositories" '[{"url": "https://github.com/initialed85/eds-game-for-ftp-game-jam-2022"}]'
repository_id="${object_id}"

do_request "rules" "[{\"branch_name\": \"main\", \"repository_id\": \"${repository_id}\"}]"
rule_id="${object_id}"

do_request "jobs" "[{\"name\": \"eds-game-for-ftp-game-jam-2022-main\", \"rule_id\": \"${rule_id}\"}]"
job_id="${object_id}"

raw_publish_script=$(
    cat <<-EOM
#!/bin/bash

set -e

./build-tag-and-push.sh
EOM
)
publish_script=$(python3 -c "import json; print(json.dumps('''${raw_publish_script}'''))")

do_request "tasks" "[{\"index\": 0, \"name\": \"publish\", \"platform\": \"linux/amd64\", \"image\": \"initialed85/the-last-ci-image-you-will-ever-need:latest\", \"script\": ${publish_script}, \"job_id\": \"${job_id}\"}]"
_="${object_id}"

echo 'done.'

#
# ----
#

do_request "repositories" '[{"url": "https://github.com/initialed85/quake-websocket-proxy"}]'
repository_id="${object_id}"

do_request "rules" "[{\"branch_name\": \"main\", \"repository_id\": \"${repository_id}\"}]"
rule_id="${object_id}"

do_request "jobs" "[{\"name\": \"quake-websocket-proxy-main\", \"rule_id\": \"${rule_id}\"}]"
job_id="${object_id}"

raw_publish_script=$(
    cat <<-EOM
#!/bin/bash

set -e

./build-tag-and-push.sh
EOM
)
publish_script=$(python3 -c "import json; print(json.dumps('''${raw_publish_script}'''))")

do_request "tasks" "[{\"index\": 0, \"name\": \"publish\", \"platform\": \"linux/amd64\", \"image\": \"initialed85/the-last-ci-image-you-will-ever-need:latest\", \"script\": ${publish_script}, \"job_id\": \"${job_id}\"}]"
_="${object_id}"

echo 'done.'

#
# ----
#

do_request "repositories" '[{"url": "https://github.com/initialed85/mqtt_things"}]'
repository_id="${object_id}"

do_request "rules" "[{\"branch_name\": \"main\", \"repository_id\": \"${repository_id}\"}]"
rule_id="${object_id}"

do_request "jobs" "[{\"name\": \"mqtt-things-main\", \"rule_id\": \"${rule_id}\"}]"
job_id="${object_id}"

raw_publish_script=$(
    cat <<-EOM
#!/bin/bash

set -e

./build-tag-and-push.sh
EOM
)
publish_script=$(python3 -c "import json; print(json.dumps('''${raw_publish_script}'''))")

do_request "tasks" "[{\"index\": 0, \"name\": \"publish\", \"platform\": \"linux/amd64\", \"image\": \"initialed85/the-last-ci-image-you-will-ever-need:latest\", \"script\": ${publish_script}, \"job_id\": \"${job_id}\"}]"
_="${object_id}"

echo 'done.'

#
# ----
#

do_request "repositories" '[{"url": "https://github.com/initialed85/the-last-ci-image-you-will-ever-need"}]'
repository_id="${object_id}"

do_request "rules" "[{\"branch_name\": \"main\", \"repository_id\": \"${repository_id}\"}]"
rule_id="${object_id}"

do_request "jobs" "[{\"name\": \"the-last-ci-image-you-will-ever-need-main\", \"rule_id\": \"${rule_id}\"}]"
job_id="${object_id}"

raw_publish_script=$(
    cat <<-EOM
#!/bin/bash

set -e

./build-tag-and-push.sh
EOM
)
publish_script=$(python3 -c "import json; print(json.dumps('''${raw_publish_script}'''))")

do_request "tasks" "[{\"index\": 0, \"name\": \"publish\", \"platform\": \"linux/amd64\", \"image\": \"initialed85/the-last-ci-image-you-will-ever-need:latest\", \"script\": ${publish_script}, \"job_id\": \"${job_id}\"}]"
_="${object_id}"

echo 'done.'
