package cart

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/spf13/viper"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte(viper.GetString("secret.key")), nil)
}

// Routes for cart
func Routes(c CartController) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))
		r.Post("/item/{item_id}", c.AddItem)
		r.Get("/", c.GetCartItems)
		r.Delete("/item/{item_id}", c.RemoveItem)
		r.Patch("/item/{item_id}", c.UpdateItem)
	})

	return r
}
