package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1/menu"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1/restaurant"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/repository"
)

func RestaurantRoutes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", repository.PaginateHandler)
	r.Get("/{id}", restaurant.GetRestaurantById)
	r.Get("/{restaurant_id}/items", menu.GetMenuByRestaurant)
	return r
}
