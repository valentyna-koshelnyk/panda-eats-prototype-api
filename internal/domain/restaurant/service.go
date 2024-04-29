package restaurant

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// RestaurantService defines an API for restaurant service to be used by presentation layer
type RestaurantService interface {
	// FindAll fetches all restaurants list
	FindAll() ([]Restaurant, error)
	// FindById fetches restaurant by Id
	FindById(id int64) (*Restaurant, error)
	// GetByCategory fetches restaurants by category
	GetByCategory(categoryId string) ([]Restaurant, error)
	// GetByPriceRange fetches restaurants by price range
	GetByPriceRange(priceRange string) ([]Restaurant, error)
	// GetByZipCode fetches restaurants by Zip Code
	GetByZipCode(zipCode string) ([]Restaurant, error)
	//FindByCategoryPriceZip retrieves restaurants by category, price and zip
	FindByCategoryPriceZip(category string, priceRange string, zip string) ([]Restaurant, error)
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

// FindAll gets the list of all restaurants
func (service RestaurantServiceImpl) FindAll() ([]Restaurant, error) {
	if service.Restaurants == nil {
		if err := service.loadRestaurants(); err != nil {
			return nil, err
		}
	}
	return service.Restaurants, nil
}

// FindById fetches the restaurant information by restaurant id
func (service *RestaurantServiceImpl) FindById(id int64) (*Restaurant, error) {
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

func (service *RestaurantServiceImpl) FindByCategoryPriceZip(category string, priceRange string, zipCode string) []Restaurant {
	if service.Restaurants == nil {
		if err := service.loadRestaurants(); err != nil {
			return nil
		}
	}
	var restaurants []Restaurant
	for _, restaurant := range service.Restaurants {
		if restaurant.Category == category && restaurant.ZipCode == zipCode && restaurant.PriceRange == priceRange {
			restaurants = append(restaurants, restaurant)
			return []Restaurant{restaurant}
		}
	}
	if len(restaurants) == 0 {
		return nil
	}
	return restaurants
}
