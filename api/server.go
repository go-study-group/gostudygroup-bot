package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ankur-anand/gostudygroup-bot/helper"
	"github.com/gorilla/mux"
)

var (
	twitterPostAPIKey = os.Getenv("TWITTER_POST_API_KEY")
	// logger
	logger = helper.Logger
)

// Server ...
type Server struct {
	Router     *mux.Router
	Controller *controller
}

// NewServer returns a new Server
func NewServer() *Server {
	return &Server{}
}

// Initialize to wire up the routes.
func (s *Server) Initialize() {
	logger.Info("Initialize Server")
	s.Router = mux.NewRouter()
	logger.Info("twitterPostAPIKey" + twitterPostAPIKey)
	s.Controller = newController(twitterPostAPIKey)
	s.initializeRoutes()
}

func (s *Server) initializeRoutes() {
	// post a tweet to api
	s.Router.HandleFunc("/api/v1/tweets/startinfive", s.Controller.createTweet).Methods("POST")

	s.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "{\"message\": \"server is up and running\"}")
	})
}

// Run method to simply start the API Server
func (s *Server) Run(addr string) {
	s.Initialize() // Initialize the server
	logger.Info("Staring Server at " + addr)

	if addr == "" {
		logger.Fatal("required addr is missing" + addr)
	}
	logger.Fatal(http.ListenAndServe(addr, s.Router))

}
