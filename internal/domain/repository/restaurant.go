package repository

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
)

//go:generate mockery --name=RestaurantRepository

// RestaurantRepository interface shows functions for interacting with DB
type RestaurantRepository interface {
	Create(restaurant entity.Restaurant) error
	GetAll() ([]entity.Restaurant, error)
	FilterRestaurants(category string, zip string, priceRange string) ([]entity.Restaurant, error)
	Update(res entity.Restaurant) error
	Delete(id int64) error
}

// repository layer for interacting with db
type restaurantRepository struct {
	db *gorm.DB
}

// NewRestaurantRepository constructor for repository layer
func NewRestaurantRepository(db *gorm.DB) RestaurantRepository {
	return &restaurantRepository{
		db: db,
	}
}

// Create TODO: to add dto for restaurant object
// Create inserts a new record into the database
func (r *restaurantRepository) Create(restaurant entity.Restaurant) error {
	result := r.db.Create(&restaurant)
	if result.Error != nil {
		log.Error(result.Error)
	}
	return nil
}

// GetAll retrieves all restaurants
func (r *restaurantRepository) GetAll() ([]entity.Restaurant, error) {
	var restaurants []entity.Restaurant
	result := r.db.Find(&restaurants)
	return restaurants, result.Error
}

// FilterRestaurants gets the list of restaurants filtered by category, zip and price range
func (r *restaurantRepository) FilterRestaurants(category, zip, priceRange string) ([]entity.Restaurant, error) {
	var restaurants []entity.Restaurant
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
func (r *restaurantRepository) Update(res entity.Restaurant) error {
	result := r.db.Save(&res)
	if result.Error != nil {
		log.Error(result.Error)
	}
	return nil
}

// Delete restaurant record
func (r *restaurantRepository) Delete(id int64) error {
	r.db.Delete(&entity.Restaurant{}, id)
	if r.db.Error != nil {
		return r.db.Error
	}
	return nil
}
