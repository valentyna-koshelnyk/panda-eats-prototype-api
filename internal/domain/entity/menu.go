package entity

import (
	"time"
)

// A Menu represents menu entity which is a menu of the restaurant
type Menu struct {
	ID           string    `gorm:"primaryKey;autoIncrement:true"`
	RestaurantID int       `json:"restaurant_id" gorm:"index:idx_restaurant_id"`
	Category     string    `json:"category"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Price        string    `json:"price"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}
