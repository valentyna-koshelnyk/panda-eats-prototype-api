package main

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1/restaurant"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
)

//const (
//	// PageIDKey refers to the context key that stores the next page id
//	PageIDKey CustomKey = "page_id"
//)

// Set handlers
//TODO implement pagination

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
	// Generate a unique identifier for every incoming HTTP request
	r.Use(middleware.RequestID)
	// Log start and end of each request
	r.Use(middleware.Logger)
	// Prevent server crashes
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
	//
	//r.Route("/api/v1", func(r chi.Router) {
	//	r.With(m.Pagination).Get("/restaurants", getAllRestaurants)
	//	// A route for menus
	//	r.With(m.Pagination.Get("/menus", getAllMenu)
	//})
	// A route for restaurants
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/restaurants", restaurant.GetAllRestaurants)
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

//// Pagination Object
//type Pagination struct {
//	Next          int
//	Previous      int
//	RecordPerPage int
//	CurrentPage   int
//	TotalPage     int
//}

//// Pagination middleware is used to extract the next page id from the url query
//func Pagination(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		PageID := r.URL.Query().Get(string(PageIDKey))
//		intPageID := 0
//		var err error
//		if PageID != "" {
//			intPageID, err = strconv.Atoi(PageID)
//			if err != nil {
//				_ = render.Render(w, r, types.ErrInvalidRequest(fmt.Errorf("couldn't read %s: %w", PageIDKey, err)))
//				return
//			}
//		}
//		ctx := context.WithValue(r.Context(), PageIDKey, intPageID)
//		next.ServeHTTP(w, r.WithContext(ctx))
//	})
//}

type ErrResponse struct {
	Err            error
	HTTPStatusCode int
	StatusText     string
}

func ErrInvalidRequest(err error) *ErrResponse {
	return &ErrResponse{Err: err, HTTPStatusCode: http.StatusBadRequest, StatusText: "Invalid request."}
}
