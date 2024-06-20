package controller

import (
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
