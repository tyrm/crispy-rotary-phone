package api

import (
	"github.com/gorilla/mux"
	"github.com/juju/loggo"
	"net/http"
	"time"
)

var logger = loggo.GetLogger("api")

type Server struct {
	router *mux.Router
	server *http.Server
}

func (s *Server) ListenAndServe() error {
	s.server = &http.Server{
		Handler:      s.router,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return s.server.ListenAndServe()
}

func (s *Server) Close() {
	err := s.server.Close()
	if err != nil {
		logger.Warningf("closing server: %s", err.Error())
	}
}

func NewServer() (*Server, error) {
	server := Server{
	}

	// setup router
	server.router = mux.NewRouter()
	server.router.Use(server.Middleware)

	// common error handlers
	server.router.NotFoundHandler = server.ErrorNotFoundHandler()
	server.router.MethodNotAllowedHandler = server.ErrorMethodNotAllowedHandler()

	return &server, nil
}