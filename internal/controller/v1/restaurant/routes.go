package restaurant

import (
	"github.com/go-chi/chi/v5"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/config"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1/menu"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/repository"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/service"
)

// @title PandaEats API
// @version 1.0
// @description This is a demo for PandaEats API app
// @host https://localhost
// @BasePath /v1
// Routes is a router for restaurants
func Routes() chi.Router {
	r := chi.NewRouter()
	db := config.GetDB()

	restaurantRepository := repository.NewRestaurantRepository(db)
	restaurantService := service.NewRestaurantService(restaurantRepository)
	restaurantController := NewRestaurantController(restaurantService)

	menuRepository := repository.NewMenuRepository(db)
	menuService := service.NewMenuService(menuRepository)
	menuController := menu.NewController(menuService)

	r.Get("/", restaurantController.GetAll)
	r.Get("/{restaurant_id}/items", menuController.GetMenuByRestaurant)
	r.Post("/", restaurantController.Create)
	r.Delete("/{restaurant_id}", restaurantController.Delete)
	r.Put("/", restaurantController.Update)
	return r
}
