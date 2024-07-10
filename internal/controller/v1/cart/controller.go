package cart

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
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

// AddItem adds item to user's cart
// @Summary      Adds items to the user's cart
// @Description user selects item from the menu of the restaurant and adds his to the cart by the item id
// @Produce      json
// @Success      201  {object}  entity.CustomResponse{data=string}
// @Failure      400  {object}  entity.CustomResponse{data=string}
// @Param        request body entity.QuantityItemRequest true "The input is quantity"
// @Param {item_id} path string true "item_id"
// @Router       /cart/item/{item_id} [post]
func (c *cartController) AddItem(w http.ResponseWriter, r *http.Request) {
	userID, err := c.tokenService.ExtractIDFromToken(r)
	w.Header().Set("Content-Type", "application/json")
	itemID := chi.URLParam(r, "item_id")

	var request entity.QuantityItemRequest

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Errorf("Error adding items: %s", err.Error())
		entity.RespondWithJSON(w, r, "", "error with formating quantity")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(data, &request)

	err = c.cartService.AddItem(userID, itemID, request.Quantity)
	if err != nil {
		log.Errorf("Error adding items: %s", err.Error())
		entity.RespondWithJSON(w, r, "", "error adding items")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	entity.RespondWithJSON(w, r, "Item added", "")
	return
}

// GetCartItems a handler for retrieval list of cart items from user's cart
// @Summary      Gets cart items
// @Description  retrieve the items that are currently in the user's cart and haven't been added to the order
// @Produce      json
// @Success      200  {object}  entity.CustomResponse{data=[]entity.Cart}
// @Failure      400  {object}  entity.CustomResponse(error=string)
// @Router       /cart/ [get]
func (c *cartController) GetCartItems(w http.ResponseWriter, r *http.Request) {
	userID, err := c.tokenService.ExtractIDFromToken(r)
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
	if len(items) == 0 {
		w.WriteHeader(http.StatusNoContent)
		log.Info("Cart is empty")
		entity.RespondWithJSON(w, r, "", "cart is empty")
		return
	}
	response := entity.CustomResponse{Data: items}
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, response)
}

// RemoveItem is a handler for removing entire item from user's cart
// @Summary      Removes item from user's cart
// @Description  remove item from the cart
// @Produce      json
// @Success      200  {object}  entity.CustomResponse{data=string}
// @Failure      400  {object}  entity.CustomResponse(error=string)
// @Param {item_id} path string true "item_id"
// @Router       /cart/item/{item_id} [delete]
func (c *cartController) RemoveItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID, err := c.tokenService.ExtractIDFromToken(r)
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
// @Summary      Updates item quantity
// @Description  updates  item quantity
// @Produce      json
// @Success      200  {object}  entity.CustomResponse{data=string}
// @Failure      400  {object}  entity.CustomResponse(error=string)
// @Param        request body entity.QuantityItemRequest true "The input is quantity"
// @Param {item_id} path string true "item_id"
// @Router       /cart/item/{item_id} [delete]
func (c *cartController) UpdateItem(w http.ResponseWriter, r *http.Request) {
	userID, err := c.tokenService.ExtractIDFromToken(r)
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
