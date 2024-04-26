package menu

import (
	"encoding/json"
	"errors"
	"github.com/spf13/viper"
	"os"
)

// MenuPath Add path to the configuration file
var MenuPath = viper.GetString("paths.menu")

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

func (service MenuServiceImpl) GetByRestaurantId(id int64) ([]Menu, error) {
	if service.Menus == nil {
		if err := service.loadMenus(); err != nil {
		}
	}
	dishesByRestaurantId := []Menu{}
	for _, menu := range service.Menus {
		if menu.MenuID == id {
			dishesByRestaurantId = append(dishesByRestaurantId, menu)
			return dishesByRestaurantId, nil
		}
	}
	return nil, errors.New("menu not found")
}

// loadMenus reads and deserializes the menus from JSON file
func (service *MenuServiceImpl) loadMenus() error {
	data, err := os.ReadFile(MenuPath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &service.Menus)
	if err != nil {
		return err
	}
	return nil
}
