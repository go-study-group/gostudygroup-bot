package config

import (
	"os"

	"github.com/ankur-anand/gostudygroup-bot/helper"
)

var (
	// Twitter Specific

	// TwitterConsumerKey envKey TWITTER_CONSUMER_KEY...
	twitterConsumerKey = getenv("TWITTER_CONSUMER_KEY")
	// TwitterConsumerSecret envKey TWITTER_CONSUMER_SECRET...
	twitterConsumerSecret = getenv("TWITTER_CONSUMER_SECRET")
	// TwitterAccessToken envKey TWITTER_ACCESS_TOKEN..
	twitterAccessToken = getenv("TWITTER_ACCESS_TOKEN")
	// TwitterAccessTokenSecret envKey TWITTER_ACCESS_TOKEN_SECRET...
	twitterAccessTokenSecret = getenv("TWITTER_ACCESS_TOKEN_SECRET")

	// Github Specific
	githubWebhookRepoAgendaSecretKey = getenv("GITHUB_WEBHOOK_REPOAGENDA_SECRET_KEY")
	// GoEnv envKey GO_ENV..
	goEnv               = getenv("GO_ENV")
	port                = getenv("PORT")
	twitterPostAPIToken = getenv("TWITTER_POST_API_TOKEN")
	// Logger ...
	logger = helper.Logger
)

func getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		logger.Fatal("required environment variable is missing " + name)
	}

	return v
}

// Config ...
type Config struct {
	TwitterConsumerKey               string
	TwitterConsumerSecret            string
	TwitterAccessToken               string
	TwitterAccessTokenSecret         string
	TwitterPostAPIToken              string
	GoEnv                            string
	GithubWebhookRepoAgendaSecretKey string
	Port                             string
}

// GetConfig returns a Config structs holding
// Environment variables.
func GetConfig() Config {
	return Config{
		TwitterConsumerKey:               twitterConsumerKey,
		TwitterConsumerSecret:            twitterConsumerSecret,
		TwitterAccessToken:               twitterAccessToken,
		TwitterAccessTokenSecret:         twitterAccessTokenSecret,
		TwitterPostAPIToken:              twitterPostAPIToken,
		GoEnv:                            goEnv,
		GithubWebhookRepoAgendaSecretKey: githubWebhookRepoAgendaSecretKey,
		Port:                             port,
	}
}
