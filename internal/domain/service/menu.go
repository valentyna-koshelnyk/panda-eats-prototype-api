package service

import (
	custom_errors "github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/custom_errors"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/repository"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/utils"
)

//go:generate mockery --name=MenuService

// MenuService interface layer for menu service
type MenuService interface {
	GetRestaurantMenu(id int, limit, offset int) (*utils.Pagination, error)
	GetItem(itemID string) (*entity.Menu, error)
}

// menuService  layer for menu
type menuService struct {
	repository repository.MenuRepository
}

// NewMenuService is a constructor for service layer of menu
func NewMenuService(r repository.MenuRepository) MenuService {
	return &menuService{repository: r}
}

// GetRestaurantMenu retrieves menu of the specific restaurant
func (s *menuService) GetRestaurantMenu(id int, limit, offset int) (*utils.Pagination, error) {
	pagination := utils.Pagination{
		Limit: limit,
		Page:  offset,
	}

	pagedMenu, err := s.repository.GetMenu(id, &pagination)
	if err != nil {
		return nil, err
	}

	if pagedMenu.TotalRows == 0 {
		return nil, custom_errors.ErrNotFound
	}

	return pagedMenu, nil
}

// GetItem retrieves menu dish by dish(item) id
func (s *menuService) GetItem(itemID string) (*entity.Menu, error) {
	item, err := s.repository.GetItem(itemID)
	if err != nil {
		return nil, err
	}
	return &item, nil
}
