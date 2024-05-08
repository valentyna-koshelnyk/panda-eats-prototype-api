package menu

import (
	"gorm.io/gorm"
)

// Repository represents methods for interacting with db, menu entity
type Repository interface {
	GetMenu(id int64) (*[]Menu, error)
}

// repository layer for interacting with db
type repository struct {
	db *gorm.DB
}

// NewRepository constructor for repository layer
func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// GetMenu retrieves menu of the specific restaurant from the database
func (r *repository) GetMenu(id int64) (*[]Menu, error) {
	var result []Menu
	err := r.db.Where("restaurant_id = ?", id).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}
