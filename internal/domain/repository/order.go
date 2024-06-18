package repository

import (
	"github.com/guregu/dynamo"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
)

type OrderRepository interface {
	CreateOrder(order entity.Order) error
}

type orderRepository struct {
	table *dynamo.Table
}

func NewOrderRepository(table *dynamo.Table) OrderRepository {
	return &orderRepository{table: table}
}

func (orderRepository *orderRepository) CreateOrder(order entity.Order) error {
	err := orderRepository.table.Put(entity.Order{
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
