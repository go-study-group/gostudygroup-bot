package api

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

// Server ...
type Server struct {
	router   *mux.Router
	twitterC *twitterController
}

// NewServer returns a new Server
func NewServer() *Server {
	return &Server{}
}

// Initialize to wire up the routes.
func (s *Server) Initialize() {
	logger.Info("Initialize Server")
	s.router = mux.NewRouter()
	s.twitterC = newTwitterController(cfg.TwitterPostAPIToken)
	s.initializeRoutes()
}

func (s *Server) initializeRoutes() {
	// post a tweet to api
	s.router.HandleFunc("/api/v1/tweets/startinfive", s.twitterC.createTweet).Methods("POST")

	s.router.HandleFunc("/", staticHomeRoute)

	s.router.HandleFunc("/webhook/github/issuetrigger", handleGithubIssueTrigger)
}

// Run method to simply start the API Server
func (s *Server) Run() {
	s.Initialize() // Initialize the server
	logger.Info("Staring Server at " + cfg.Port)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	addr := ":" + cfg.Port

	// httpServer.
	h := &http.Server{Addr: addr, Handler: s.router}

	go func() {
		logger.Fatal(h.ListenAndServe())
	}()

	<-stop
	logger.Infof("Shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	h.Shutdown(ctx)

	logger.Infof("Server gracefully stopped")

}
