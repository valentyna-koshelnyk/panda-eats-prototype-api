package menu

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// Service defines an API for menu service to be used by presentation layer
type MenuService interface {
	// FindAll fetches all dishes (menus) list
	FindAll() ([]Menu, error)
	// FindByRestaurantId fetches menu by restaurant Id
	FindByRestaurantId(restaurantId int64) ([]Menu, error)
}

// Cache menus list after the first load
type menuService struct {
	Menus []Menu
}

func NewMenuService() MenuService {
	return &menuService{
		Menus: []Menu{},
	}
}

func (service *menuService) FindAll() ([]Menu, error) {
	if service.Menus == nil {
		if err := service.loadMenus(); err != nil {
			return nil, err
		}
	}
	return service.Menus, nil
}

func (service *menuService) FindByRestaurantId(id int64) ([]Menu, error) {
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
func (service *menuService) loadMenus() error {
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
