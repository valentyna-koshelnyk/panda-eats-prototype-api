package repository

import (
	"github.com/guregu/dynamo"
	log "github.com/sirupsen/logrus"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
)

// CartRepository serves a repository for cart, layer for interaction with dynamoDB
type CartRepository interface {
	AddItem(cart entity.Cart) error
	GetCartItems(UserID string) ([]entity.Cart, error)
	RemoveItem(itemID, userID string) error
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

// GetCartItems retrieves dishes that were added to cart of the user
func (c *cartRepository) GetCartItems(userID string) ([]entity.Cart, error) {
	var cart []entity.Cart

	err := c.table.Scan().Filter("'user_id' = ?", userID).All(&cart)
	if err != nil {
		log.Printf("Failed to retrieve cart items for user %s: %v", userID, err)
		return nil, err
	}
	return cart, nil
}

// RemoveItem removes item from user's cart
func (c *cartRepository) RemoveItem(userID, itemID string) error {
	err := c.table.Delete("user_id", userID).Range("product_id", itemID).Run()
	if err != nil {
		log.Printf("Failed to delete cart item %s for user %s: %v", itemID, userID, err)
		return err
	}
	log.Printf("Successfully deleted cart item %s for user %s", itemID, userID)
	return nil
}
