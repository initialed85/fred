package job_executor

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
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
			api.TriggerTableJobExecutionStartedAtColumn,
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

			now := time.Now().UTC()
			trigger.JobExecutorClaimedUntil = now
			trigger.JobExecutionStartedAt = &now

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

			doTask := func(task *api.Task, execution *api.Execution) error {
				ctx, cancel := context.WithCancel(ctx)
				defer cancel()

				output := &api.Output{
					TaskID: task.ID,
					Status: internal.ExecutionOrTaskStatusRunning,
				}

				var exitStatus *int

				logBuffer := make([]byte, 0)

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

					if task == job.BuildTaskIDObject {
						execution.BuildOutputID = &output.ID
					} else if task == job.TestTaskIDObject {
						execution.TestOutputID = &output.ID
					} else if task == job.PublishTaskIDObject {
						execution.PublishOutputID = &output.ID
					} else if task == job.DeployTaskIDObject {
						execution.DeployOutputID = &output.ID
					} else if task == job.ValidateTaskIDObject {
						execution.ValidateOutputID = &output.ID
					} else {
						return fmt.Errorf("assertion failed: could not work out if %#+v was for build / test / publish / deploy / validate", output)
					}

					execution.Status = internal.ExecutionOrTaskStatusRunning

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
				if err != nil {
					return err
				}

				updateOutput := func(givenErr error) error {
					tx, err := db.Begin(ctx)
					if err != nil {
						return err
					}

					defer func() {
						_ = tx.Rollback(ctx)
					}()

					if logBuffer != nil {
						output.Buffer = string(logBuffer)
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

				apiClient, err := client.NewClientWithOpts(client.FromEnv)
				if err != nil {
					_ = updateOutput(err)
					return err
				}

				apiClient.NegotiateAPIVersion(ctx)

				defer func() {
					_ = apiClient.Close()
				}()

				containerName := fmt.Sprintf("change-%s-job-%s-task-%s", trigger.ChangeID, job.ID, task.ID)

				tempDir, err := os.MkdirTemp(tempPath, fmt.Sprintf("%s-*", containerName))
				if err != nil {
					_ = updateOutput(err)
					return err
				}

				dockerEntrypointOutsideFilePath := filepath.Join(tempDir, "docker-entrypoint.sh")
				err = os.WriteFile(dockerEntrypointOutsideFilePath, []byte(dockerEntrypointScript), 0o777)
				if err != nil {
					_ = updateOutput(err)
					return err
				}
				log.Printf("prepared %s", dockerEntrypointOutsideFilePath)

				taskEntrypointOutsideFilePath := filepath.Join(tempDir, "task-entrypoint.sh")
				err = os.WriteFile(taskEntrypointOutsideFilePath, []byte(task.Script), 0o777)
				if err != nil {
					_ = updateOutput(err)
					return err
				}
				log.Printf("prepared %s", taskEntrypointOutsideFilePath)

				log.Printf("pulling %s", task.Image)

				_, err = apiClient.ImagePull(
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

				log.Printf("creating %s", containerName)

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
				os := platformParts[0]
				arch := platformParts[1]

				containerCreateResponse, err := apiClient.ContainerCreate(
					ctx,
					&container.Config{
						Image: task.Image,
						Env: []string{
							"CI=true",
							"COMPOSE_DOCKER_CLI_BUILD=1",
							"DOCKER_BUILDKIT=1",
							fmt.Sprintf("COMPOSE_PROJECT_NAME=%s", containerName),
							fmt.Sprintf("REPOSITORY_URL=%s", repositoryUrl),
							fmt.Sprintf("REPOSITORY_BRANCH_NAME=%s", repositoryBranchName),
							fmt.Sprintf("REPOSITORY_FOLDER_NAME=%s", repositoryFolderName),
							fmt.Sprintf("REPOSITORY_COMMIT_HASH=%s", repositoryCommitHash),
						},
						WorkingDir:   tempDir,
						Entrypoint:   []string{fmt.Sprintf("/%s/docker-entrypoint.sh", tempDir)},
						Cmd:          []string{},
						StopTimeout:  helpers.Ptr(1),
						Tty:          true,
						AttachStdout: true,
						AttachStderr: true,
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
								Source: "/var/run",
								Target: "/var/run",
							},
							{
								Type:   "bind",
								Source: "/var/lib/docker",
								Target: "/var/lib/docker",
							},
							{
								Type:   "bind",
								Source: tempDir,
								Target: tempDir,
							},
						},
					},
					&network.NetworkingConfig{},
					&ocispec.Platform{
						Architecture: arch,
						OS:           os,
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
						log.Printf("warning: ContainerStop: %s", err.Error())
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
						log.Printf("warning: ContainerStop: %s", err.Error())
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

				for {
					p := make([]byte, 65536)

					n, err := containerLogsReader.Read(p)
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

				err = updateOutput(err)
				if err != nil {
					return err
				}

				log.Printf("%s done", containerName)

				return nil
			}

			execution := &api.Execution{
				TriggerID: trigger.ID,
				JobID:     job.ID,
				Status:    internal.ExecutionOrTaskStatusCreated,
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

			failed := false

			if !failed && job.BuildTaskIDObject != nil {
				err = doTask(job.BuildTaskIDObject, execution)
				if err != nil {
					return err
				}

				failed = execution.BuildOutputIDObject != nil && execution.BuildOutputIDObject.Status == internal.ExecutionOrTaskStatusFailed
			}

			if !failed && job.TestTaskIDObject != nil {
				err = doTask(job.TestTaskIDObject, execution)
				if err != nil {
					return err
				}

				failed = execution.TestOutputIDObject != nil && execution.TestOutputIDObject.Status == internal.ExecutionOrTaskStatusFailed
			}

			if !failed && job.PublishTaskIDObject != nil {
				err = doTask(job.PublishTaskIDObject, execution)
				if err != nil {
					return err
				}

				failed = execution.PublishOutputIDObject != nil && execution.PublishOutputIDObject.Status == internal.ExecutionOrTaskStatusFailed
			}

			if !failed && job.DeployTaskIDObject != nil {
				err = doTask(job.DeployTaskIDObject, execution)
				if err != nil {
					return err
				}

				failed = execution.DeployOutputIDObject != nil && execution.DeployOutputIDObject.Status == internal.ExecutionOrTaskStatusFailed
			}

			if !failed && job.ValidateTaskIDObject != nil {
				err = doTask(job.ValidateTaskIDObject, execution)
				if err != nil {
					return err
				}

				failed = execution.ValidateOutputIDObject != nil && execution.ValidateOutputIDObject.Status == internal.ExecutionOrTaskStatusFailed
			}

			err = func() error {
				tx, err = db.Begin(ctx)
				if err != nil {
					return err
				}

				defer func() {
					_ = tx.Rollback(ctx)
				}()

				if failed {
					execution.Status = internal.ExecutionOrTaskStatusFailed
				} else {
					execution.Status = internal.ExecutionOrTaskStatusSucceeded
				}

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
