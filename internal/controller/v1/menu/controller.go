package menu

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"

	custom_errors "github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/custom-errors"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/service"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/utils"
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

	var err error
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithJSON(w, r, "", "no restaurant found")
		http.Error(w, "Invalid restaurant ID", http.StatusBadRequest)
		return
	}

	queryParams := r.URL.Query()
	limit, err := strconv.Atoi(queryParams.Get("limit"))
	if err != nil {
		limit = utils.DefaultLimit
	}

	offset, err := strconv.Atoi(queryParams.Get("offset"))
	if err != nil {
		offset = utils.DefaultOffset
	}

	pagedMenu, err := c.s.GetMenu(id, limit, offset)
	if err != nil && !errors.Is(err, custom_errors.ErrNotFound) {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorf("Error getting menu: %s", err)
		utils.RespondWithJSON(w, r, "", "error getting menu")
		return
	}

	if errors.Is(err, custom_errors.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		log.Errorf("No items available")
		utils.RespondWithJSON(w, r, "", "no items available")
		return
	}

	response := utils.NewPaginatedResponse(pagedMenu)

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, response)
}
