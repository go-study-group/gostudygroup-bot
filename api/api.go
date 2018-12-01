package api

import (
	"encoding/json"
	"net/http"

	"github.com/ankur-anand/gostudygroup-bot/twitterbot"
)

// Controller ...
type controller struct {
	twitterPostAPIKey string
}

// newController creates a new Controller.
func newController(key string) *controller {
	if key == "" {
		logger.Fatal("required twitterPostAPIKey is missing " + key)
	}
	return &controller{
		twitterPostAPIKey: key,
	}
}

// PostTweet decode the body has token, while posting tweet.
type PostTweet struct {
	Token string
}

// create a newTweet and post it to the twitter bot
func (c *controller) createTweet(w http.ResponseWriter, r *http.Request) {
	var pT PostTweet
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&pT); err != nil {
		resWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if pT.Token != c.twitterPostAPIKey {
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

func resWithError(w http.ResponseWriter, code int, message string) {
	resWithJSON(w, code, map[string]string{"error": message})
}

func resWithSuccess(w http.ResponseWriter, message string) {
	resWithJSON(w, http.StatusOK, map[string]string{"result": message})
}

func resWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
