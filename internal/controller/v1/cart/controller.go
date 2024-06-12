package cart

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/service"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/utils"
	"io"
	"net/http"
)

// Controller is struct for cart controller, which takes cart service as an attribute
type Controller struct {
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

// NewCartController is a contrsuctor for cart
func NewCartController(cartService service.CartService, tokenService service.TokenService) Controller {
	return Controller{
		cartService:  cartService,
		tokenService: tokenService,
	}
}

// AddItem is a handler for adding item to user's cart
func (c *Controller) AddItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Using hardcoded value as 1 before authentication is implemented, then "userID := chi.URLParam(r, "user_id")"
	userID := "1"
	itemID := chi.URLParam(r, "item_id")

	var request utils.QuantityItemRequest

	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(data, &request)

	err = c.cartService.AddItem(userID, itemID, request.Quantity)
	if err != nil {
		utils.RespondWithJSON(w, r, "", err.Error())
		http.Error(w, "", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusCreated)
	utils.RespondWithJSON(w, r, "Item added", "")
	return
}

// GetCartItems a handler for retrieval list of dishes from user's cart
func (c *Controller) GetCartItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Using hardcoded value as 1 before authentication is implemented, then "userID := chi.URLParam(r, "user_id")"
	userID := "1"
	// Get items from the cart
	items, err := c.cartService.GetItemsList(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorf("Error getting items: %s", err.Error())
		utils.RespondWithJSON(w, r, "", "error getting items")
		return
	}
	// Handle empty cart
	if len(items) == 0 {
		w.WriteHeader(http.StatusNoContent)
		log.Info("Cart is empty")
		utils.RespondWithJSON(w, r, "", "cart is empty")
		return
	}
	// Respond with cart items
	response := utils.CustomResponse{Data: items}
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, response)
}

// RemoveItem is a handler for removing entire item from user's cart
func (c *Controller) RemoveItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Hardcoded 1 before authentication is implemented
	//userID := chi.URLParam(r, "user_id")
	userID := "1"
	itemID := chi.URLParam(r, "item_id")

	err := c.cartService.RemoveItem(userID, itemID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("Error removing item: %s", err.Error())
		utils.RespondWithJSON(w, r, "", "error removing item")
		return
	}
	w.WriteHeader(http.StatusOK)
	utils.RespondWithJSON(w, r, "Item removed", "")
	return
}

// UpdateItem is a handler for updating item quantity
func (c *Controller) UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Hardcoded 1 before authentication is implemented
	//userID := chi.URLParam(r, "user_id")
	userID := "1"
	itemID := chi.URLParam(r, "item_id")

	var request utils.QuantityItemRequest

	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(data, &request)

	if err = c.cartService.UpdateUserItem(userID, itemID, request.Quantity); err != nil {
		utils.RespondWithJSON(w, r, "", err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	utils.RespondWithJSON(w, r, "Item updated", "")
	return
}
