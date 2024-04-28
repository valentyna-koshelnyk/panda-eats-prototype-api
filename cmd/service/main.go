package main

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1/menu"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1/restaurant"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/repository"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
)

var version string

func initConfig() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config.dev")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}

func init() {
	// TODO: implement a custom structured logger
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	// Log for informational purposes, then depends on necessity use log.SetLevel(log.WarnLevel)
	log.SetLevel(log.InfoLevel)
	log.Info("starting up API...", log.WithField("version", version))

}

func main() {
	initConfig()
	// Set a router
	r := chi.NewRouter()
	// Generate a unique identifier for every incoming HTTP request
	r.Use(middleware.RequestID)
	// Log start and end of each request
	r.Use(middleware.Logger)
	// Prevent server crashes
	r.Use(middleware.Recoverer)

	// Initiate a server
	server := &http.Server{
		Addr:    ":" + viper.GetString("server.port"),
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
	r.Route("/api/v1/menus", func(r chi.Router) {
		r.Get("/", menu.GetAllMenus)
	})
	r.Route("/api/v1/restaurants", func(r chi.Router) {
		r.Get("/", repository.PaginateHandler)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", restaurant.GetRestaurantById)
		})
	})

	r.Route("/api/v1/restaurants/{restaurant_id}/items", func(r chi.Router) {
		r.Get("/", menu.GetMenuByRestaurant)
	})

	// Start the server
	go func() {
		log.Info("Starting server on port :", viper.GetString("server.port"))
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
