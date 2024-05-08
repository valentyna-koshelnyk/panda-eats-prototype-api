package service

import (
	e "github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	r "github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/repository"
)

// Service defines an API for restaurant service to be used by presentation layer
type RestaurantService interface {
	FilterRestaurants(category string, zip string, priceRange string) ([]e.Restaurant, error)
	CreateRestaurant(restaurant e.Restaurant) error
	UpdateRestaurant(restaurant e.Restaurant) error
	DeleteRestaurant(id int64) error
}

// service gets repository
type restaurantService struct {
	repository r.RestaurantRepository
}

// NewRestaurantService is a constructor with pointer to service struct which returned as instance of the interface
func NewRestaurantService(r r.RestaurantRepository) RestaurantService {
	return &restaurantService{repository: r}
}

func (s *restaurantService) FilterRestaurants(category string, zip string, priceRange string) ([]e.Restaurant, error) {
	return s.repository.FilterRestaurants(category, zip, priceRange)
}

func (s *restaurantService) CreateRestaurant(restaurant e.Restaurant) error {
	err := s.repository.Create(restaurant)
	if err != nil {
		return err
	}
	return nil
}

func (s *restaurantService) UpdateRestaurant(restaurant e.Restaurant) error {
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
