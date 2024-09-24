package internal

const (
	TriggerTableJobExecutorClaimedUntilAdvisoryLockID = 1
	TriggerProducerAdvisoryLockID                     = 2
	JobExecutorAdvisoryLockID                         = 3
	ExecutionOrTaskStatusCreated                      = "created"
	ExecutionOrTaskStatusRunning                      = "running"
	ExecutionOrTaskStatusSucceeded                    = "succeded"
	ExecutionOrTaskStatusFailed                       = "failed"
	ExecutionOrTaskStatusErrored                      = "errored"
)
