package app

import (
	"net/http"
	"time"

	"github.com/corentings/kafejo-books/config"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	Router *chi.Mux
	Stage  string
	Config *config.Config
}

func NewServer(router *chi.Mux, config *config.Config) *Server {
	return &Server{
		Router: router,
		Config: config,
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
