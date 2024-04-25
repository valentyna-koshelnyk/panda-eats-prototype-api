package menu

import (
	"encoding/json"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/menu"
	"net/http"
)

// GetAllMenu is a handler for getting all menus
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

// GetMenusByRestaurant is a handler for getting menu fetched by restaurant Id
func GetMenusByRestaurant(id int64, w http.ResponseWriter, r *http.Request) {
	service := menu.MenuServiceImpl{}
	restaurants, err := service.GetByRestaurantId(id)
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
