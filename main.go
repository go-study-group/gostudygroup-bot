package main

import (
	"github.com/gomods/twitter-bot/api"
	"github.com/gomods/twitter-bot/helper"
)

var logger = helper.Logger

func main() {

	logger.Info("Staring Bot AppLication")

	apiServer := api.NewServer()

	apiServer.Run()
	logger.Info("Stoping Bot AppLication")
	logger.Sync()
}
