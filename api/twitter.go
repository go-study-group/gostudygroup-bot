package api

import (
	"encoding/json"
	"net/http"

	"github.com/ankur-anand/gostudygroup-bot/twitterbot"
)

// twitterController ...
type twitterController struct {
	twitterPostAPIToken string
}

// newtwitterController creates a new twitterController.
func newTwitterController(key string) *twitterController {
	if key == "" {
		logger.Fatal("required twitterPostAPIToken is missing " + key)
	}
	return &twitterController{
		twitterPostAPIToken: key,
	}
}

// PostTweet decode the body has token, while posting tweet.
type PostTweet struct {
	Token string
}

// create a newTweet and post it to the twitter bot
func (c *twitterController) createTweet(w http.ResponseWriter, r *http.Request) {
	var pT PostTweet
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&pT); err != nil {
		resWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if pT.Token != c.twitterPostAPIToken {
		resWithError(w, http.StatusUnauthorized, "UnAuthorized Request")
		return
	}

	res, err := twitterbot.PostNewTweet()

	// if err is not nil repond with something went wrong
	if err != nil {
		resWithError(w, http.StatusInternalServerError, "Something went Wrong")
		return
	}

	resWithSuccess(w, res)
}
