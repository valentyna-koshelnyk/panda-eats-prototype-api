package restaurant

import (
	"github.com/go-chi/chi/v5"
)

// @title PandaEats API
// @version 1.0
// @description This is a demo for PandaEats API app
// @host https://localhost
// @BasePath /v1
// Routes is a router for restaurants
func Routes(c RestaurantController) chi.Router {
	r := chi.NewRouter()
	r.Get("/", c.GetAll)
	r.Delete("/{restaurant_id}", c.Delete)
	r.Put("/", c.Update)
	return r
}
