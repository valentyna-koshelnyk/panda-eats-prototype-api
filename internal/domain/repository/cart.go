package repository

import (
	"github.com/guregu/dynamo"
	log "github.com/sirupsen/logrus"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
)

// CartRepository serves a repository for cart, layer for interaction with dynamoDB
//
//go:generate mockery --name=CartRepository
type CartRepository interface {
	AddItem(cart entity.Cart) error
	GetCartItems(UserID string) ([]entity.Cart, error)
	GetCartItem(userID, itemID string) (*entity.Cart, error)
	RemoveItem(itemID, userID string) error
	UpdateCartItems(userID string, item entity.Cart) error
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

// GetCartItem retrieves a single item added by user to his cart
func (c *cartRepository) GetCartItem(userID, itemID string) (*entity.Cart, error) {
	var cart entity.Cart
	err := c.table.Get("user_id", userID).Range("product_id", dynamo.Equal, itemID).One(&cart)
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

// GetCartItems retrieves dishes that were added to cart of the user
func (c *cartRepository) GetCartItems(userID string) ([]entity.Cart, error) {
	var cart []entity.Cart

	err := c.table.Get("user_id", userID).All(&cart)
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

// UpdateCartItems updates quantity of item in the user's cart
func (c *cartRepository) UpdateCartItems(userID string, item entity.Cart) error {
	err := c.table.Update("user_id", userID).Range("product_id", item.Item.ID).Set("quantity", item.Quantity).Set("total_price", item.TotalPrice).Run()
	if err != nil {
		log.Printf("Failed to update cart item %s for user %s: %v", item.ItemID, userID, err)
		return err
	}
	log.Printf("Successfully updated cart item %s for user %s", item.ItemID, userID)
	return nil
}
