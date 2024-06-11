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
	s service.CartService
}

// CartController interface for handlers of different CRUD operations for cart
type CartController interface {
	AddItem(w http.ResponseWriter, r *http.Request)
}

// NewCartController is a contrsuctor for cart
func NewCartController(s service.CartService) Controller {
	return Controller{s: s}
}

// AddItem is a handler for adding item to user's cart
func (c *Controller) AddItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID := chi.URLParam(r, "user_id")
	itemID := chi.URLParam(r, "item_id")

	var request utils.AddItemRequest

	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(data, &request)

	err = c.s.AddItem(userID, itemID, request.Quantity)
	if err != nil {
		utils.RespondWithJSON(w, r, "", err.Error())
		http.Error(w, "", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusCreated)
	utils.RespondWithJSON(w, r, "Item added", "")
	return
}

// GetItem a handler for retrieval list of dishes from user's cart
func (c *Controller) GetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := chi.URLParam(r, "user_id")
	items, err := c.s.GetItemsList(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("Error getting items: %s", err.Error())
		utils.RespondWithJSON(w, r, "", "error getting items")
		return
	}

	if len(items) == 0 {
		w.WriteHeader(http.StatusNotFound)
		log.Errorf("Cart is empty")
		utils.RespondWithJSON(w, r, "", "cart is empty")
		return
	}

	response := utils.CustomResponse{Data: items}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, response)
	return
}

// RemoveItem is a handler for removing entire item from user's cart
func (c *Controller) RemoveItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID := chi.URLParam(r, "user_id")
	itemID := chi.URLParam(r, "item_id")

	err := c.s.RemoveItem(userID, itemID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("Error removing item: %s", err.Error())
		utils.RespondWithJSON(w, r, "", "error removing item")
		return
	}
	w.WriteHeader(http.StatusNoContent)
	render.JSON(w, r, "Item removed")
	return
}
