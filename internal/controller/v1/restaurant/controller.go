package restaurant

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	ce "github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/custom-errors"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/service"
	"io"
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
	w.Header().Set("Content-Type", "application/json")

	category := r.URL.Query().Get("category")
	priceRange := r.URL.Query().Get("price_range")
	zipCode := r.URL.Query().Get("zip_code")

	response, err := c.s.FilterRestaurants(category, zipCode, priceRange)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("Error getting restaurants: %s", err.Error())
		ce.RespondWithError(w, r, "error getting restaurants")
		return
	}
	if response == nil || len(response.Data.Items) == 0 {
		w.WriteHeader(http.StatusNotFound)
		log.Errorf("Restaurant not found: %s", err.Error())
		ce.RespondWithError(w, r, "restaurant data not found")
		return
	}
	w.WriteHeader(http.StatusOK)

	render.JSON(w, r, response)
	return
}

// @Summary Adds a new restaurant
// @Description Adds a new restaurant to the restaurants table
// @Accept  json
// @Produce  json
// @Param restaurant body restaurant.Restaurant true "Restaurant"
func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var restaurant entity.Restaurant
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("Error reading request body: %s", err.Error())
		ce.RespondWithError(w, r, "invalid request body")
		return
	}
	err = json.Unmarshal(data, &restaurant)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Error("Error decoding restaurant", err.Error())
		ce.RespondWithError(w, r, "error creating restaurant")
		return
	}

	err = c.s.CreateRestaurant(restaurant)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		log.Error("Error creating restaurant", err.Error())
		ce.RespondWithError(w, r, "error creating restaurant")
		return
	}
	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, "restaurant created")
	return
}

// @Summary Updates a restaurant information
// @Description Updates info about the restaurant
// @Accept  json
// @Produce  json
// @Param restaurant body restaurant.Restaurant true "Restaurant"
func (c *Controller) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var restaurant entity.Restaurant
	data, err := io.ReadAll(r.Body)
	err = json.Unmarshal(data, &restaurant)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		log.Error("Error decoding restaurant", err.Error())
		ce.RespondWithError(w, r, "error updating restaurant")
		return
	}

	err = c.s.UpdateRestaurant(restaurant)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("Error updating restaurant: %s", err.Error())
		ce.RespondWithError(w, r, "Error updating restaurant")
		return
	}
	w.WriteHeader(http.StatusNoContent)
	render.JSON(w, r, "successfully updated the restaurant")
	return
}

// @Summary Deletes a restaurant record
// @Description Deletes the restaurant
// @Accept  json
// @Produce  json
func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := chi.URLParam(r, "restaurant_id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	err := c.s.DeleteRestaurant(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("Error deleting restaurant: %s", err.Error())
		ce.RespondWithError(w, r, "error deleting restaurant")
		return
	}

	w.WriteHeader(http.StatusNoContent)
	render.JSON(w, r, "Restaurant deleted successfully")
	return
}
