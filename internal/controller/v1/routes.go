package v1

import (
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1/cart"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1/menu"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1/order"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1/restaurant"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/controller/v1/user"
)

// HTTPController is an object which handles all controllers as attributes for their initialisation at one entry point
type HTTPController struct {
	Menu       menu.MenuController
	Restaurant restaurant.RestaurantController
	Cart       cart.CartController
	Order      order.OrderController
	User       user.UserController
}

// Routes mounts routes of v1 API
func Routes(c *HTTPController) chi.Router {
	r := chi.NewRouter()

	r.Mount("/cart", cart.Routes(c.Cart))
	r.Mount("/menu", menu.Routes(c.Menu))
	r.Mount("/restaurants", restaurant.Routes(c.Restaurant))
	r.Mount("/auth", user.Routes(c.User))
	r.Mount("/order", order.Routes(c.Order))
	r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("doc.json")))
	return r
}
