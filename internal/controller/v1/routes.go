package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1/cart"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1/menu"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1/order"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1/restaurant"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1/user"
)

// Routes mounts routes of v1 API
func Routes(c *controller.HTTPController) chi.Router {
	r := chi.NewRouter()
	var tokenAuth *jwtauth.JWTAuth
	key := []byte(viper.GetString("secret.key"))
	tokenAuth = jwtauth.New("HS256", key, nil)
	r.Use(middleware.Logger)
	r.Mount("/menu", menu.Routes(c.Menu))
	r.Mount("/restaurants", restaurant.Routes(c.Restaurant))
	r.Mount("/auth", user.Routes(c.User))
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))
		r.Mount("/order", order.Routes(c.Order))
		r.Mount("/cart", cart.Routes(c.Cart))
	})
	r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("doc.json")))
	return r
}
