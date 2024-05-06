package restaurant

import (
	"github.com/go-chi/chi/v5"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/config"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1/menu"
	menu2 "github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/menu"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/restaurant"
)

// Routes is a router for restaurants
func Routes() chi.Router {
	r := chi.NewRouter()
	db := config.GetDB()
	repo := restaurant.NewRepository(db)
	service := restaurant.NewRestaurantService(repo)
	controller := NewRestaurantController(service)

	repoM := menu2.NewRepository(db)
	serviceM := menu2.NewService(repoM)
	controllerM := menu.NewController(serviceM)

	r.Get("/", controller.GetAll)
	r.Get("/{restaurant_id}/items", controllerM.GetMenuByRestaurant)
	r.Post("/", controller.Post)
	r.Delete("/{restaurant_id}", controller.Delete)
	r.Put("/", controller.Update)
	return r
}
