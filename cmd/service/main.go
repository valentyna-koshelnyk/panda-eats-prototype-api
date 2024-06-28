package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "github.com/valentyna-koshelnyk/panda-eats-prototype-api/docs"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/server"
)

var version string

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	// Log for informational purposes, then depends on necessity use log.SetLevel(log.WarnLevel)
	log.SetLevel(log.InfoLevel)
	log.Info("starting up API...", log.WithField("version", version))

}

func main() {
	srv := server.SetupApplication()
	log.Info("Starting server on port :", viper.GetString("server.port"))

	go func() {
		fmt.Println("Starting HTTP server...")
		err := srv.Start()
		if err != nil && err != http.ErrServerClosed {
			fmt.Printf("HTTP server error: %s\n", err)
		}
	}()

	// Create a channel to receive  notifications from signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Context for shutdown process
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("HTTP server shutdown error: %v", err)
	}
}
