package entity

import "time"

// OrderStatus numerical value of order status
type OrderStatus int

// OrderStatus is enum representing actual status of order
const (
	InProcess OrderStatus = iota + 1
	Shipped
	Delivered
)

// String returns string value of order status
func (status OrderStatus) String() string {
	return [...]string{"InProcess", "Shipped", "Delivered"}[status-1]
}

// EnumIndex returns numerical value of order status
func (status OrderStatus) EnumIndex() int {
	return int(status)
}

// Order struct is takes items from user's cart and transforms to order
type Order struct {
	OrderID         string      `dynamo:"order_id"`
	UserID          string      `dynamo:"user_id"`
	Items           []Cart      `dynamo:"cart"`
	TotalOrderPrice float64     `dynamo:"total_price"`
	AddedAt         time.Time   `dynamo:"added_at"`
	Status          OrderStatus `dynamo:"status"`
}
