package repository

import (
	"gorm.io/gorm"

	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
)

//go:generate mockery --name=MenuRepository

// MenuRepository represents methods for interacting with db, menu entity
type MenuRepository interface {
	GetMenu(id int64) ([]entity.Menu, error)
}

// menuRepository layer for interacting with db
type menuRepository struct {
	db *gorm.DB
}

// NewMenuRepository constructor for repository layer
func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &menuRepository{
		db: db,
	}
}

// GetMenu retrieves menu of the specific restaurant from the database
func (r *menuRepository) GetMenu(id int64) ([]entity.Menu, error) {
	var result []entity.Menu
	err := r.db.Where("restaurant_id = ?", id).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
