package order

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

// Routes for order controller
func Routes(c OrderController) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))
		r.Post("/", c.CreateOrder)
		r.Patch("/shipped", c.UpdateOrderStatusShipped)
		r.Patch("/deliver", c.UpdateOrderStatusDelivered)
		r.Get("/orders", c.GetOrdersHistory)
	})
	return r
}
