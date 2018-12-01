package twitterbot

import (
	"bytes"
	"fmt"
	"net/url"
	"os"
	"text/template"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/ankur-anand/gostudygroup-bot/helper"
)

// YeildWhen ...
type YeildWhen struct {
	When string
}

const (
	tweetTemplate = `The study group is starting in 5 minutes! Come join us, absolutely everyone is welcome!

	https://zoom.us/j/714787795
	When: {{ .When}} at 17:00-18:00 UTC
	Details: https://docs.google.com/document/d/16m99AvcTL_BJOIbR4jkUDSHyApYUDb0VgC9UPBJMed0/edit#
	`
	// sort form that we can refer to.
	dateLayout = "02-Jan-2006"
	// below is just used for testing purpose.
	dateLayoutTest = time.RFC850
)

var (
	consumerKey       = getenv("TWITTER_CONSUMER_KEY")
	consumerSecret    = getenv("TWITTER_CONSUMER_SECRET")
	accessToken       = getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret = getenv("TWITTER_ACCESS_TOKEN_SECRET")
	goEnv             = getenv("GO_ENV")
	prod              = "production"
	// twtTmpl stores the parsed template
	twtTmpl *template.Template
	// logger
	logger = helper.Logger
)

func getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		logger.Fatal("required environment variable is missing " + name)
	}

	return v
}

func getCurrentDate() string {
	// Assuming this will be running from the
	// some sort of cron job the date will
	// be always correct even in different
	// timezone as it's handled at configuartion
	// level
	if goEnv == prod {
		return time.Now().Format(dateLayout)
	}
	return time.Now().Format(dateLayoutTest)
}

func init() {

	var err error
	// a named Template.
	twtTmpl = template.New("twitterText")
	// parse the twitter status content and generate the template.
	twtTmpl, err = twtTmpl.Parse(tweetTemplate)

	if err != nil {
		logger.Fatal("Twitterbot Template Parse Error: ", err)
	}

}

func getTweetText() string {
	yieldWhen := YeildWhen{
		When: getCurrentDate(),
	}

	var out bytes.Buffer
	err := twtTmpl.Execute(&out, yieldWhen)

	if err != nil {
		// TODO: Handle it in good way.
		logger.Fatal("Twitterbot Template Execute Error: ", err)
	}

	return out.String()
}

// PostNewTweet Post's a new tweet to account.
func PostNewTweet() (string, error) {

	tweetStatus := getTweetText()

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	logger.Infof("Posting Tweet with Content")
	logger.Infof(tweetStatus)

	twt, err := api.PostTweet(tweetStatus, url.Values{})

	// these error's need to be monitored precisely
	// as these error can mostly be from twitter api.
	if err != nil {
		logger.Warn(err)
		return err.Error(), nil
	}
	message := fmt.Sprintf("The tweet has been successfully posted, to Handle [%s] and status ID is [%d]",
		twt.User.ScreenName, twt.Id)

	logger.Infof(message)

	return message, nil
}
