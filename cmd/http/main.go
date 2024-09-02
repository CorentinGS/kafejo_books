package main

import (
	"log"

	"github.com/corentings/kafejo-books/app"
	"github.com/corentings/kafejo-books/config"
	"github.com/go-chi/chi/v5"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("❌ Error loading config: %s", err.Error())
	}

	chiRouter := chi.NewRouter()

	server := app.NewServer(chiRouter, cfg)

	app.RegisterRoutes(server)

	if err = server.Run(":" + cfg.Port); err != nil {
		log.Fatalf("failed to run server: %v", err.Error())
	}
}
