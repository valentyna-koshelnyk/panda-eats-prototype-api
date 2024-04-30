package menu

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// Service Cache menus list after the first load
type Service struct {
	Menus []Menu
}

// FindByRestaurantID fetches menu of the restaurant by restaurant Id and returns the list of dishes
func (service *Service) FindByRestaurantID(id int64) ([]Menu, error) {
	if service.Menus == nil {
		if err := service.loadMenus(); err != nil {
			return nil, err
		}
	}
	var menus []Menu
	for _, menu := range service.Menus {
		if id == menu.RestaurantID {
			menus = append(menus, menu)
		}
	}
	if len(menus) == 0 {
		return nil, fmt.Errorf("no menu found for restaurant with ID %d", id)
	}
	return menus, nil
}

// loadMenus reads and deserializes the menus from JSON file
func (service *Service) loadMenus() error {
	x := viper.GetString("paths.menu")
	data, err := os.ReadFile(x)

	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &service.Menus)
	if err != nil {
		return err
	}
	return nil
}
