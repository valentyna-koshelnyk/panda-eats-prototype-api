package menu

import "github.com/go-chi/chi/v5"

func Routes(c MenuController) chi.Router {
	r := chi.NewRouter()
	r.Get("restaurant/{restaurant_id}", c.GetMenuByRestaurant)
	return r
}
