package main

import (
	"context"
	"fmt"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1/cart"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1/menu"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1/order"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1/restaurant"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1/user"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/repository"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "github.com/valentyna-koshelnyk/panda-eats-prototype-api/docs"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/config"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/server"
)

var version string

func init() {
	// TODO: implement a custom structured logger
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	// Log for informational purposes, then depends on necessity use log.SetLevel(log.WarnLevel)
	log.SetLevel(log.InfoLevel)
	log.Info("starting up API...", log.WithField("version", version))

}

func main() {
	config.InitViperConfig()

	cartTable := config.InitDynamoSession("cart")
	orderTable := config.InitDynamoSession("order")
	db := config.GetDB()

	restaurantRepository := repository.NewRestaurantRepository(db)
	userRepository := repository.NewUserRepository(db)
	menuRepository := repository.NewMenuRepository(db)
	cartRepository := repository.NewCartRepository(&cartTable)
	orderRepository := repository.NewOrderRepository(&orderTable)

	userTokenService := service.NewTokenService(viper.GetString("secret.key"))
	userAuthService := service.NewAuthService()

	restaurantService := service.NewRestaurantService(restaurantRepository)
	userService := service.NewUserService(userRepository, userAuthService, userTokenService)
	menuService := service.NewMenuService(menuRepository)
	cartService := service.NewCartService(cartRepository, userService, menuService)
	orderService := service.NewOrderService(orderRepository, cartService, userService)

	controllers := new(controller.HTTPController)
	controllers.Restaurant = restaurant.NewRestaurantController(restaurantService)
	controllers.Cart = cart.NewCartController(cartService, userTokenService)
	controllers.User = user.NewUserController(userService)
	controllers.Order = order.NewController(orderService, userTokenService)
	controllers.Menu = menu.NewController(menuService)

	port := viper.GetString("server.port")
	srv := server.CreateNewServer(port, controllers)
	log.Info("Starting server on port :", port)

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
