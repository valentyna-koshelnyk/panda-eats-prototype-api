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

type Controller struct {
	service restaurant.Service
}

func NewRestaurantController(service restaurant.Service) Controller {
	return Controller{service: service}
}

func (c *Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	//	service := restaurant.NewRestaurantService()
	restaurants, err := c.service.GetAllRestaurants()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.WithError(err).Error("Failed to fetch restaurants")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	restaurantsJSON, _ := json.Marshal(restaurants)
	w.Write(restaurantsJSON)
	if err != nil {
		return
	}
	{
		log.WithError(err).Error("Failed to encode restaurants into JSON")
		return
	}
}

func (c *Controller) GetSelected(w http.ResponseWriter, r *http.Request) {
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
	w.Write(restaurantsJSON)
	{
		log.WithError(err).Error("Failed to encode restaurants into JSON")
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

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
	//TODO: to update formatting of the output
	fmt.Fprintf(w, "Restaurant %v", res)
	if err != nil {
		return
	}
}

func (c *Controller) Update(w http.ResponseWriter, r *http.Request) {
	var res restaurant.Restaurant
	err := c.service.UpdateRestaurant(res)
	if err != nil {
		return
	}
	err = json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
	}
	return
}

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "restaurant_id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	err := c.service.DeleteRestaurant(id)
	if err != nil {
		return
	}
}
