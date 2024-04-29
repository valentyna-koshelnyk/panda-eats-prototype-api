package restaurant

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// RestaurantService defines an API for restaurant service to be used by presentation layer
type restaurantService interface {
	// FindAll fetches all restaurants list
	FindAll() ([]Restaurant, error)
	// FindById fetches restaurant by Id
	FindById(id int64) (*Restaurant, error)
	//FindByCategoryPriceZip retrieves restaurants by category, price and zip
	FindByCategoryPriceZip(category string, priceRange string, zip string) ([]Restaurant, error)
}

// RestaurantService Cache restaurant list after the first load
type RestaurantService struct {
	Restaurants []Restaurant
}

var restaurantJSON = viper.GetString("paths.restaurants")

// loadRestaurants returns list of restaurants
func (service *RestaurantService) loadRestaurants() error {
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
func (service RestaurantService) FindAll() ([]Restaurant, error) {
	if service.Restaurants == nil {
		if err := service.loadRestaurants(); err != nil {
			return nil, err
		}
	}
	return service.Restaurants, nil
}

// FindById fetches the restaurant information by restaurant id
func (service *RestaurantService) FindById(id int64) (*Restaurant, error) {
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

func (service *RestaurantService) FindByCategoryPriceZip(category string, priceRange string, zipCode string) []Restaurant {
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
