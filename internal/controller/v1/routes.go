package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1/restaurant"
)

// Routes mounts routes of v1 API
func Routes() chi.Router {
	r := chi.NewRouter()
	r.Mount("/restaurants", restaurant.Routes())
	return r
}
