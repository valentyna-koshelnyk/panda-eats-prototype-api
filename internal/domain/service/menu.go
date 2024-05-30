package service

//go:generate mockery --name=MenuService
import (
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/repository"
)

// MenuService interface layer for menu service
type MenuService interface {
	GetMenu(id int64) (*entity.Response, error)
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
func (s *menuService) GetMenu(id int64) (*entity.Response, error) {
	menus, err := s.repository.GetMenu(id)
	if err != nil {
		return nil, err
	}

	if len(menus) == 0 {
		return entity.NewResponse([]entity.Item{}), nil
	}

	var items []entity.Item
	for _, menu := range menus {
		items = append(items, menu)
	}
	response := entity.NewResponse(items)
	return response, nil
}
