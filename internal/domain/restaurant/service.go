package restaurant

import (
	"encoding/json"
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

// Cache restaurant list after the first load
type RestaurantServiceImpl struct {
	Restaurants []Restaurant
}

const restaurantJSON = "internal/data/restaurants.json"

// loadRestaurants returns list of restaurants
func (service *RestaurantServiceImpl) loadRestaurants() error {
	data, err := os.ReadFile(restaurantJSON)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &service.Restaurants)
	if err != nil {
		return err
	}
	return nil
}

func (service RestaurantServiceImpl) GetAll() ([]Restaurant, error) {
	if service.Restaurants == nil {
		if err := service.loadRestaurants(); err != nil {
			return nil, err
		}
	}
	return service.Restaurants, nil
}
