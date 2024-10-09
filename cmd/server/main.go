package main

import (
	"doollm/internal/api"
	"doollm/internal/repository"
	"doollm/pkg/config"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig("./configs/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	repository.InitDB(&cfg.Database)

	api.SetupRoutes()

	log.Printf("Server starting on http://localhost%s/", cfg.Server.Port)
	if err := http.ListenAndServe(cfg.Server.Port, nil); err != nil {
		log.Fatal(err)
	}
}
