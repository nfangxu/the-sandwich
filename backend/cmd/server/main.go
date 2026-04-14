package main

import (
	"fmt"
	"log"

	"github.com/the-sandwich/backend/internal/infrastructure"
	"github.com/the-sandwich/backend/internal/interface/api"
	"github.com/the-sandwich/backend/internal/interface/ws"
)

func main() {
	cfg, err := infrastructure.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	app, err := infrastructure.InitializeApp(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	hub, err := ws.NewHub()
	if err != nil {
		log.Fatalf("Failed to create hub: %v", err)
	}
	go hub.Run()

	handlers := api.NewHandlers(app.AuthSvc)
	router, server := api.SetupRouter(hub, handlers)

	addr := fmt.Sprintf("%s:%d", cfg.App.Host, cfg.App.Port)
	log.Printf("Starting server on %s", addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	_ = router
}
