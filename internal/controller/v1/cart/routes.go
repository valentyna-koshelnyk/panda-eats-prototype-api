package cart

import (
	"github.com/go-chi/chi/v5"
)

// Routes for cart
func Routes(c CartController) chi.Router {
	r := chi.NewRouter()
	r.Post("/item/{item_id}", c.AddItem)
	r.Get("/", c.GetCartItems)
	r.Delete("/item/{item_id}", c.RemoveItem)
	r.Patch("/item/{item_id}", c.UpdateItem)

	return r
}
