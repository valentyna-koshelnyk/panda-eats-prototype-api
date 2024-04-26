package restaurant

import (
	"encoding/json"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/restaurant"
	"net/http"
)

func GetAllRestaurants(w http.ResponseWriter, r *http.Request) {
	service := restaurant.RestaurantServiceImpl{}
	restaurants, err := service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(restaurants)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
