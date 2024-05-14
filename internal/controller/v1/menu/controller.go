package menu

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/service"
	"reflect"

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
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, "Invalid restaurant ID", http.StatusBadRequest)
		return
	}
	menuList, err := c.service.GetMenu(id)
	if err != nil {
		log.Errorf("Error while writing response: %s", err)
		http.Error(w, "Error retrieving menu", http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if menuList == nil || reflect.ValueOf(menuList).IsZero() {
		w.WriteHeader(http.StatusNotFound)
		_, err = w.Write([]byte(""))
		return
	}
	w.WriteHeader(http.StatusOK)
	menuJSON, _ := json.Marshal(menuList)
	_, err = w.Write([]byte(menuJSON))
	if err != nil {
		log.WithError(err).Error("Failed to encode restaurants into JSON")
		return
	}
}
