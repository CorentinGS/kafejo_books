package app

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Router *chi.Mux
	Stage  string
}

func NewServer(router *chi.Mux, stage string) *Server {
	return &Server{
		Router: router,
		Stage:  stage,
	}
}

func (s *Server) Run(addr string) error {
	httpServer := &http.Server{
		Addr:              addr,
		ReadHeaderTimeout: 5 * time.Second, //nolint:mnd // This is the default value
		Handler:           s.Router,
	}

	return httpServer.ListenAndServe()
}
