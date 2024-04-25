package main

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5/middleware"
	log "github.com/sirupsen/logrus"
	scripts "github.com/valentyna-koshelnyk/panda-eats-prototype-api/scripts"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
)

var (
	restaurantsJSON = scripts.ConverterRestaurant()
	menuJSON        = scripts.ConverterMenu()
)

// Set handlers
//TODO implement pagination

// Fetch all restaurants handler
func getAllRestaurants(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(restaurantsJSON); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

// Fetch all menus handler
func getAllMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(menuJSON); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

func init() {
	// TODO: implement a custom structured logger
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	// Log for informational purposes, then depends on necessity use log.SetLevel(log.WarnLevel)
	log.SetLevel(log.InfoLevel)

}

func main() {
	// Set a router
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Initiate a server
	server := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	//Set routes

	// A route for health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("OK"))
		if err != nil {
			log.Error("Failed to write response: %v", err)
		}
	})
	// A route for restaurants
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/restaurants", getAllRestaurants)
		r.Get("/menus", getAllMenu)
	})

	// Start the server
	go func() {
		log.Info("Starting server on port :3000")
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

// Pagination Object
type Pagination struct {
	Next          int
	Previous      int
	RecordPerPage int
	CurrentPage   int
	TotalPage     int
}
