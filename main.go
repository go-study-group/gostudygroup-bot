package main

import (
	"os"

	"github.com/ankur-anand/gostudygroup-bot/api"
	"github.com/ankur-anand/gostudygroup-bot/helper"
)

var logger = helper.Logger

func main() {

	logger.Info("Staring Bot AppLication")

	apiServer := api.NewServer()
	port := os.Getenv("PORT")
	apiServer.Run(":" + port)
	logger.Info("Stoping Bot AppLication")
	logger.Sync()
}
