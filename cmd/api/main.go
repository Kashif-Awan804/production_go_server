package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Kashif-Awan804/production_go_server/internal/config"
	"github.com/Kashif-Awan804/production_go_server/internal/handlers"
)

func main() {

	cfg := config.LoadConfig()

	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.Health)

	server := &http.Server{
		Addr:              ":" + cfg.PORT,
		Handler:           mux,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Println("server running on :8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
