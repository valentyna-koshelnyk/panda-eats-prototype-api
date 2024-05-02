package restaurant

import (
	"github.com/go-chi/chi/v5"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1/menu"
)

// Routes is a router for restaurants
func Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", GetAllRestaurants)
	r.Get("/{restaurant_id}/items", menu.GetMenuByRestaurant)
	r.Get("/", GetByCategoryPriceZip)
	return r
}
