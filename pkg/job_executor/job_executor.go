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
	"github.com/jackc/pgx/v5/pgxpool"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"

	_ "embed"
)

//go:embed docker-entrypoint.sh
var dockerEntrypointScript string

var log = helpers.GetLogger("job_executor")

const claimDuration = time.Second * 30

func formatUUID(u uuid.UUID) string {
	return strings.ReplaceAll(u.String(), "-", "")
}

func getExecutionCtx(ctx context.Context) context.Context {
	ctx = query.WithLoad(ctx, api.ChangeTable)
	ctx = query.WithLoad(ctx, api.TriggerTable)
	ctx = query.WithLoad(ctx, api.JobTable)
	ctx = query.WithLoad(ctx, api.TaskTable)
	ctx = query.WithLoad(ctx, api.RepositoryTable)

	return ctx
}

func getOutputCtx(ctx context.Context) context.Context {
	ctx = query.WithLoad(ctx, api.TaskTable)
	ctx = query.WithLoad(ctx, api.LogTable)

	return ctx
}

func HandleExecution(ctx context.Context, db *pgxpool.Pool, execution *api.Execution, apiClient *client.Client, tempPath string) error {
	log.Printf("handling %s", internal.GetExecutionSummary(execution))

	var outputs []*api.Output

	err := func() error {
		var err error

		tx, err := db.Begin(ctx)
		if err != nil {
			return err
		}
		defer func() {
			_ = tx.Rollback(ctx)
		}()

		outputs, _, _, _, _, err = api.SelectOutputs(
			getOutputCtx(ctx),
			tx,
			fmt.Sprintf("%s = $$??", api.OutputTableExecutionIDColumn),
			nil,
			nil,
			nil,
			execution.ID,
		)
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
		err = execution.Update(getExecutionCtx(ctx), tx, false)
		if err != nil {
			return
		}

		for _, output := range outputs {
			output.Status = internal.ExecutionOrTaskStatusErrored
			output.EndedAt = helpers.Ptr(time.Now().UTC())
			err = output.Update(getOutputCtx(ctx), tx, false)
			if err != nil {
				return
			}
		}

		err = tx.Commit(ctx)
		if err != nil {
			return
		}
	}()

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

				execution.JobExecutorClaimedUntil = time.Now().UTC()
				err = execution.Update(getExecutionCtx(ctx), tx, false)
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

				execution.JobExecutorClaimedUntil = time.Now().UTC().Add(time.Second * 60)
				err = execution.Update(getExecutionCtx(ctx), tx, false)
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
				log.Printf("warning: execution claim refresh failed: %s", err.Error())
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

		err = output.Update(getOutputCtx(ctx), tx, false)
		if err != nil {
			return err
		}

		err = tx.Commit(ctx)
		if err != nil {
			return err
		}

		volumeName := fmt.Sprintf("fred-%s-%s", execution.ChangeIDObject.CommitHash, formatUUID(execution.ID))
		containerName := fmt.Sprintf("%s-%s", volumeName, task.Name)

		if execution.Status == internal.ExecutionOrTaskStatusFailing || execution.Status == internal.ExecutionOrTaskStatusErroring {
			log.Printf("%s skipped (due to failing / erroring parent execution)", containerName)
			return nil
		}

		var exitStatus *int
		logBuffer := make([]byte, 0)

		updateOutput := func(givenErr error) error {
			outputLog := output.LogIDObject

			tx, err := db.Begin(ctx)
			if err != nil {
				return err
			}

			defer func() {
				_ = tx.Rollback(ctx)
			}()

			if logBuffer != nil {
				outputLog.Buffer = logBuffer
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

			err = output.LogIDObject.Update(ctx, tx, false)
			if err != nil {
				return err
			}

			err = output.Update(getOutputCtx(ctx), tx, false)
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

		err = func() error {
			tx, err = db.Begin(ctx)
			if err != nil {
				return err
			}

			defer func() {
				_ = tx.Rollback(ctx)
			}()

			err = execution.ChangeIDObject.Reload(getExecutionCtx(ctx), tx)
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

		repositoryUrl := execution.ChangeIDObject.RepositoryIDObject.URL
		repositoryBranchName := execution.ChangeIDObject.BranchName
		parts := strings.Split(repositoryUrl, "/")
		repositoryFolderName := parts[len(parts)-1]
		repositoryCommitHash := execution.ChangeIDObject.CommitHash

		platformParts := strings.Split(task.Platform, "/")
		if len(platformParts) != 2 {
			err := fmt.Errorf("platform %#+v doesn't seem sane (want [os]/[arch])", task.Platform)
			_ = updateOutput(err)
			return err
		}
		platformOS := platformParts[0]
		platformArch := platformParts[1]

		builtinEnvVars := []string{
			fmt.Sprintf("FRED_REPOSITORY_ID=%s", execution.ChangeIDObject.RepositoryID.String()),
			fmt.Sprintf("FRED_CHANGE_ID=%s", execution.ChangeID.String()),
			fmt.Sprintf("FRED_RULE_ID=%s", execution.TriggerIDObject.RuleID.String()),
			fmt.Sprintf("FRED_JOB_ID=%s", execution.JobIDObject.ID.String()),
			fmt.Sprintf("FRED_JOB_NAME=%s", execution.JobIDObject.Name),
			fmt.Sprintf("FRED_TASK_ID=%s", task.ID.String()),
			fmt.Sprintf("FRED_TASK_NAME=%s", task.Name),
			fmt.Sprintf("FRED_TASK_PLATFORM=%s", task.Platform),
			fmt.Sprintf("FRED_TASK_PLATFORM_OS=%s", platformOS),
			fmt.Sprintf("FRED_TASK_PLATFORM_ARCH=%s", platformArch),
			fmt.Sprintf("FRED_TASK_IMAGE=%s", task.Image),
			fmt.Sprintf("FRED_TRIGGER_ID=%s", execution.TriggerIDObject.ID.String()),
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
				AttachStdout: true,
				AttachStderr: true,
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
						BindOptions: &mount.BindOptions{
							CreateMountpoint: true,
						},
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

	execution.Status = internal.ExecutionOrTaskStatusRunning
	execution.StartedAt = helpers.Ptr(time.Now().UTC())

	err = func() error {
		tx, err := db.Begin(ctx)
		if err != nil {
			return err
		}

		defer func() {
			_ = tx.Rollback(ctx)
		}()

		err = execution.Update(getExecutionCtx(ctx), tx, false)
		if err != nil {
			return fmt.Errorf("attempt to update execution to running failed: %s", err.Error())
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

	log.Printf("execution %s started", execution.ID)

	slices.SortFunc(execution.JobIDObject.ReferencedByTaskJobIDObjects, func(a *api.Task, b *api.Task) int {
		if a.Index < b.Index {
			return -1
		} else if a.Index > b.Index {
			return 1
		} else {
			return 0
		}
	})

	for _, output := range outputs {
		err = doTask(output, execution)
		if err != nil {
			return err
		}
	}

	err = func() error {
		tx, err := db.Begin(ctx)
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

		err = execution.Update(getExecutionCtx(ctx), tx, false)
		if err != nil {
			return fmt.Errorf("attempt to update execution to complete failed: %s", err.Error())
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

	log.Printf("execution %s done", execution.ID)

	return nil
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

	lastRunFoundNothing := false

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

			execution, err := api.JobExecutorClaimExecution(
				getExecutionCtx(ctx),
				tx,
				time.Now().UTC().Add(claimDuration),
				time.Second*2,
				fmt.Sprintf(
					"%s = $$??",
					api.ExecutionTableChangeIDColumn,
				),
				helpers.Ptr(fmt.Sprintf("%s DESC", api.ExecutionTableCreatedAtColumn)),
			)
			if err != nil {
				return fmt.Errorf("attempt to claim an execution failed: %s", err.Error())
			}

			if execution == nil {
				if !lastRunFoundNothing {
					lastRunFoundNothing = true
					log.Printf("no executions to claim")
				}

				return nil
			}

			defer func() {
				if execution != nil {
					ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
					defer cancel()

					tx, err := db.Begin(ctx)
					if err != nil {
						return
					}

					defer func() {
						_ = tx.Rollback(ctx)
					}()

					execution.JobExecutorClaimedUntil = time.Now().UTC()
					err = execution.Update(getExecutionCtx(ctx), tx, false)
					if err != nil {
						err = fmt.Errorf("attempt to release claim on %#+v failed: %s", execution, err.Error())
						log.Printf("warning: %s", err.Error())
					}

					_ = tx.Commit(ctx)
				}
			}()

			lastRunFoundNothing = false

			err = tx.Commit(ctx)
			if err != nil {
				return err
			}

			log.Printf("claimed %s", internal.GetExecutionSummary(execution))

			err = HandleExecution(ctx, db, execution, apiClient, tempPath)
			if err != nil {
				return fmt.Errorf("attempt to handle %s failed: %s", internal.GetExecutionSummary(execution), err)
			}

			return nil
		},
	)
}
