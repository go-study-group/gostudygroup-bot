package githubbot

import (
	"context"
	"net/http"
	"time"

	"github.com/ankur-anand/gostudygroup-bot/config"
	"github.com/ankur-anand/gostudygroup-bot/helper"
	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/v19/github"
)

var (
	logger = helper.Logger
)

// GithubBot Represnt's a bot with
// github client
type GithubBot struct {
	client *github.Client
}

// New returns a new Github bot.
func New(cfg config.Config) GithubBot {
	// initialize the token
	tr := http.DefaultTransport

	// Wrap the shared transport for use
	itr, err := ghinstallation.NewKeyFromFile(tr, cfg.GithubIssueLabelerIntegrationID, cfg.GithubIssueLabelerInstallationID, cfg.GithubIssueLabelerPemFilePath)
	if err != nil {
		logger.Fatal(err)
	}
	client := github.NewClient(&http.Client{Transport: itr})

	return GithubBot{
		client: client,
	}
}

// LabelIssue ...
func (g GithubBot) LabelIssue(event *github.IssuesEvent, label []string) {

	// get the repo structs
	repo := event.GetRepo()
	issue := event.GetIssue()
	// issue Number
	issueNumber := issue.GetNumber()
	user := repo.GetOwner()
	owner := user.GetLogin()
	repoName := repo.GetName()

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	lbsl, res, err := g.client.Issues.AddLabelsToIssue(ctx, owner, repoName, issueNumber, label)
	if err != nil {
		logger.Info("Error while adding labels to the issue")
		logger.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		logger.Info("Status Code of AddLabelsToIssue Not Ok " + string(res.StatusCode))
		return
	}
	logger.Info("The following labels were Added to Issue")
	logger.Info(lbsl)
}
