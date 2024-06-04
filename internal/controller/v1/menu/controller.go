package menu

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"

	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/service"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/utils"
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
		utils.RespondWithJSON(w, r, "", "no restaurant found")
		http.Error(w, "Invalid restaurant ID", http.StatusBadRequest)
		return
	}

	response, err := c.s.GetMenu(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorf("Error getting menu: %s", err)
		utils.RespondWithJSON(w, r, "", "error getting menu")
		return
	}
	if response == nil || len(response.Data.Items) == 0 {
		w.WriteHeader(http.StatusNotFound)
		log.Errorf("No items available")
		utils.RespondWithJSON(w, r, "", "no items available")
		return
	}
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, response)
}
