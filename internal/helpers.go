package internal

import (
	"fmt"

	"github.com/initialed85/fred/pkg/api"
)

func GetChangeSummary(change *api.Change) string {
	repositoryURL := fmt.Sprintf("repository %s", change.RepositoryID)
	if change.RepositoryIDObject != nil {
		repositoryURL = change.RepositoryIDObject.URL
	}

	return fmt.Sprintf(
		"change %s for %s:%s@%s",
		change.ID.String(),
		repositoryURL,
		change.BranchName,
		change.CommitHash,
	)
}

func GetJobSummary(job *api.Job) string {
	return fmt.Sprintf(
		"job %s %#+v",
		job.ID.String(),
		job.Name,
	)
}

func GetTriggerSummary(trigger *api.Trigger) string {
	jobSummary := fmt.Sprintf("job %s", trigger.JobID.String())
	if trigger.JobIDObject != nil {
		jobSummary = GetJobSummary(trigger.JobIDObject)
	}

	return fmt.Sprintf(
		"trigger %s for %s",
		trigger.ID.String(),
		jobSummary,
	)
}

func GetExecutionSummary(execution *api.Execution) string {
	jobSummary := fmt.Sprintf("job %s", execution.JobID.String())
	if execution.JobIDObject != nil {
		jobSummary = GetJobSummary(execution.JobIDObject)
	}

	return fmt.Sprintf(
		"execution %s for %s",
		execution.ID.String(),
		jobSummary,
	)
}
