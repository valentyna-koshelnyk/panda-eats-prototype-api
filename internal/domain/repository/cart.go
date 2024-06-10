package repository

import (
	"github.com/guregu/dynamo"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
)

// CartRepository serves a repository for cart, layer for interaction with dynamoDB
type CartRepository interface {
	AddItem(cart entity.Cart) error
}

// cartRepository struct which takes dynamo table as attribute (orm)
type cartRepository struct {
	table *dynamo.Table
}

// NewCartRepository constructor for cart repository
func NewCartRepository(table *dynamo.Table) CartRepository {
	return &cartRepository{
		table: table,
	}
}

// AddItem adds item to dynamodb table 'cart'
// TODO: to fix unmarshalling of entity
func (c *cartRepository) AddItem(cart entity.Cart) error {
	err := c.table.Put(entity.Cart{
		UserID:       cart.UserID,
		ItemID:       cart.ItemID,
		Item:         cart.Item,
		Quantity:     cart.Quantity,
		PricePerUnit: cart.PricePerUnit,
		TotalPrice:   cart.TotalPrice,
		AddedAt:      cart.AddedAt,
	}).Run()
	if err != nil {
		return err
	}
	return nil
}
