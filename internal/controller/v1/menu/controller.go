package menu

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/menu"
	"net/http"
	"strconv"
)

// GetAllMenu is a handler for getting all menus
func GetAllMenus(w http.ResponseWriter, r *http.Request) {
	service := menu.NewMenuService()
	menus, err := service.FindAll()
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

// GetMenuByRestaurant is a handler for getting dish fetched by restaurant Id
func GetMenuByRestaurant(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "restaurant_id")
	service := menu.NewMenuService()
	id, err := strconv.ParseInt(idStr, 10, 64)

	menu, err := service.FindByRestaurantId(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(menu)
	if err != nil {
		return
	}
}
