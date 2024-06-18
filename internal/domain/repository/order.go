package repository

import (
	"github.com/guregu/dynamo"
	log "github.com/sirupsen/logrus"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
)

// OrderRepository crud operations for interacting with order table
type OrderRepository interface {
	CreateOrder(order entity.Order) error
	UpdateOrderStatus(order *entity.Order) error
	GetOrderInformation(userID, orderID string) (*entity.Order, error)
	GetOrderHistory(userID string) ([]entity.Order, error)
}

// orderRepository is a struct which takes dynamo orm as an attribute
type orderRepository struct {
	table *dynamo.Table
}

// NewOrderRepository is a constructor for order
func NewOrderRepository(table *dynamo.Table) OrderRepository {
	return &orderRepository{table: table}
}

// CreateOrder persists order in the order table
func (r *orderRepository) CreateOrder(order entity.Order) error {
	err := r.table.Put(entity.Order{
		OrderID:         order.OrderID,
		UserID:          order.UserID,
		Items:           order.Items,
		TotalOrderPrice: order.TotalOrderPrice,
		AddedAt:         order.AddedAt,
		Status:          order.Status,
	}).Run()
	if err != nil {
		return err
	}
	return nil
}

// UpdateOrderStatus updates order status
func (r *orderRepository) UpdateOrderStatus(order *entity.Order) error {
	err := r.table.
		Update("user_id", order.UserID).
		Range("order_id", order.OrderID).
		Set("status", order.Status).Run()
	if err != nil {
		log.Printf("Failed to update order status %s for order %s: %v", order.Status, order.OrderID, err)
		return err
	}
	log.Printf("Successfully updated order status as %s for order %s", order.Status, order.OrderID)
	return nil
}

// GetOrderInformation retrieves order object based on user/order ids
func (r *orderRepository) GetOrderInformation(userID, orderID string) (*entity.Order, error) {
	var order entity.Order
	err := r.table.
		Get("user_id", userID).
		Range("order_id", dynamo.Equal, orderID).
		One(&order)
	if err != nil {
		log.Error("order not found: ", err)
		return nil, err
	}
	return &order, nil
}

// GetOrderHistory retrieves orders from the table based on userID
func (r *orderRepository) GetOrderHistory(userID string) ([]entity.Order, error) {
	var orders []entity.Order

	err := r.table.Get("user_id", userID).All(&orders)
	if err != nil {
		log.Printf("Failed to retrieve order history for user %s: %v", userID, err)
		return nil, err
	}
	return orders, nil
}
