package restaurant

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/service"
	"net/http"
	"strconv"
)

// Controller datatype for controller layer
type Controller struct {
	s service.RestaurantService
}

// NewRestaurantController constructor for controller layer
func NewRestaurantController(service service.RestaurantService) Controller {
	return Controller{s: service}
}

// @Summary Get all restaurants
// @Description Retrieves the list of all restaurants from the database
// @Produce  json
func (c *Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	priceRange := r.URL.Query().Get("price_range")
	zipCode := r.URL.Query().Get("zip_code")

	w.Header().Set("Content-Type", "application/json")

	restaurants, err := c.s.FilterRestaurants(category, zipCode, priceRange)

	if err != nil {
		log.Errorf("Error getting restaurants: %s", err.Error())
	} else if len(restaurants) == 0 || restaurants == nil {
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte(""))
		if err != nil {
			log.Errorf("Error getting restaurants: %s", err.Error())
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
	} else {
		w.WriteHeader(http.StatusOK)
		restaurantsJSON, _ := json.Marshal(restaurants)
		_, err = w.Write([]byte(restaurantsJSON))
		if err != nil {
			log.WithError(err).Error("Failed to encode restaurants into JSON")
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

// @Summary Adds a new restaurant
// @Description Adds a new restaurant to the restaurants table
// @Accept  json
// @Produce  json
// @Param restaurant body restaurant.Restaurant true "Restaurant"
func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	var res entity.Restaurant
	err := json.NewDecoder(r.Body).Decode(&res)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		log.Error("Error decoding restaurant")
	}
	err = c.s.CreateRestaurant(res)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		log.Error("Error creating restaurant")
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Restaurant created")
}

// @Summary Updates a restaurant information
// @Description Updates info about the restaurant
// @Accept  json
// @Produce  json
// @Param restaurant body restaurant.Restaurant true "Restaurant"
func (c *Controller) Update(w http.ResponseWriter, r *http.Request) {
	var restaurant entity.Restaurant
	err := json.NewDecoder(r.Body).Decode(&restaurant)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = c.s.UpdateRestaurant(restaurant)
	if err != nil {
		log.Errorf("Error updating restaurant: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Restaurant updated successfully")
}

// @Summary Deletes a restaurant record
// @Description Deletes the restaurant
// @Accept  json
// @Produce  json
func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "restaurant_id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	err := c.s.DeleteRestaurant(id)
	if err != nil {
		log.Errorf("Error deleting restaurant: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Restaurant deleted successfully")
}
