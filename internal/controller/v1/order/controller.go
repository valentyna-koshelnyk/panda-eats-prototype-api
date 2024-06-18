package order

import (
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/service"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/utils"
	"net/http"
)

// Controller type for order controller layer
type orderController struct {
	orderService service.OrderService
	tokenService service.TokenService
}

type OrderController interface {
	CreateOrder(w http.ResponseWriter, r *http.Request)
}

// NewController constructor for controller layer
func NewController(orderService service.OrderService, tokenService service.TokenService) OrderController {
	return &orderController{
		orderService: orderService,
		tokenService: tokenService,
	}
}

func (c *orderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	//token := jwtauth.TokenFromHeader(r)
	//userID, err := c.tokenService.ExtractEmailFromToken(token, viper.GetString("secret.key"))
	userEmail := "user@deliveryhero.com"
	w.Header().Set("Content-Type", "application/json")

	err := c.orderService.CreateOrder(userEmail)
	if err != nil {
		return
	}

	utils.RespondWithJSON(w, r, "Order created", "")
	return
}
