package menu

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/menu"
	"net/http"
	"strconv"
)

// GetMenuByRestaurant is a handler for getting dish fetched by restaurant Id
func GetMenuByRestaurant(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "restaurant_id")
	service := menu.Service{}
	id, err := strconv.ParseInt(idStr, 10, 64)

	m, err := service.FindByRestaurantID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(m)
	if err != nil {
		return
	}
}
