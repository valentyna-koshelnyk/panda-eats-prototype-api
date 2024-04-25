package restaurant

import (
	"encoding/json"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/menu"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/restaurant"
	"net/http"
)

func GetAllMenu(w http.ResponseWriter, r *http.Request) {
	service := menu.MenuServiceImpl{}
	menus, err := service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(menus)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

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
