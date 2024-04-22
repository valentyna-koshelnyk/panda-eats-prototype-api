package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	// Log for informational purposes, then depends on necessity use log.SetLevel(log.WarnLevel)
	log.SetLevel(log.InfoLevel)

}

func main() {
	// Set a router
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("OK"))
		if err != nil {
			log.Error("Failed to write response: %v", err)
		}
	})

	log.Info("Starting server on port :3000")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Error("Error starting server: %v", err)
	}

}
