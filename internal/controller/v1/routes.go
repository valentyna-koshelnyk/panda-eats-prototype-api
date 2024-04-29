package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1/menu"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/repository"
)

// Define chi router, path to controllers, I can muntain (build a chain between multiple routers)
// I can chain to routers from main file to routes.go (creating chain of routers). in main we don't maintain a
// check go chi documentation on how to do it
func RestaurantRoutes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", repository.PaginateHandler)
	//	r.Get("/{id}", restaurant.GetRestaurantById)
	r.Get("/{restaurant_id}/items", menu.GetMenuByRestaurant)
	return r
}
