package menu

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// MenuPath Add path to the configuration file

// MenuService defines an API for menu service to be used by presentation layer
type MenuService interface {
	// GetAll fetches all dishes (menus) list
	GetAll() ([]*Menu, error)
	// GetByMenuId fetches menu by dish Id
	GetByMenuId(id int64) (*Menu, error)
	// GetByRestaurantId fetches menu by restaurant Id
	GetByRestaurantId(restaurantId int64) (*Menu, error)
}

// Cache menus list after the first load
type MenuServiceImpl struct {
	Menus []Menu
}

func (service MenuServiceImpl) GetAll() ([]Menu, error) {
	if service.Menus == nil {
		if err := service.loadMenus(); err != nil {
			return nil, err
		}
	}
	return service.Menus, nil
}

func (service *MenuServiceImpl) GetById(id int64) ([]Menu, error) {
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
func (service *MenuServiceImpl) loadMenus() error {
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
