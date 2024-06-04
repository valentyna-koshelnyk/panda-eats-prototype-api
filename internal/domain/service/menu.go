package service

import (
	custom_errors "github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/custom-errors"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/repository"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/utils"
)

//go:generate mockery --name=MenuService

// MenuService interface layer for menu service
type MenuService interface {
	GetMenu(id int, limit, offset int) (*utils.Pagination, error)
}

// menuService  layer for menu
type menuService struct {
	repository repository.MenuRepository
}

// NewMenuService is a constructor for service layer of menu
func NewMenuService(r repository.MenuRepository) MenuService {
	return &menuService{repository: r}
}

// GetMenu retrieves menu of the specific restaurant
func (s *menuService) GetMenu(id int, limit, offset int) (*utils.Pagination, error) {
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
