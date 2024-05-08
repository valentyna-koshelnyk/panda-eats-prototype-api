package restaurant

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	e "github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	s "github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/service"
	"net/http"
	"strconv"
)

// Controller datatype for controller layer
type Controller struct {
	service s.RestaurantService
}

// NewRestaurantController constructor for controller layer
func NewRestaurantController(service s.RestaurantService) Controller {
	return Controller{service: service}
}

// @Summary Get all restaurants
// @Description Retrieves the list of all restaurants from the database
// @Produce  json
func (c *Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	priceRange := r.URL.Query().Get("price_range")
	zipCode := r.URL.Query().Get("zip_code")

	restaurants, err := c.service.FilterRestaurants(category, zipCode, priceRange)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.WithError(err).Error("Failed to fetch restaurants")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	restaurantsJSON, _ := json.Marshal(restaurants)
	_, err = w.Write([]byte(restaurantsJSON))
	if err != nil {
		log.WithError(err).Error("Failed to encode restaurants into JSON")
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// @Summary Adds a new restaurant
// @Description Adds a new restaurant to the restaurants table
// @Accept  json
// @Produce  json
// @Param restaurant body restaurant.Restaurant true "Restaurant"
func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	var res e.Restaurant
	err := c.service.CreateRestaurant(res)
	if err != nil {
		return
	}

	err = json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Restaurant %v", res)
}

// @Summary Updates a restaurant information
// @Description Updates info about the restaurant
// @Accept  json
// @Produce  json
// @Param restaurant body restaurant.Restaurant true "Restaurant"
func (c *Controller) Update(w http.ResponseWriter, r *http.Request) {
	var res e.Restaurant
	err := c.service.UpdateRestaurant(res)
	if err != nil {
		return
	}
	err = json.NewDecoder(r.Body).Decode(&res)

}

// @Summary Deletes a restaurant record
// @Description Deletes the restaurant
// @Accept  json
// @Produce  json
func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "restaurant_id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	err := c.service.DeleteRestaurant(id)
	if err != nil {
		return
	}
}
