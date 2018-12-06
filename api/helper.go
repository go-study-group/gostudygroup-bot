package api

import (
	"encoding/json"
	"net/http"

	"github.com/ankur-anand/gostudygroup-bot/config"
	"github.com/ankur-anand/gostudygroup-bot/helper"
)

var (
	cfg    = config.Cfg
	logger = helper.Logger
)

// resWithError JSON
func resWithError(w http.ResponseWriter, code int, message string) {
	resWithJSON(w, code, map[string]string{"error": message})
}

// resWithSuccess JSON
func resWithSuccess(w http.ResponseWriter, message string) {
	resWithJSON(w, http.StatusOK, map[string]string{"result": message})
}

func resWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
