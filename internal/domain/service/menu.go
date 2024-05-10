package service

import (
	e "github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	r "github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/repository"
)

// Service  layer for menu
type MenuService struct {
	repository r.MenuRepository
}

// NewService is a constructor for service layer of menu
func NewMenuService(r r.MenuRepository) MenuService {
	return MenuService{repository: r}
}

// GetMenu retrieves menu of the specific restaurant
func (s *MenuService) GetMenu(id int64) (*[]e.Menu, error) {
	return s.repository.GetMenu(id)
}

//func (s *restaurantService) FindRestaurantByItem(item string) (*[] e.Restaurant, error) {
//	err := s.repository.
//}
