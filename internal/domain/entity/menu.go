package entity

import (
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/config"
	"time"
)

// A Menu represents menu entity which is a menu of the restaurant
type Menu struct {
	ID           int64     `gorm:"primaryKey;autoIncrement:true"`
	RestaurantID int64     `json:"restaurant_id" gorm:"index:idx_restaurant_id"`
	Category     string    `json:"category"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Price        string    `json:"price"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

//TODO: to add validators to menu

func init() {
	db := config.GetDB()
	err := db.AutoMigrate(&Menu{})
	if err != nil {
		return
	}
}
