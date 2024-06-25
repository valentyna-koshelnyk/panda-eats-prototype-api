package server

import (
	"context"
	"github.com/spf13/viper"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/config"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1/cart"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1/menu"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1/order"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1/restaurant"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1/user"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/repository"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/service"
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

func SetupApplication() *Server {
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
	return CreateNewServer(port, controllers)
}
