package restaurant

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Repository interface shows functions for interacting with DB
type Repository interface {
	Create(restaurant Restaurant) error
	GetAll() ([]Restaurant, error)
	FilterRestaurants(category string, zip string, priceRange string) ([]Restaurant, error)
	Update(res Restaurant) error
	Delete(id int64) error
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

// Create TODO: to add dto for restaurant object
// Create inserts a new record into the database
func (r *repository) Create(restaurant Restaurant) error {
	result := r.db.Create(&restaurant)
	if result.Error != nil {
		log.Error(result.Error)
	}
	return nil
}

// GetAll retrieves all restaurants
func (r *repository) GetAll() ([]Restaurant, error) {
	var restaurants []Restaurant
	result := r.db.Find(&restaurants)
	return restaurants, result.Error
}

// FilterRestaurants gets the list of restaurants filtered by category, zip and price range
func (r *repository) FilterRestaurants(category string, zip string, priceRange string) ([]Restaurant, error) {
	var restaurants []Restaurant
	query := r.db
	if category != "" {
		query = query.Where("category = ?", category)
	}
	if zip != "" {
		query = query.Where("zip_code = ?", zip)
	}
	if priceRange != "" {
		query = query.Where("price_range = ?", priceRange)
	}

	result := query.Find(&restaurants)
	if result.Error != nil {
		log.Error("Failed to filter restaurants: ", result.Error)
		return nil, result.Error
	}
	return restaurants, nil
}

// Update restaurant information or save new restaurant
func (r *repository) Update(res Restaurant) error {
	result := r.db.Save(&res)
	if result.Error != nil {
		log.Error(result.Error)
	}
	return nil
}

// Delete restaurant record
func (r *repository) Delete(id int64) error {
	r.db.Delete(&Restaurant{}, id)
	if r.db.Error != nil {
		return r.db.Error
	}
	return nil
}
