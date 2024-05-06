package restaurant

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/restaurant"
	"net/http"
	"strconv"
)

// Controller datatype for controller layer
type Controller struct {
	service restaurant.Service
}

// NewRestaurantController constructor for controller layer
func NewRestaurantController(service restaurant.Service) Controller {
	return Controller{service: service}
}

// GetAll retrieves all restaurants list OR filters restaurants by category/ price range/ zip code
func (c *Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	priceRange := r.URL.Query().Get("price_range")
	zipCode := r.URL.Query().Get("zip_code")
	restaurants, err := c.service.FilterRestaurants(category, priceRange, zipCode)
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

// Post handler for adding restaurant record
func (c *Controller) Post(w http.ResponseWriter, r *http.Request) {
	var res restaurant.Restaurant
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

// Update handler for update restaurant record
func (c *Controller) Update(w http.ResponseWriter, r *http.Request) {
	var res restaurant.Restaurant
	err := c.service.UpdateRestaurant(res)
	if err != nil {
		return
	}
	err = json.NewDecoder(r.Body).Decode(&res)

}

// Delete handler for delete restaurant record
func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "restaurant_id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	err := c.service.DeleteRestaurant(id)
	if err != nil {
		return
	}
}
