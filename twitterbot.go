package main

import (
	"bytes"
	"fmt"
	"net/url"
	"text/template"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/ankur-anand/gostudygroup-bot/config"
	"github.com/ankur-anand/gostudygroup-bot/helper"
)

// YeildWhen ...
type YeildWhen struct {
	When string
}

const (
	tweetTemplate = `The weekly developer meeting is happening in 1 hour!

	Absolutely everyone is welcome to join, hope to see you there!
	
	Video link: https://aka.ms/athensdevzoom`
	// sort form that we can refer to.
	dateLayout = "02-Jan-2006"
	// below is just used for testing purpose.
	dateLayoutTest = time.RFC850
	prod           = "production"
)

var (

	// logger
	logger = helper.Logger
)

// TwitterBot ...
type TwitterBot struct {
	// twtTmpl stores the parsed template
	twtTmpl *template.Template
	cfg     config.Config
}

// New returns a new twitterBot
func New(cfg config.Config) TwitterBot {

	var err error
	// a named Template.
	twtTmpl := template.New("twitterText")
	// parse the twitter status content and generate the template.
	twtTmpl, err = twtTmpl.Parse(tweetTemplate)

	if err != nil {
		logger.Fatal("Twitterbot Template Parse Error: ", err)
	}

	return TwitterBot{
		twtTmpl: twtTmpl,
		cfg:     cfg,
	}
}

func getCurrentDate(goEnv string) string {
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

func (b TwitterBot) getTweetText() string {
	yieldWhen := YeildWhen{
		When: getCurrentDate(b.cfg.GoEnv),
	}

	var out bytes.Buffer
	err := b.twtTmpl.Execute(&out, yieldWhen)

	if err != nil {
		// TODO: Handle it in good way.
		logger.Fatal("Twitterbot Template Execute Error: ", err)
	}

	return out.String()
}

// PostNewTweet Post's a new tweet to account.
func (b TwitterBot) PostNewTweet() (string, error) {

	tweetStatus := b.getTweetText()

	anaconda.SetConsumerKey(b.cfg.TwitterConsumerKey)
	anaconda.SetConsumerSecret(b.cfg.TwitterConsumerSecret)
	api := anaconda.NewTwitterApi(b.cfg.TwitterAccessToken, b.cfg.TwitterAccessTokenSecret)

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
