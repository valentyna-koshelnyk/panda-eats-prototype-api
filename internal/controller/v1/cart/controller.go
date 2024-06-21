package cart

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/service"
	"io"
	"net/http"
)

// Controller is struct for cart controller, which takes cart service as an attribute
type cartController struct {
	cartService  service.CartService
	tokenService service.TokenService
}

// CartController interface for handlers of different CRUD operations for cart
type CartController interface {
	AddItem(w http.ResponseWriter, r *http.Request)
	GetCartItems(w http.ResponseWriter, r *http.Request)
	RemoveItem(w http.ResponseWriter, r *http.Request)
	UpdateItem(w http.ResponseWriter, r *http.Request)
}

// NewCartController is a constructor for cart
func NewCartController(cartService service.CartService, tokenService service.TokenService) CartController {
	return &cartController{
		cartService:  cartService,
		tokenService: tokenService,
	}
}

// AddItem is a handler for adding item to user's cart
func (c *cartController) AddItem(w http.ResponseWriter, r *http.Request) {
	token := jwtauth.TokenFromHeader(r)
	userID, err := c.tokenService.ExtractIDFromToken(token, viper.GetString("secret.key"))
	w.Header().Set("Content-Type", "application/json")
	itemID := chi.URLParam(r, "item_id")

	var request entity.QuantityItemRequest

	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(data, &request)

	err = c.cartService.AddItem(userID, itemID, request.Quantity)
	if err != nil {
		entity.RespondWithJSON(w, r, "", err.Error())
		http.Error(w, "", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusCreated)
	entity.RespondWithJSON(w, r, "Item added", "")
	return
}

// GetCartItems a handler for retrieval list of dishes from user's cart
func (c *cartController) GetCartItems(w http.ResponseWriter, r *http.Request) {
	token := jwtauth.TokenFromHeader(r)
	userID, err := c.tokenService.ExtractIDFromToken(token, viper.GetString("secret.key"))
	if err != nil {
		entity.RespondWithJSON(w, r, "", err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	items, err := c.cartService.GetItemsList(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorf("Error getting items: %s", err.Error())
		entity.RespondWithJSON(w, r, "", "error getting items")
		return
	}
	// Handle empty cart
	if len(items) == 0 {
		w.WriteHeader(http.StatusNoContent)
		log.Info("Cart is empty")
		entity.RespondWithJSON(w, r, "", "cart is empty")
		return
	}
	// Respond with cart items
	response := entity.CustomResponse{Data: items}
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, response)
}

// RemoveItem is a handler for removing entire item from user's cart
func (c *cartController) RemoveItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := jwtauth.TokenFromHeader(r)
	userID, err := c.tokenService.ExtractIDFromToken(token, viper.GetString("secret.key"))
	itemID := chi.URLParam(r, "item_id")
	err = c.cartService.RemoveItem(userID, itemID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("Error removing item: %s", err.Error())
		entity.RespondWithJSON(w, r, "", "error removing item")
		return
	}
	w.WriteHeader(http.StatusOK)
	entity.RespondWithJSON(w, r, "Item removed", "")
	return
}

// UpdateItem is a handler for updating item quantity
func (c *cartController) UpdateItem(w http.ResponseWriter, r *http.Request) {
	token := jwtauth.TokenFromHeader(r)
	userID, err := c.tokenService.ExtractIDFromToken(token, viper.GetString("secret.key"))
	w.Header().Set("Content-Type", "application/json")
	itemID := chi.URLParam(r, "item_id")

	var request entity.QuantityItemRequest

	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(data, &request)

	if err = c.cartService.UpdateUserItem(userID, itemID, request.Quantity); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		entity.RespondWithJSON(w, r, "", err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	entity.RespondWithJSON(w, r, "Item updated", "")
	return
}
