package order

import (
	"encoding/json"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/service"
	"io"
	"net/http"
)

// orderController type for order controller layer
type orderController struct {
	orderService service.OrderService
	tokenService service.TokenService
}

// OrderController interface for orders
type OrderController interface {
	CreateOrder(w http.ResponseWriter, r *http.Request)
	UpdateOrderStatusShipped(w http.ResponseWriter, r *http.Request)
	UpdateOrderStatusDelivered(w http.ResponseWriter, r *http.Request)
	GetOrdersHistory(w http.ResponseWriter, r *http.Request)
}

// NewController constructor for controller layer
func NewController(orderService service.OrderService, tokenService service.TokenService) OrderController {
	return &orderController{
		orderService: orderService,
		tokenService: tokenService,
	}
}

// CreateOrder is a handler for placing the order
func (c *orderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	//token := jwtauth.TokenFromHeader(r)
	//userID, err := c.tokenService.ExtractEmailFromToken(token, viper.GetString("secret.key"))
	userID := "50aa4686-bb62-4202-b2ce-471df794adea"
	w.Header().Set("Content-Type", "application/json")

	err := c.orderService.CreateOrder(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("error creating order: %s", err)
		entity.RespondWithJSON(w, r, "", "error creating order")
		return
	}

	entity.RespondWithJSON(w, r, "Order created", "")
	return
}

// UpdateOrderStatusShipped updates status of the order to be set as shipped
func (c *orderController) UpdateOrderStatusShipped(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var orderID entity.OrderIDRequest
	userID := "50aa4686-bb62-4202-b2ce-471df794adea"
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("error reading body: %s", err)
		entity.RespondWithJSON(w, r, "", "error retrieving user")
		return
	}
	err = json.Unmarshal(data, &orderID)

	err = c.orderService.UpdateOrderStatusShipped(userID, orderID.OrderID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("error updating the order: %s", err)
		entity.RespondWithJSON(w, r, "", "error updating the order")
		return
	}

	w.WriteHeader(http.StatusAccepted)
	entity.RespondWithJSON(w, r, "Order has been shipped", "")
	return
}

// UpdateOrderStatusDelivered updates status of the order to be set as delivered
func (c *orderController) UpdateOrderStatusDelivered(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var orderID entity.OrderIDRequest
	userID := "50aa4686-bb62-4202-b2ce-471df794adea"
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("error reading body: %s", err)
		entity.RespondWithJSON(w, r, "", "error retrieving user")
		return
	}
	err = json.Unmarshal(data, &orderID)

	if err = c.orderService.UpdateOrderStatusDelivered(userID, orderID.OrderID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("error updating the order: %s", err)
		entity.RespondWithJSON(w, r, "", "error updating the order")
		return
	}

	w.WriteHeader(http.StatusAccepted)
	entity.RespondWithJSON(w, r, "Order has been delivered", "")
	return
}

// GetOrdersHistory retrieves orders history  of the particular user
func (c *orderController) GetOrdersHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := "50aa4686-bb62-4202-b2ce-471df794adea"

	orders, err := c.orderService.GetOrderHistory(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("error getting order history: %s", err.Error())
		entity.RespondWithJSON(w, r, "", "error getting order history")
		return
	}

	if len(orders) == 0 {
		w.WriteHeader(http.StatusNotFound)
		log.Errorf("o")
	}

	response := entity.CustomResponse{Data: orders}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, response)
	return

}
