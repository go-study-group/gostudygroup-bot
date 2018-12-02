package api

import (
	"net/http"

	"github.com/google/go-github/v19/github"
)

const (
	newIssue = "opened"
)

func handleGithubIssueTrigger(w http.ResponseWriter, r *http.Request) {
	payload, err := github.ValidatePayload(r, []byte(cfg.GithubWebhookRepoAgendaSecretKey))
	if err != nil {
		logger.Fatal(err)
	}

	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		logger.Fatal(err)
	}

	switch eT := event.(type) {
	case *github.IssuesEvent:
		processIssuesEvent(eT)
	default:
		logger.Info("Info not github issue")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// processIssuesEventProcess the github IssuesEvent
// for various actions.
func processIssuesEvent(event *github.IssuesEvent) {

	// action has to be "opened" to take action.
	action := event.GetAction()
	if action != newIssue {
		return
	}

	issue := event.GetIssue()

	// we are only interested in github issue for this trigger.
	// and issue should not be pull request too, as every pull
	// request is also an issue in github api
	if issue.IsPullRequest() == true {
		return
	}

	// check if labels if already present.
	// if present no need to assign a new label.
	var labels []string
	for _, label := range issue.Labels {
		name := label.GetName()
		labels = append(labels, name)
	}

	if len(labels) != 0 {
		return
	}
}
