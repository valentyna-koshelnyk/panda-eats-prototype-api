package restaurant

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"

	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/service"
)

// Controller datatype for controller layer
type restaurantController struct {
	s service.RestaurantService
}

type RestaurantController interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

// NewRestaurantController constructor for controller layer
func NewRestaurantController(service service.RestaurantService) RestaurantController {
	return &restaurantController{s: service}
}

// GetAll retrieves all restaurants
// @Summary Get all restaurants
// @Description Retrieves the list of all restaurants from the database
// @Produce  json
// @Success 200 {array} entity.Restaurant
// @Router  /restaurants [get]
func (c *restaurantController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	category := r.URL.Query().Get("category")
	priceRange := r.URL.Query().Get("price_range")
	zipCode := r.URL.Query().Get("zip_code")

	restaurants, err := c.s.FilterRestaurants(category, zipCode, priceRange)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("Error getting restaurants: %s", err.Error())
		entity.RespondWithJSON(w, r, "", "error getting restaurants")
		return
	}

	if len(restaurants) == 0 {
		w.WriteHeader(http.StatusNotFound)
		log.Errorf("Restaurant not found")
		entity.RespondWithJSON(w, r, "", "restaurant not found")
		return
	}

	response := entity.CustomResponse{Data: restaurants}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, response)
	return
}

// Create creates a restaurant
// @Summary Adds a new restaurant
// @Description Adds a new restaurant to the restaurants table
// @Accept  json
// @Produce  json
// @Success 201 {string} string "Restaurant created"
// @Failure 400 {string} string "Error reading request body or decoding restaurant"
// @Failure 422 {string} string "Error creating restaurant"
// @Router /restaurants [post]
func (c *restaurantController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var restaurant entity.Restaurant
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("Error reading request body: %s", err.Error())
		entity.RespondWithJSON(w, r, "", "error reading request body")
		return
	}
	err = json.Unmarshal(data, &restaurant)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Error("Error decoding restaurant", err.Error())
		entity.RespondWithJSON(w, r, "", "error decoding restaurant")
		return
	}

	err = c.s.CreateRestaurant(restaurant)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		log.Error("Error creating restaurant", err.Error())
		entity.RespondWithJSON(w, r, "", "error creating restaurant")
		return
	}
	w.WriteHeader(http.StatusCreated)
	entity.RespondWithJSON(w, r, "restaurant created", "")
	return
}

// Update updates restaurant information
// @Summary Updates a restaurant information
// @Description Updates info about the restaurant
// @Accept  json
// @Produce  json
// @Param id path string true "Restaurant ID"
// @Param restaurant body entity.Restaurant true "Restaurant"
// @Success 204 {string} string "Successfully updated restaurant"
// @Failure 400 {string} string "Error updating restaurant"
// @Failure 409 {string} string "Error decoding restaurant"
// @Router /restaurants/{id} [put]
func (c *restaurantController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var restaurant entity.Restaurant
	data, err := io.ReadAll(r.Body)
	err = json.Unmarshal(data, &restaurant)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		log.Error("Error decoding restaurant", err.Error())
		entity.RespondWithJSON(w, r, "", "error decoding restaurant")
		return
	}

	err = c.s.UpdateRestaurant(restaurant)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("Error updating restaurant: %s", err.Error())
		entity.RespondWithJSON(w, r, "", "error updating restaurant")
		return
	}
	w.WriteHeader(http.StatusNoContent)
	entity.RespondWithJSON(w, r, "successfully updated restaurant", "")
	return
}

// Delete deletes restaurant from the list
// @Summary Deletes a restaurant record
// @Description Deletes the restaurant
// @Produce  json
// @Param id   path      int  true  "Restaurant ID"
// @Router /restaurants/{id} [delete]
func (c *restaurantController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := chi.URLParam(r, "restaurant_id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	err := c.s.DeleteRestaurant(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("Error deleting restaurant: %s", err.Error())
		entity.RespondWithJSON(w, r, "", "error deleting restaurant")
		return
	}

	w.WriteHeader(http.StatusNoContent)
	entity.RespondWithJSON(w, r, "Restaurant deleted successfully", "restaurant deleted")
	return
}
