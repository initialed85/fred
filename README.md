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

- [DONE] Database schema (and therefore endpoints because Djangolang)
- [DONE] Change producer (syncs w/ repo)
- [WIP] Trigger producer (consume Changes, produce Triggers for Jobs)
  - It's working, but it doesn't do anything pipeline-y yet
- [WIP] Job executor (consume Triggers, run Tasks under an Execution)
  - It's pulling and running Docker images but there's a few bugs to iron out with ports and volumes (mostly around Docker-in-Docker)

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
```
