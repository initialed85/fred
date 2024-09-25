# fred

Among the numerous CI system out there, there's Jenkins, and then there's Fred.

Fred's stubborn and opinionated and not very full-featured.

## Concepts

My thinking is that we can do everything of value with the following phases:

- Build
- Test
- Publish
- Deploy
- Validate

On the configuration side, we'll have these definitions:

- Repository
  - Contains info about a code repository to watch
- Rule
  - Refers to a Repository
  - Contains info about which branches and folders to watch
- Job
  - Refers to a Rule
  - Refers to a number of Tasks (Build, Test, Publish, Deploy, Validate)
- Task
  - Contains a script to do something

Probably we'll throw those in some kind of folder structure e.g. `.fred` in the repo.

When things happen, we'll have these events:

- Change
  - Refers to a Repository
  - Contains info about a commit
- Trigger
  - Refers to a Rule
  - Contains info about which branch / folders have changed
- Execution
  - Refers to a Job
  - Refers to a number of Outputs (Build, Test, Publish, Deploy, Validate)
- Output
  - Refers to a Task
  - Contains the output of a script that did something

## Loose ideas

- Try to do everything with environment variables
  - Both as inputs and as outputs

## Status

Done / WIP at the top, TODOs are in priority order.

- [DONE] Database schema (and therefore endpoints because Djangolang)
- [DONE] Change producer (syncs w/ repo)
- [WIP] Trigger producer (consume Changes, produce Triggers for Jobs)
  - [TODO] Add folder filter for rules
  - [TODO] Implement pipelines feature (split / merge / wait for other jobs etc)
- [DONE] Job executor (consume Triggers, run Tasks under an Execution)
  - [DONE] Fix Task failures not bubbling up to Execution failures
- [DONE] Use volumes for Repository (and other asset) reuse between the Tasks of an Execution
- [WIP] Have some sort of lean UI
- [TODO] Fix `api/outputs/{primaryKey}` somehow generated w/ `Execution` model (no idea- if one is wrong, they should all be wrong)
- [TODO] Fix perf issue relating to `depth` params + frontend (this might be Djangolang thing)
- [TODO] Fix jobs stuck in "Running" when tasks are all "Errored"
- [TODO] Add some timestamps to the various states etc
- [TODO] Make it clear when a job executor is pulling the CI image (show it in the logs or something)
- [TODO] Make it clear which node a job executor is running on
- [TODO] Carry any environment variables set during a Task execution between the Tasks of an Execution (don't override CI-set ones though)
- [TODO] SSH key authentication for Repositories
- [TODO] Get it deployed to Kubernetes
- [TODO] Support for tags as well as branches
- [TODO] Username / password (or token) authentication for Repositories
- [TODO] Support for Repository webhooks (at least GitHub for now)
- [TODO] Be able to trigger jobs manually

## Dev notes

```shell
# shell 1
./run-env.sh

# shell 2
./run-for-dev.sh

# shell 3
websocat ws://localhost:7070/api/__stream --exit-on-eof | jq

# shell 4
./bootstrap-for-dev.sh

# shell 5
DJANGOLANG_API_ROOT=/api/ POSTGRES_DB=fred POSTGRES_PASSWORD=NoCI\!11 go run ./cmd/change_producer/

# shell 6
DJANGOLANG_API_ROOT=/api/ POSTGRES_DB=fred POSTGRES_PASSWORD=NoCI\!11 go run ./cmd/trigger_producer

# shell 7 (worker 1)
DJANGOLANG_API_ROOT=/api/ POSTGRES_DB=fred POSTGRES_PASSWORD=NoCI\!11 go run ./cmd/job_executor

# shell 8 (worker 2)
DJANGOLANG_API_ROOT=/api/ POSTGRES_DB=fred POSTGRES_PASSWORD=NoCI\!11 go run ./cmd/job_executor

# shell 9
echo -e "$(curl 'http://localhost:7070/api/executions?created_at__desc=&limit=2&depth=2' | jq | sed 's/\\u001b/\\033/g')"
```
