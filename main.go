package main

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("OK"))
		if err != nil {
			log.Error("Failed to write response: %v", err)
		}
	})

	// Create a server
	server := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	// Start the server
	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
		log.Info("Stopped serving new connections.")
	}()

	// Create a channel to receive  notifications from signal
	sigChan := make(chan os.Signal, 1)

	//Register the given channel to receive notifications
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	// Block until the signal is received
	<-sigChan

	//Shut down gracefully
	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}
	log.Info("Graceful shutdown complete.")
}
