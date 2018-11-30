package main

import (
	"github.com/ankur-anand/gostudygroup-bot/helper"
	"github.com/ankur-anand/gostudygroup-bot/twitterbot"
)

var logger = helper.Logger

func main() {

	logger.Info("Staring Bot AppLication")
	twitterbot.PostNewTweet()
	logger.Info("Stoping Bot AppLication")
	logger.Sync()
}
