package cart

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/spf13/viper"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/config"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/repository"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/service"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte(viper.GetString("secret.key")), nil)
}

// Routes for cart
func Routes() chi.Router {
	r := chi.NewRouter()

	table := config.InitDynamoSession()
	db := config.GetDB()

	userRepository := repository.NewUserRepository(db)
	userTokenService := service.NewTokenService()
	userAuthService := service.NewAuthService()
	userService := service.NewUserService(userRepository, userAuthService, userTokenService)

	menuRepository := repository.NewMenuRepository(db)
	menuService := service.NewMenuService(menuRepository)

	cartRepository := repository.NewCartRepository(&table)
	cartService := service.NewCartService(cartRepository, userService, menuService)
	controller := NewCartController(cartService, userTokenService)

	r.Use(middleware.Logger)
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))

		r.Post("/item/{item_id}", controller.AddItem)
		r.Get("/", controller.GetCartItems)
		r.Delete("/item/{item_id}", controller.RemoveItem)
		r.Patch("/item/{item_id}", controller.UpdateItem)
	})

	return r
}
