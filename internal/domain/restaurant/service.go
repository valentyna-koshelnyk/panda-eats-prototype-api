package restaurant

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// Service defines an API for restaurant service to be used by presentation layer
type RestaurantService interface {
	// FindAll fetches all restaurants list
	FindAll() ([]Restaurant, error)
	// FindById fetches restaurant by Id
	FindById(id int64) (*Restaurant, error)
	//FindByCategoryPriceZip retrieves restaurants by category, price and zip
	FindByCategoryPriceZip(category string, priceRange string, zip string) ([]Restaurant, error)
}

// service Cache restaurant list after the first load
type restaurantService struct {
	Restaurants []Restaurant
}

func NewRestaurantService() RestaurantService {
	return &restaurantService{
		Restaurants: []Restaurant{},
	}
}

// loadRestaurants returns list of restaurants
func (service *restaurantService) loadRestaurants() error {
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
func (service *restaurantService) FindAll() ([]Restaurant, error) {
	if service.Restaurants == nil {
		if err := service.loadRestaurants(); err != nil {
			return nil, err
		}
	}
	return service.Restaurants, nil
}

// FindById fetches the restaurant information by restaurant id
func (service *restaurantService) FindById(id int64) (*Restaurant, error) {
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

func (service *restaurantService) FindByCategoryPriceZip(category string, priceRange string, zipCode string) ([]Restaurant, error) {
	if service.Restaurants == nil {
		if err := service.loadRestaurants(); err != nil {
			return nil, err
		}
	}
	var restaurants []Restaurant
	for _, restaurant := range service.Restaurants {
		if restaurant.Category == category && restaurant.ZipCode == zipCode && restaurant.PriceRange == priceRange {
			restaurants = append(restaurants, restaurant)
			return []Restaurant{restaurant}, nil
		}
	}
	if len(restaurants) == 0 {
		return nil, errors.New("restaurants not found")
	}
	return restaurants, nil
}
