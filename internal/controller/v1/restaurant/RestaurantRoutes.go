package restaurant

import (
	"github.com/go-chi/chi/v5"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1/menu"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/repository"
)

// Routes is a router for restaurants
func Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", repository.GetAllRestaurants)
	r.Get("/{restaurant_id}/items", menu.GetMenuByRestaurant)
	r.Get("/", repository.GetByCategoryPriceZip)
	return r
}
