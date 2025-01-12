# fred

Among the numerous CI system out there, there's Jenkins, and then there's Fred.

Fred's stubborn and opinionated and not very full-featured.

## Bugs

- Need to ensure binfmt is always there

## Status

Done / WIP at the top, TODOs are in priority order.

- [DONE] Database schema (and therefore endpoints because Djangolang)
- [DONE] Repository syncer (syncs w/ repo)
- [WIP] Job coordinator (consume Changes, produce stub Executions for Jobs)
  - [TODO] Add folder filter for rules
  - [WIP] Implement pipelines feature (split / merge / wait for other jobs etc)
    - [DONE] Wait for other jobs
- [DONE] Job executor (consume Triggers, run Tasks under an Execution)
  - [DONE] Fix Task failures not bubbling up to Execution failures
- [DONE] Use volumes for Repository (and other asset) reuse between the Tasks of an Execution
- [DONE] Get it deployed to Kubernetes
  - [DONE] Get to the bottom of this (Kubernetes Docker-in-Docker):
    - `Error response from daemon: failed to create task for container: failed to create shim task: OCI runtime create failed: runc create failed: unable to start container process: error during container init: error running hook #0: fork/exec /proc/251/exe: no such file or directory: unknown`
- [WIP] Have some sort of lean UI
- [WIP] SSH key authentication for Repositories
  - [DONE] Pass through SSH creds from the host's root user
- [WIP] Docker registry authentication for job executors
  - [DONE] Pass through Docker creds from the host's root user
- [DONE] Fix storage of odd bytes in logs (ref.: `ERROR: invalid byte sequence for encoding "UTF8"`); change to a
  `bytea` column for logs
- [DONE] Fix `api/outputs/{primaryKey}` somehow generated w/ `Execution` model (no idea- if one is wrong, they should
  all be wrong)
- [WIP] Fix perf issue relating to `depth` params + frontend (this might be Djangolang thing)
- [DONE] Fix jobs stuck in "Running" when tasks are all "Errored"
- [TODO] Fix anything to do with orphaned containers
- [TODO] Fix anything to do with outputs or executions stuck in "Running" (some sort of cleanup for hard kills)
- [DONE] Ensure job executor gracefully waits for Docker (for Kubernetes Docker-in-Docker)
- [TODO] Add some timestamps to the various states etc
- [DONE] Add timestamps to the log output
- [TODO] Have a streaming WebSocket for the logs
- [DONE] Make it clear when a job executor is pulling the CI image (show it in the logs or something)
- [TODO] Make it clear which node a job executor is running on
- [TODO] Carry any environment variables set during a Task execution between the Tasks of an Execution (don't override
  CI-set ones though)
- [TODO] Support for tags as well as branches
- [TODO] Username / password (or token) authentication for Repositories
- [TODO] Support for Repository webhooks (at least GitHub for now)
- [WIP] Be able to trigger jobs manually
  - [DONE] Be able to trigger a specific job for a commit
  - [DONE] Be able to trigger all relevant jobs for a commit
- [TODO] Be able to trigger jobs from other jobs

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
DJANGOLANG_API_ROOT=/api/ POSTGRES_DB=fred POSTGRES_PASSWORD=NoCI\!11 go run ./cmd/repository_syncer/

# shell 6
DJANGOLANG_API_ROOT=/api/ POSTGRES_DB=fred POSTGRES_PASSWORD=NoCI\!11 go run ./cmd/job_coordinator/

# shell 7
DJANGOLANG_API_ROOT=/api/ POSTGRES_DB=fred POSTGRES_PASSWORD=NoCI\!11 go run ./cmd/job_executor/

# shell 8
echo -e "$(curl 'http://localhost:7070/api/executions?created_at__desc=&limit=2&depth=2' | jq | sed 's/\\u001b/\\033/g')"
```
