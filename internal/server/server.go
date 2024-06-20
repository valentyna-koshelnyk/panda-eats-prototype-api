package server

import (
	"context"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	v1 "github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1"
)

// Server represents a chi router and HTTP server
type Server struct {
	Router     *chi.Mux
	httpServer *http.Server
}

// CreateNewServer should return a server struct
func CreateNewServer(port string, controllers *controller.HTTPController) *Server {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Mount("/api/v1", v1.Routes(controllers))

	server := &Server{
		Router: router,
		httpServer: &http.Server{
			Addr:    ":" + port,
			Handler: router,
		},
	}
	return server
}

// Start begins listening and serving HTTP requests.
func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

// Shutdown gracefully shuts down the server
func (s *Server) Shutdown(ctx context.Context) error {
	if s.httpServer != nil {
		return s.httpServer.Shutdown(ctx)
	}
	return nil
}
