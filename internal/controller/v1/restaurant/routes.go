package restaurant

import (
	"github.com/go-chi/chi/v5"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/config"
	m "github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1/menu"
	rpstr "github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/repository"
	s "github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/service"
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

	restaurantRepository := rpstr.NewRestaurantRepository(db)
	restaurantService := s.NewRestaurantService(restaurantRepository)
	restaurantController := NewRestaurantController(restaurantService)

	menuRepository := rpstr.NewMenuRepository(db)
	menuService := s.NewMenuService(menuRepository)
	menuController := m.NewController(menuService)

	r.Get("/", restaurantController.All)
	r.Get("/{restaurant_id}/items", menuController.MenuByRestaurant)
	r.Post("/", restaurantController.Create)
	r.Delete("/{restaurant_id}", restaurantController.Delete)
	r.Put("/", restaurantController.Update)
	return r
}
