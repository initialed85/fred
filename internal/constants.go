package internal

const (
	TriggerTableJobExecutorClaimedUntilAdvisoryLockID = 1
	TriggerProducerAdvisoryLockID                     = 2
	JobExecutorAdvisoryLockID                         = 3
	ExecutionOrTaskStatusCreated                      = "created"
	ExecutionOrTaskStatusRunning                      = "running"
	ExecutionOrTaskStatusSucceeded                    = "succeded"
	ExecutionOrTaskStatusFailing                      = "failing"
	ExecutionOrTaskStatusFailed                       = "failed"
	ExecutionOrTaskStatusErroring                     = "erroring"
	ExecutionOrTaskStatusErrored                      = "errored"
	ExecutionOrTaskStatusSkipped                      = "skipped"
)
