package restaurant

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/restaurant"
	"net/http"
	"strconv"
)

// GetRestaurantByID a handler for retrieving information about the restaurant based on its id
func GetRestaurantByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	service := restaurant.NewRestaurantService()
	id, err := strconv.ParseInt(idStr, 10, 64)

	res, err := service.FindByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		return
	}
}
