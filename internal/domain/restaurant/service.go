package restaurant

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// RestaurantService defines an API for restaurant service to be used by presentation layer
type RestaurantService interface {
	// GetAll fetches all restaurants list
	GetAll() ([]Restaurant, error)
	// GetById fetches restaurant by Id
	GetById(id int64) (*Restaurant, error)
	// GetByCategory fetches restaurants by category
	GetByCategory(categoryId string) ([]Restaurant, error)
	// GetByPriceRange fetches restaurants by price range
	GetByPriceRange(priceRange string) ([]Restaurant, error)
	// GetByZipCode fetches restaurants by Zip Code
	GetByZipCode(zipCode string) ([]Restaurant, error)
}

// RestaurantServiceImpl Cache restaurant list after the first load
type RestaurantServiceImpl struct {
	Restaurants []Restaurant
}

var restaurantJSON = viper.GetString("paths.restaurants")

// loadRestaurants returns list of restaurants
func (service *RestaurantServiceImpl) loadRestaurants() error {
	x := viper.GetString("paths.restaurants")

	data, err := os.ReadFile(x)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &service.Restaurants)
	if err != nil {
		return err
	}
	return nil
}

// GetAll gets the list of all restaurants
func (service RestaurantServiceImpl) GetAll() ([]Restaurant, error) {
	if service.Restaurants == nil {
		if err := service.loadRestaurants(); err != nil {
			return nil, err
		}
	}
	return service.Restaurants, nil
}

// GetById fetches the restaurant information by restaurant id
func (service *RestaurantServiceImpl) GetById(id int64) (*Restaurant, error) {
	if service.Restaurants == nil {
		if err := service.loadRestaurants(); err != nil {
			return nil, err
		}
	}

	for i, restaurant := range service.Restaurants {
		if id == restaurant.ID {
			return &service.Restaurants[i], nil
		}
	}

	return nil, fmt.Errorf("restaurant with ID %d not found", id)
}
