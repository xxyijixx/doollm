package server

import (
	"doollm/config"
	"doollm/internal/api"
	"doollm/internal/repository"
	"fmt"
	"log"
	"net/http"
)

func Run() {
	// g := gin.Default()
	// RouteMount(g)

	repository.InitDB()

	api.SetupRoutes()

	StartScheduledTask()
	log.Printf("Server starting on http://localhost:%d/", config.EnvConfig.PORT)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.EnvConfig.PORT), nil); err != nil {
		log.Fatal(err)
	}

	// g.Run(":9090")
}
