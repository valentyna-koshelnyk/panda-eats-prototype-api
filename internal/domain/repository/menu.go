package repository

import (
	e "github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"gorm.io/gorm"
)

// Repository represents methods for interacting with db, menu entity
type MenuRepository interface {
	GetMenu(id int64) (*[]e.Menu, error)
}

// repository layer for interacting with db
type menuRepository struct {
	db *gorm.DB
}

// NewRepository constructor for repository layer
func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &menuRepository{
		db: db,
	}
}

// GetMenu retrieves menu of the specific restaurant from the database
func (r *menuRepository) GetMenu(id int64) (*[]e.Menu, error) {
	var result []e.Menu
	err := r.db.Where("restaurant_id = ?", id).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

//func (r *menuRepository) GetRestaurantByDishName(name string) (*[]e.Menu, error) {
//
//}
