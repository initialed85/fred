package job_executor

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/docker/cli/cli/streams"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/google/uuid"
	"github.com/initialed85/djangolang/pkg/helpers"
	"github.com/initialed85/djangolang/pkg/query"
	"github.com/initialed85/fred/internal"
	"github.com/initialed85/fred/pkg/api"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"

	_ "embed"
)

//go:embed docker-entrypoint.sh
var dockerEntrypointScript string

var log = helpers.GetLogger("job_executor")

func formatUUID(u uuid.UUID) string {
	return strings.ReplaceAll(u.String(), "-", "")
}

func ClaimTriggerForJobExecutor(ctx context.Context, tx pgx.Tx, claimDuration time.Duration) (*api.Trigger, error) {
	now := time.Now().UTC()

	claimUntil := now.Add(claimDuration)

	if claimUntil.Sub(now) <= 0 {
		return nil, fmt.Errorf("claim_duration_seconds too short; must result in a claim that expires in the future")
	}

	trigger := &api.Trigger{}

	err := trigger.AdvisoryLockWithRetries(ctx, tx, internal.TriggerTableJobExecutorClaimedUntilAdvisoryLockID, claimDuration+(time.Second*2), time.Second*1)
	if err != nil {
		return nil, err
	}

	triggers, _, _, _, _, err := api.SelectTriggers(
		ctx,
		tx,
		fmt.Sprintf(
			"%s < now() AND %s IS null",
			api.TriggerTableJobExecutorClaimedUntilColumn,
			api.TriggerTableHandledAtColumn,
		),
		helpers.Ptr(fmt.Sprintf(
			"%v ASC",
			api.TriggerTableJobExecutorClaimedUntilColumn,
		)),
		helpers.Ptr(1),
		nil,
	)
	if err != nil {
		return nil, err
	}

	if len(triggers) == 0 {
		return nil, nil
	}

	if len(triggers) > 1 {
		return nil, fmt.Errorf("wanted exactly 1 unclaimed trigger, got %d", len(triggers))
	}

	trigger = triggers[0]

	trigger.JobExecutorClaimedUntil = claimUntil

	err = trigger.Update(ctx, tx, false)
	if err != nil {
		return nil, err
	}

	err = trigger.Reload(ctx, tx)
	if err != nil {
		return nil, err
	}

	return trigger, nil
}

func ClaimJobForJobExecutor(ctx context.Context, tx pgx.Tx, trigger *api.Trigger, claimDuration time.Duration) (*api.Job, error) {
	now := time.Now().UTC()

	claimUntil := now.Add(claimDuration)

	if claimUntil.Sub(now) <= 0 {
		return nil, fmt.Errorf("claim_duration_seconds too short; must result in a claim that expires in the future")
	}

	job := &api.Job{}

	err := job.AdvisoryLockWithRetries(ctx, tx, internal.JobExecutorAdvisoryLockID, claimDuration+(time.Second*2), time.Second*1)
	if err != nil {
		return nil, err
	}

	ctx = query.WithMaxDepth(ctx, helpers.Ptr(2))

	jobs, _, _, _, _, err := api.SelectJobs(
		ctx,
		tx,
		fmt.Sprintf(
			"%s < now() AND %s = $$??",
			api.JobTableJobExecutorClaimedUntilColumn,
			api.JobTableRuleIDColumn,
		),
		helpers.Ptr(fmt.Sprintf(
			"%v ASC",
			api.JobTableJobExecutorClaimedUntilColumn,
		)),
		helpers.Ptr(1),
		nil,
		trigger.RuleID,
	)
	if err != nil {
		return nil, err
	}

	if len(jobs) == 0 {
		return nil, nil
	}

	if len(jobs) > 1 {
		return nil, fmt.Errorf("wanted exactly 1 unclaimed trigger, got %d", len(jobs))
	}

	job = jobs[0]

	job.JobExecutorClaimedUntil = claimUntil

	err = job.Update(ctx, tx, false)
	if err != nil {
		return nil, err
	}

	return job, nil
}

func Run() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	tempPath := filepath.Join(wd, "temp")

	err = os.MkdirAll(tempPath, 0o777)
	if err != nil {
		return err
	}

	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel((context.Background()))
	defer cancel()

	for {
		err = func() error {
			pingCtx, pingCancel := context.WithTimeout(ctx, time.Second*1)
			defer pingCancel()

			pingResponse, err := apiClient.Ping(pingCtx)
			if err != nil {
				return err
			}
			_ = pingResponse

			return nil
		}()
		if err != nil {
			log.Printf("warning: Docker daemon not available / not ready: %s; retrying...", err.Error())
			time.Sleep(time.Second * 1)
			continue
		}

		break
	}

	apiClient.NegotiateAPIVersion(ctx)

	defer func() {
		_ = apiClient.Close()
	}()

	return internal.Run(
		log,
		func(ctx context.Context, db *pgxpool.Pool) error {
			tx, err := db.Begin(ctx)
			if err != nil {
				return err
			}

			defer func() {
				_ = tx.Rollback(ctx)
			}()

			//
			// first we claim a trigger, but just long enough to claim a job
			//

			trigger, err := ClaimTriggerForJobExecutor(ctx, tx, time.Second*60)
			if err != nil {
				return err
			}

			if trigger == nil {
				return nil
			}

			log.Printf("claimed %s", internal.GetTriggerSummary(trigger))

			//
			// then we claim a job, and we'll keep it claimed until we've finished executing it
			//

			job, err := ClaimJobForJobExecutor(ctx, tx, trigger, time.Second*60)
			if err != nil {
				return err
			}

			if job == nil {
				log.Printf("warning: no jobs for %s; skipping", internal.GetTriggerSummary(trigger))
				return nil
			}

			log.Printf("claimed %s", internal.GetJobSummary(job))

			var execution *api.Execution
			outputs := make([]*api.Output, 0)

			defer func() {
				if execution != nil && execution.Status != internal.ExecutionOrTaskStatusRunning {
					return
				}

				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()

				tx, err := db.Begin(ctx)
				if err != nil {
					return
				}

				defer func() {
					_ = tx.Rollback(ctx)
				}()

				execution.Status = internal.ExecutionOrTaskStatusErrored
				execution.EndedAt = helpers.Ptr(time.Now().UTC())
				err = execution.Update(ctx, tx, false)
				if err != nil {
					return
				}

				for _, output := range outputs {
					output.Status = internal.ExecutionOrTaskStatusErrored
					output.EndedAt = helpers.Ptr(time.Now().UTC())
					err = output.Update(ctx, tx, false)
					if err != nil {
						return
					}
				}

				err = tx.Commit(ctx)
				if err != nil {
					return
				}
			}()

			now := time.Now().UTC()
			trigger.JobExecutorClaimedUntil = now
			trigger.HandledAt = &now

			ctx = query.WithMaxDepth(ctx, helpers.Ptr(3))

			err = trigger.Update(ctx, tx, false)
			if err != nil {
				return err
			}

			err = tx.Commit(ctx)
			if err != nil {
				return err
			}

			ctx, cancel := context.WithCancel(ctx)
			defer cancel()

			go func() {
				t := time.NewTicker(time.Second)

				defer func() {
					err := func() error {
						ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
						defer cancel()

						tx, err := db.Begin(ctx)
						if err != nil {
							return err
						}

						defer func() {
							_ = tx.Rollback(ctx)
						}()

						job.JobExecutorClaimedUntil = time.Now().UTC()
						err = job.Update(ctx, tx, false)
						if err != nil {
							return err
						}

						err = tx.Commit(ctx)
						if err != nil {
							return err
						}

						return nil
					}()
					if err != nil {
						log.Printf("warning: job claim release failed: %s", err.Error())
					}
				}()

				for {
					select {
					case <-ctx.Done():
						return
					case <-t.C:
					}

					err := func() error {
						tx, err := db.Begin(ctx)
						if err != nil {
							return err
						}

						defer func() {
							_ = tx.Rollback(ctx)
						}()

						job.JobExecutorClaimedUntil = time.Now().UTC().Add(time.Second * 60)
						err = job.Update(ctx, tx, false)
						if err != nil {
							return err
						}

						err = tx.Commit(ctx)
						if err != nil {
							return err
						}

						return nil
					}()
					if err != nil {
						log.Printf("warning: job claim refresh failed: %s", err.Error())
						continue
					}
				}
			}()

			doTask := func(output *api.Output, execution *api.Execution) error {
				task := output.TaskIDObject

				ctx, cancel := context.WithCancel(ctx)
				defer cancel()

				defer func() {
					if output.Status == internal.ExecutionOrTaskStatusErrored || output.Status == internal.ExecutionOrTaskStatusFailed {
						tx, err := db.Begin(ctx)
						if err != nil {
							log.Printf("warning: %s", err.Error())
							return
						}

						defer func() {
							_ = tx.Rollback(ctx)
						}()

						if output.Status == internal.ExecutionOrTaskStatusErrored {
							execution.Status = internal.ExecutionOrTaskStatusErrored
						} else if output.Status == internal.ExecutionOrTaskStatusFailed {
							execution.Status = internal.ExecutionOrTaskStatusFailed
						}

						err = tx.Commit(ctx)
						if err != nil {
							log.Printf("warning: %s", err.Error())
							return
						}
					}
				}()

				if execution.Status == internal.ExecutionOrTaskStatusFailing || execution.Status == internal.ExecutionOrTaskStatusErroring {
					output.Status = internal.ExecutionOrTaskStatusSkipped
					return nil
				}

				tx, err := db.Begin(ctx)
				if err != nil {
					return err
				}

				defer func() {
					_ = tx.Rollback(ctx)
				}()

				output.Status = internal.ExecutionOrTaskStatusRunning

				output.LogIDObject = &api.Log{
					OutputID: output.ID,
				}

				err = output.LogIDObject.Insert(ctx, tx, false, false)
				if err != nil {
					return err
				}

				err = tx.Commit(ctx)
				if err != nil {
					return err
				}

				volumeName := fmt.Sprintf("fred-%s-%s", trigger.ChangeIDObject.CommitHash, formatUUID(execution.ID))
				containerName := fmt.Sprintf("%s-%s", volumeName, task.Name)

				if execution.Status == internal.ExecutionOrTaskStatusFailing || execution.Status == internal.ExecutionOrTaskStatusErroring {
					log.Printf("%s skipped (due to failing / erroring parent execution)", containerName)
					return nil
				}

				var exitStatus *int
				logBuffer := make([]byte, 0)

				updateOutput := func(givenErr error) error {
					tx, err := db.Begin(ctx)
					if err != nil {
						return err
					}

					defer func() {
						_ = tx.Rollback(ctx)
					}()

					if logBuffer != nil {
						output.LogIDObject.Buffer = logBuffer
					}

					if givenErr != nil {
						output.Error = helpers.Ptr(givenErr.Error())
					}

					if exitStatus != nil {
						output.ExitStatus = int64(*exitStatus)
					}

					if output.Error != nil {
						output.Status = internal.ExecutionOrTaskStatusErrored
					} else if exitStatus != nil {
						if *exitStatus != 0 {
							output.Status = internal.ExecutionOrTaskStatusFailed
						} else {
							output.Status = internal.ExecutionOrTaskStatusSucceeded
						}
					}

					err = output.Update(ctx, tx, false)
					if err != nil {
						return err
					}

					err = tx.Commit(ctx)
					if err != nil {
						return err
					}

					return nil
				}

				tempDir := filepath.Join(tempPath, volumeName)
				err = os.MkdirAll(tempDir, 0o777)
				if err != nil {
					_ = updateOutput(err)
					return err
				}

				repositoryUrl := trigger.ChangeIDObject.RepositoryIDObject.URL
				repositoryBranchName := trigger.ChangeIDObject.BranchName
				parts := strings.Split(repositoryUrl, "/")
				repositoryFolderName := parts[len(parts)-1]
				repositoryCommitHash := trigger.ChangeIDObject.CommitHash

				platformParts := strings.Split(task.Platform, "/")
				if len(platformParts) != 2 {
					err := fmt.Errorf("platform %#+v doesn't seem sane (want [os]/[arch])", task.Platform)
					_ = updateOutput(err)
					return err
				}
				platformOS := platformParts[0]
				platformArch := platformParts[1]

				builtinEnvVars := []string{
					fmt.Sprintf("FRED_REPOSITORY_ID=%s", trigger.RepositoryID.String()),
					fmt.Sprintf("FRED_CHANGE_ID=%s", trigger.ChangeID.String()),
					fmt.Sprintf("FRED_RULE_ID=%s", trigger.RuleID.String()),
					fmt.Sprintf("FRED_JOB_ID=%s", job.ID.String()),
					fmt.Sprintf("FRED_JOB_NAME=%s", job.Name),
					fmt.Sprintf("FRED_TASK_ID=%s", task.ID.String()),
					fmt.Sprintf("FRED_TASK_NAME=%s", task.Name),
					fmt.Sprintf("FRED_TASK_PLATFORM=%s", task.Platform),
					fmt.Sprintf("FRED_TASK_PLATFORM_OS=%s", platformOS),
					fmt.Sprintf("FRED_TASK_PLATFORM_ARCH=%s", platformArch),
					fmt.Sprintf("FRED_TASK_IMAGE=%s", task.Image),
					fmt.Sprintf("FRED_TRIGGER_ID=%s", trigger.ID.String()),
					fmt.Sprintf("FRED_EXECUTION_ID=%s", execution.ID.String()),
					fmt.Sprintf("FRED_OUTPUT_ID=%s", output.ID.String()),
					fmt.Sprintf("FRED_REPOSITORY_URL=%s", repositoryUrl),
					fmt.Sprintf("FRED_REPOSITORY_BRANCH_NAME=%s", repositoryBranchName),
					fmt.Sprintf("FRED_REPOSITORY_FOLDER_NAME=%s", repositoryFolderName),
					fmt.Sprintf("FRED_REPOSITORY_COMMIT_HASH=%s", repositoryCommitHash),
					"CI=true",
					"COMPOSE_DOCKER_CLI_BUILD=1",
					"DOCKER_BUILDKIT=1",
					fmt.Sprintf("COMPOSE_PROJECT_NAME=%s", containerName),
					// "GIT_TERMINAL_PROMPT=1",
					// "GIT_TRACE=1",
					// "GIT_CURL_VERBOSE=1",
					"GIT_SSH_COMMAND=ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -o BatchMode=yes",
					fmt.Sprintf("DOCKER_HOST=%s", os.Getenv("DOCKER_HOST")),
				}

				logAndUpdateOutput := func(msg string, extra ...any) {
					log.Printf(msg, extra...)
					logBuffer = append(logBuffer, []byte(fmt.Sprintf(msg+"\n", extra...))...)
					_ = updateOutput(nil)
				}

				builtinEnvVarsFilePath := filepath.Join(tempDir, fmt.Sprintf("%s-builtin-env-vars.txt", task.Name))
				err = os.WriteFile(builtinEnvVarsFilePath, []byte(strings.Join(builtinEnvVars, "\n")+"\n"), 0o777)
				if err != nil {
					_ = updateOutput(err)
					return err
				}
				logAndUpdateOutput("prepared %s", builtinEnvVarsFilePath)

				dockerEntrypointOutsideFilePath := filepath.Join(tempDir, fmt.Sprintf("%s-docker-entrypoint.sh", task.Name))
				err = os.WriteFile(dockerEntrypointOutsideFilePath, []byte(dockerEntrypointScript), 0o777)
				if err != nil {
					_ = updateOutput(err)
					return err
				}
				logAndUpdateOutput("prepared %s", dockerEntrypointOutsideFilePath)

				taskEntrypointOutsideFilePath := filepath.Join(tempDir, fmt.Sprintf("%s-task-entrypoint.sh", task.Name))
				err = os.WriteFile(taskEntrypointOutsideFilePath, []byte(task.Script), 0o777)
				if err != nil {
					_ = updateOutput(err)
					return err
				}
				logAndUpdateOutput("prepared %s", taskEntrypointOutsideFilePath)

				logAndUpdateOutput("pulling %s", task.Image)

				imagePullReader, err := apiClient.ImagePull(
					ctx,
					task.Image,
					image.PullOptions{
						Platform: task.Platform,
					},
				)
				if err != nil {
					_ = updateOutput(err)
					return err
				}

				defer func() {
					_ = imagePullReader.Close()
				}()

				logBufferWriter := bytes.NewBuffer(logBuffer)

				_ = jsonmessage.DisplayJSONMessagesToStream(
					imagePullReader,
					streams.NewOut(logBufferWriter),
					func(j jsonmessage.JSONMessage) {
						_ = updateOutput(nil)
					},
				)
				defer func() {
					_, _ = io.ReadAll(imagePullReader)
				}()

				logAndUpdateOutput("creating %s", containerName)

				envVars := make([]string, 0)
				envVars = append(envVars, builtinEnvVars...)

				// TODO
				dockerSockPath := "/var/run/docker.sock"
				if os.Getenv("DOCKER_HOST") != "" && strings.Contains(os.Getenv("DOCKER_HOST"), "/var/run/docker/docker.sock") {
					dockerSockPath = "/var/run/docker/docker.sock"
				}

				containerCreateResponse, err := apiClient.ContainerCreate(
					ctx,
					&container.Config{
						Image:        task.Image,
						Env:          envVars,
						WorkingDir:   tempDir,
						Entrypoint:   []string{fmt.Sprintf("/%s/%s-docker-entrypoint.sh", tempDir, task.Name)},
						Cmd:          []string{},
						StopTimeout:  helpers.Ptr(1),
						Tty:          false,
						OpenStdin:    false,
						AttachStdin:  false,
						AttachStdout: false,
						AttachStderr: false,
						User:         "root",
					},
					&container.HostConfig{
						RestartPolicy: container.RestartPolicy{
							Name: container.RestartPolicyDisabled,
						},
						Privileged: true,
						LogConfig: container.LogConfig{
							Type: "json-file",
						},
						Mounts: []mount.Mount{
							{
								Type:   "bind",
								Source: dockerSockPath,
								Target: dockerSockPath,
							},
							{
								Type:   "bind",
								Source: "/var/lib/docker",
								Target: "/var/lib/docker",
							},
							{
								Type:   "bind",
								Source: "/root/.ssh",
								Target: "/root/.ssh",
								BindOptions: &mount.BindOptions{
									CreateMountpoint: true,
								},
							},
							{
								Type:   "bind",
								Source: tempDir,
								Target: tempDir,
							},
							{
								Type:   "bind",
								Source: "/root/.cache",
								Target: "/root/.cache",
								BindOptions: &mount.BindOptions{
									CreateMountpoint: true,
								},
							},
							{
								Type:   "bind",
								Source: "/root/.npm",
								Target: "/root/.npm",
								BindOptions: &mount.BindOptions{
									CreateMountpoint: true,
								},
							},
						},
					},
					&network.NetworkingConfig{},
					&ocispec.Platform{
						Architecture: platformArch,
						OS:           platformOS,
					},
					containerName,
				)
				if err != nil {
					_ = updateOutput(err)
					return err
				}

				defer func() {
					err = apiClient.ContainerStop(
						ctx,
						containerName,
						container.StopOptions{
							Signal:  "SIGINT",
							Timeout: helpers.Ptr(1),
						},
					)
					if err != nil {
						logAndUpdateOutput("warning: ContainerStop: %s", err.Error())
					}

					for {
						containerInspectResponse, err := apiClient.ContainerInspect(
							ctx,
							containerName,
						)
						if err != nil {
							time.Sleep(time.Second * 1)
							continue
						}

						if containerInspectResponse.State == nil {
							time.Sleep(time.Second * 1)
							continue
						}

						if containerInspectResponse.State.Running {
							time.Sleep(time.Second * 1)
							continue
						}

						if containerInspectResponse.State.ExitCode >= 0 {
							exitStatus = &containerInspectResponse.State.ExitCode
						}

						updateOutput(err)

						break
					}

					err = apiClient.ContainerRemove(
						ctx,
						containerName,
						container.RemoveOptions{
							RemoveVolumes: true,
							Force:         true,
						},
					)
					if err != nil {
						logAndUpdateOutput("warning: ContainerStop: %s", err.Error())
					}

				}()

				err = apiClient.ContainerStart(
					ctx,
					containerCreateResponse.ID,
					container.StartOptions{},
				)
				if err != nil {
					_ = updateOutput(err)
					return err
				}

				for {
					containerInspectResponse, err := apiClient.ContainerInspect(
						ctx,
						containerName,
					)
					if err != nil {
						time.Sleep(time.Second * 1)
						continue
					}

					if containerInspectResponse.State == nil {
						time.Sleep(time.Second * 1)
						continue
					}

					if containerInspectResponse.State.StartedAt == "" {
						time.Sleep(time.Second * 1)
						continue
					}

					if !containerInspectResponse.State.Running {
						time.Sleep(time.Second * 1)
						continue
					}

					break
				}

				containerLogsReader, err := apiClient.ContainerLogs(
					ctx,
					containerName,
					container.LogsOptions{
						ShowStdout: true,
						ShowStderr: true,
						Timestamps: true,
						Follow:     true,
					},
				)
				if err != nil {
					return err
				}

				lastWrite := time.Now().Add(-time.Millisecond * 100)

				b := make([]byte, 0)
				n := 0

				for {
					p := make([]byte, 65536)

					n, err = containerLogsReader.Read(p)
					if err != nil {
						if errors.Is(err, io.EOF) {
							break
						}
					}
					p = p[:n]

					if len(p) == 0 {
						continue
					}

					logBuffer = append(logBuffer, p...)
					b = append(b, p...)

					if time.Since(lastWrite) > time.Millisecond*100 {
						err = updateOutput(err)
						if err != nil {
							return err
						}

						lastWrite = time.Now()
						log.Printf("%d: %#+v", n, string(b))
						b = make([]byte, 0)
					}
				}

				err = updateOutput(nil)
				if err != nil {
					return err
				}

				log.Printf("%d: %#+v", n, string(b))

				err = updateOutput(err)
				if err != nil {
					return err
				}

				logAndUpdateOutput("%s done", containerName)

				return nil
			}

			execution = &api.Execution{
				JobName:      job.Name,
				Status:       internal.ExecutionOrTaskStatusRunning,
				StartedAt:    helpers.Ptr(time.Now().UTC()),
				RepositoryID: trigger.RepositoryID,
				ChangeID:     trigger.ChangeID,
				RuleID:       trigger.RuleID,
				TriggerID:    trigger.ID,
				JobID:        job.ID,
			}

			err = func() error {
				tx, err = db.Begin(ctx)
				if err != nil {
					return err
				}

				defer func() {
					_ = tx.Rollback(ctx)
				}()

				err = execution.Insert(ctx, tx, false, false)
				if err != nil {
					return err
				}

				err = tx.Commit(ctx)
				if err != nil {
					return err
				}

				return nil
			}()

			log.Printf("execution %s started", execution.ID)

			slices.SortFunc(job.ReferencedByTaskJobIDObjects, func(a *api.Task, b *api.Task) int {
				if a.Index < b.Index {
					return -1
				} else if a.Index > b.Index {
					return 1
				} else {
					return 0
				}
			})

			for _, task := range job.ReferencedByTaskJobIDObjects {
				friendlyTaskName := strings.ReplaceAll(strings.ReplaceAll(strings.ToLower(task.Name), " ", "-"), "_", "-")

				output := &api.Output{
					TaskIndex:    task.Index,
					TaskName:     friendlyTaskName,
					RepositoryID: execution.RepositoryID,
					ChangeID:     execution.ChangeID,
					RuleID:       execution.RuleID,
					TriggerID:    execution.TriggerID,
					JobID:        execution.JobID,
					ExecutionID:  execution.ID,
					TaskID:       task.ID,
					Status:       internal.ExecutionOrTaskStatusPending,
					TaskIDObject: task, // note: we've put this here
				}

				err = func() error {
					tx, err := db.Begin(ctx)
					if err != nil {
						return err
					}

					defer func() {
						_ = tx.Rollback(ctx)
					}()

					err = output.Insert(ctx, tx, false, false)
					if err != nil {
						return err
					}

					outputs = append(outputs, output)

					if !(execution.Status == internal.ExecutionOrTaskStatusFailing || execution.Status == internal.ExecutionOrTaskStatusErroring) {
						execution.Status = internal.ExecutionOrTaskStatusRunning

						err = execution.Update(ctx, tx, false)
						if err != nil {
							return err
						}
					}

					err = tx.Commit(ctx)
					if err != nil {
						return err
					}

					return nil
				}()
				if err != nil {
					return err
				}
			}

			for _, output := range outputs {
				err = doTask(output, execution)
				if err != nil {
					return err
				}
			}

			err = func() error {
				tx, err = db.Begin(ctx)
				if err != nil {
					return err
				}

				defer func() {
					_ = tx.Rollback(ctx)
				}()

				if execution.Status == internal.ExecutionOrTaskStatusRunning {
					execution.Status = internal.ExecutionOrTaskStatusSucceeded
				}

				execution.EndedAt = helpers.Ptr(time.Now().UTC())

				err = execution.Update(ctx, tx, false)
				if err != nil {
					return err
				}

				err = tx.Commit(ctx)
				if err != nil {
					return err
				}

				return nil
			}()

			log.Printf("execution %s done", execution.ID)

			return nil
		},
	)
}
