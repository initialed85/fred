package internal

import (
	"fmt"

	"github.com/initialed85/fred/pkg/api"
)

func GetChangeSummary(change *api.Change) string {
	repositoryURL := fmt.Sprintf("repository (%s)", change.RepositoryID)
	if change.RepositoryIDObject != nil {
		repositoryURL = change.RepositoryIDObject.URL
	}

	return fmt.Sprintf(
		"change (%s) for %s:%s:%s",
		change.ID.String(),
		repositoryURL,
		change.BranchName,
		change.CommitHash,
	)
}

func GetTriggerSummary(trigger *api.Trigger) string {
	changeSummary := fmt.Sprintf("change (%s)", trigger.ChangeID.String())
	if trigger.ChangeIDObject != nil {
		changeSummary = GetChangeSummary(trigger.ChangeIDObject)
	}

	return fmt.Sprintf(
		"trigger (%s) for %s",
		trigger.ID.String(),
		changeSummary,
	)
}

func GetJobSummary(job *api.Job) string {
	return fmt.Sprintf(
		"job (%s - %#+v)",
		job.ID.String(),
		job.Name,
	)
}
