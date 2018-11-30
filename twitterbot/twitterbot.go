package twitterbot

import (
	"bytes"
	"fmt"
	"log"
	"net/url"
	"os"
	"text/template"
	"time"

	"github.com/ChimeraCoder/anaconda"
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
	// twtTmpl stores the parsed template
	twtTmpl *template.Template
)

func getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("required environment variable is missing " + name)
	}

	return v
}

func getCurrentDate() string {
	// Assuming this will be running from the
	// some sort of cron job the date will
	// be always correct even in different
	// timezone as it's handled at configuartion
	// level
	return time.Now().Format(dateLayoutTest)
}

func init() {

	var err error
	// a named Template.
	twtTmpl = template.New("twitterText")
	// parse the twitter status content and generate the template.
	twtTmpl, err = twtTmpl.Parse(tweetTemplate)

	if err != nil {
		log.Fatalln("Twitterbot Template Parse Error: ", err)
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
		log.Fatalln("Twitterbot Template Execute Error: ", err)
	}

	return out.String()
}

// PostNewTweet Post's a new tweet to account.
func PostNewTweet() error {

	tweetStatus := getTweetText()

	fmt.Println(tweetStatus)

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	twt, err := api.PostTweet(tweetStatus, url.Values{})
	if err != nil {
		panic(err)
	}

	fmt.Println(twt.Text)
	return nil
}
