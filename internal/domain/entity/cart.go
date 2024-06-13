package entity

import (
	"time"
)

// Cart is a struct for user's cart management, which entity values stored in dynamoDB table
type Cart struct {
	UserID       string    `dynamo:"user_id"`
	ItemID       string    `dynamo:"product_id"`
	Item         Menu      `dynamo:"item"`
	Quantity     int64     `dynamo:"quantity"`
	PricePerUnit float64   `dynamo:"price"`
	TotalPrice   float64   `dynamo:"total_price"`
	AddedAt      time.Time `dynamo:"added_at"`
}
