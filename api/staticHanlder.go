package api

import (
	"net/http"
)

func staticHomeRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resWithSuccess(w, "server is up and running")
}
