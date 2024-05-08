package repository

import (
	log "github.com/sirupsen/logrus"
	e "github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"gorm.io/gorm"
)

// Repository interface shows functions for interacting with DB
type RestaurantRepository interface {
	Create(restaurant e.Restaurant) error
	All() ([]e.Restaurant, error)
	FilterRestaurants(category string, zip string, priceRange string) ([]e.Restaurant, error)
	Update(res e.Restaurant) error
	Delete(id int64) error
}

// repository layer for interacting with db
type restaurantRepository struct {
	db *gorm.DB
}

// NewRepository constructor for repository layer
func NewRestaurantRepository(db *gorm.DB) RestaurantRepository {
	return &restaurantRepository{
		db: db,
	}
}

// Create TODO: to add dto for restaurant object
// Create inserts a new record into the database
func (r *restaurantRepository) Create(restaurant e.Restaurant) error {
	result := r.db.Create(&restaurant)
	if result.Error != nil {
		log.Error(result.Error)
	}
	return nil
}

// GetAll retrieves all restaurants
func (r *restaurantRepository) All() ([]e.Restaurant, error) {
	var restaurants []e.Restaurant
	result := r.db.Find(&restaurants)
	return restaurants, result.Error
}

// FilterRestaurants gets the list of restaurants filtered by category, zip and price range
func (r *restaurantRepository) FilterRestaurants(category string, zip string, priceRange string) ([]e.Restaurant, error) {
	var restaurants []e.Restaurant
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
func (r *restaurantRepository) Update(res e.Restaurant) error {
	result := r.db.Save(&res)
	if result.Error != nil {
		log.Error(result.Error)
	}
	return nil
}

// Delete restaurant record
func (r *restaurantRepository) Delete(id int64) error {
	r.db.Delete(&e.Restaurant{}, id)
	if r.db.Error != nil {
		return r.db.Error
	}
	return nil
}
