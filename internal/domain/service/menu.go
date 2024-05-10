package service

//go:generate mockery --name=MenuService
import (
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/repository"
)

type MenuService interface {
	GetMenu(id int64) (*[]entity.Menu, error)
}

// Service  layer for menu
type menuService struct {
	repository repository.MenuRepository
}

// NewService is a constructor for service layer of menu
func NewMenuService(r repository.MenuRepository) MenuService {
	return &menuService{repository: r}
}

// GetMenu retrieves menu of the specific restaurant
func (s *menuService) GetMenu(id int64) (*[]entity.Menu, error) {
	return s.repository.GetMenu(id)
}
