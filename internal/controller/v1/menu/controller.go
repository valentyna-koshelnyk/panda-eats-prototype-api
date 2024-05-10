package menu

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/service"

	"net/http"
	"strconv"
)

// Controller datatype for menu controller layer
type Controller struct {
	service service.MenuService
}

// NewController constructor for controller layer
func NewController(service service.MenuService) *Controller {
	return &Controller{
		service: service,
	}
}

// @Summary List menu of the restaurant
// @Description get menu by restaurantID
// @ID restaurant_id
// @Accept  json
// @Produce  json
// @Param id path int true "restaurant_id"
// @Success 200 {array} Menu
// @Router /restaurants/{restaurant_id}/items [get]
func (c *Controller) GetMenuByRestaurant(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "restaurant_id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid restaurant ID", http.StatusBadRequest)
		return
	}
	m, err := c.service.GetMenu(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	menuJSON, _ := json.Marshal(m)
	_, err = w.Write([]byte(menuJSON))
	if err != nil {
		log.WithError(err).Error("Failed to encode restaurants into JSON")
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
