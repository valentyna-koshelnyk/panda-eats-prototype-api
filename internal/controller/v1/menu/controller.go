package menu

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	ce "github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/custom-errors"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/service"
	"net/http"
	"strconv"
)

// Controller type for menu controller layer
type Controller struct {
	s service.MenuService
}

// NewController constructor for controller layer
func NewController(s service.MenuService) *Controller {
	return &Controller{
		s: s,
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
	w.Header().Set("Content-Type", "application/json")
	idStr := chi.URLParam(r, "restaurant_id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid restaurant ID", http.StatusBadRequest)
		return
	}

	response, err := c.s.GetMenu(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorf("Error getting menu: %s", err)
		ce.RespondWithError(w, r, "error getting menu")
		return
	}
	if response == nil || len(response.Data.Items) == 0 {
		w.WriteHeader(http.StatusNotFound)
		log.Errorf("No items available")
		ce.RespondWithError(w, r, "restaurant data not found")
		return
	}
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, response)
}
