package restaurant

import (
	log "github.com/sirupsen/logrus"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/config"
	"gorm.io/gorm"
)

var db = config.InitDB()

type Repository interface {
	Create(restaurant Restaurant) error
	GetAll() ([]Restaurant, error)
	FilterRestaurants(category string, zip string, priceRange string) ([]Restaurant, error)
	Update(res Restaurant) error
	Delete(id int64) error
}

// restaurantRepository
type repository struct {
	db *gorm.DB
}

func NewRestaurantRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// Create TODO: to add dto for restaurant object
// Create inserts a new record into the database
func (r *repository) Create(restaurant Restaurant) error {
	result := db.Create(&restaurant)
	if result.Error != nil {
		log.Error(result.Error)
	}
	return nil
}

// GetAll retrieves all restaurants
func (r *repository) GetAll() ([]Restaurant, error) {
	var restaurants []Restaurant
	result := db.Find(&restaurants)
	return restaurants, result.Error
}

// FilterRestaurants gets the list of restaurants filtered by category, zip and price range
func (r *repository) FilterRestaurants(category string, zip string, priceRange string) ([]Restaurant, error) {
	var restaurants []Restaurant
	result := db.Where("category = ? OR zip_code LIKE ? OR price_range LIKE ?", category, zip, priceRange).Find(&restaurants)
	return restaurants, result.Error
}

// Update restaurant information or save new restaurant
func (r *repository) Update(res Restaurant) error {
	result := db.Save(&res)
	if result.Error != nil {
		log.Error(result.Error)
	}
	return nil
}

// Delete restaurant record
func (r *repository) Delete(id int64) error {
	db.Delete(&Restaurant{}, id)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
