package repository

import (
	"gorm.io/gorm"

	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/utils"
)

//go:generate mockery --name=MenuRepository

// MenuRepository represents methods for interacting with db, menu entity
type MenuRepository interface {
	GetMenu(id int, pagination *utils.Pagination) (*utils.Pagination, error)
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
func (r *menuRepository) GetMenu(restaurantID int, pagination *utils.Pagination) (*utils.Pagination, error) {
	var menuList []entity.Menu
	err := r.db.
		Where("restaurant_id = ?", restaurantID).
		Scopes(utils.Paginate(entity.Menu{RestaurantID: restaurantID}, pagination, r.db)).
		Find(&menuList).Error
	if err != nil {
		return nil, err
	}

	pagination.Rows = menuList

	return pagination, nil
}
