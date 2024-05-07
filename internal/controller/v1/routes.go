package v1

import (
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1/restaurant"
)

// Routes mounts routes of v1 API
func Routes() chi.Router {
	r := chi.NewRouter()
	r.Mount("/restaurants", restaurant.Routes())
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("doc.json"),
	))
	return r
}
