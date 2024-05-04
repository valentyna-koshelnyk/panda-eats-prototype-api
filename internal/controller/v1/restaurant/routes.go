package restaurant

import (
	"github.com/go-chi/chi/v5"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/config"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1/menu"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/restaurant"
)

var db = config.GetDB()

// Routes is a router for restaurants
func Routes() chi.Router {
	r := chi.NewRouter()
	restaurantService := restaurant.NewRestaurantService(restaurant.NewRestaurantRepository(db))

	controller := NewRestaurantController(restaurantService)

	r.Get("/", controller.GetAll)
	r.Get("/{restaurant_id}/items", menu.GetMenuByRestaurant)
	r.Get("/", controller.GetSelected)
	r.Post("/", controller.Post)
	r.Delete("/{restaurant_id}", controller.Delete)
	r.Put("/", controller.Update)
	return r
}
