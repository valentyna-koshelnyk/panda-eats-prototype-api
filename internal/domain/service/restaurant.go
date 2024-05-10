package service

import (
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/repository"
)

// RestaurantService defines an API for restaurant service to be used by presentation layer
type RestaurantService interface {
	FilterRestaurants(category string, zip string, priceRange string) ([]entity.Restaurant, error)
	CreateRestaurant(restaurant entity.Restaurant) error
	UpdateRestaurant(restaurant entity.Restaurant) error
	DeleteRestaurant(id int64) error
}

// restaurantService gets repository
type restaurantService struct {
	repository repository.RestaurantRepository
}

// NewRestaurantService is a constructor with pointer to service struct which returned as instance of the interface
func NewRestaurantService(r repository.RestaurantRepository) RestaurantService {
	return &restaurantService{repository: r}
}

func (s *restaurantService) FilterRestaurants(category string, zip string, priceRange string) ([]entity.Restaurant, error) {
	return s.repository.FilterRestaurants(category, zip, priceRange)
}

func (s *restaurantService) CreateRestaurant(restaurant entity.Restaurant) error {
	err := s.repository.Create(restaurant)
	if err != nil {
		return err
	}
	return nil
}

func (s *restaurantService) UpdateRestaurant(restaurant entity.Restaurant) error {
	err := s.repository.Update(restaurant)
	if err != nil {
		return err
	}
	return nil
}

func (s *restaurantService) DeleteRestaurant(id int64) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
