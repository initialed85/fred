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
		change.Branch,
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

func GetExecutionSummary(execution *api.Execution) string {
	jobSummary := fmt.Sprintf("job %s", execution.JobID.String())
	if execution.JobIDObject != nil {
		jobSummary = GetJobSummary(execution.JobIDObject)
	}

	return fmt.Sprintf(
		"execution %s (%s) for %s",
		execution.ID.String(),
		execution.Status,
		jobSummary,
	)
}
