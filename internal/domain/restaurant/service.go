package restaurant

import (
	"encoding/json"
	"github.com/spf13/viper"
	"os"
	"strings"
)

// Service defines an API for restaurant service to be used by presentation layer
type Service interface {
	// FindAll fetches all restaurants list
	FindAll() ([]Restaurant, error)
	//FilterByCategoryPriceZip retrieves restaurants by category, price and zip
	FilterByCategoryPriceZip(category string, priceRange string, zip string) []Restaurant
}

// service Cache restaurant list after the first load
type service struct {
	Restaurants []Restaurant
}

// NewRestaurantService is a constructor with pointer to service struct which returned as instance of the interface
func NewRestaurantService() Service {
	return &service{}
}

var restaurantJSON = viper.GetString("paths.restaurants")

// loadRestaurants returns list of restaurants
func (service *service) loadRestaurants() error {
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
func (service service) FindAll() ([]Restaurant, error) {
	if service.Restaurants == nil {
		if err := service.loadRestaurants(); err != nil {
			return nil, err
		}
	}
	return service.Restaurants, nil
}

// FilterByCategoryPriceZip filters restaurants by category, price range or zip
func (service *service) FilterByCategoryPriceZip(category string, priceRange string, zipCode string) []Restaurant {
	if service.Restaurants == nil {
		if err := service.loadRestaurants(); err != nil {
			return nil
		}
	}
	var restaurants []Restaurant
	for _, restaurant := range service.Restaurants {
		matchCategory := category == "" || strings.Contains(restaurant.Category, category)
		matchZipCode := zipCode == "" || restaurant.ZipCode == zipCode
		matchPriceRange := priceRange == "" || restaurant.PriceRange == priceRange

		if matchCategory || matchZipCode || matchPriceRange {
			restaurants = append(restaurants, restaurant)
		}
	}

	if len(restaurants) == 0 {
		return nil
	}
	return restaurants
}
