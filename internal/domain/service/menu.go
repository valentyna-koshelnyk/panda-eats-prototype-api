package service

//go:generate mockery --name=MenuService
import (
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/repository"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/utils"
)

// MenuService interface layer for menu service
type MenuService interface {
	GetMenu(id int64) (*utils.PaginatedResponse, error)
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
func (s *menuService) GetMenu(id int64) (*utils.PaginatedResponse, error) {
	menus, err := s.repository.GetMenu(id)
	if err != nil {
		return nil, err
	}

	if len(menus) == 0 {
		return utils.NewPaginatedResponse([]utils.Item{}), nil
	}

	var items []utils.Item
	for _, menu := range menus {
		items = append(items, menu)
	}
	response := utils.NewPaginatedResponse(items)
	return response, nil
}
