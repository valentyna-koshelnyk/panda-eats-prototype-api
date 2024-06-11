package service

import (
	"errors"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/repository"
	"strconv"
	"strings"
	"time"
)

// CartService is interface for cart service having methods for adding item and verification of user/item
type CartService interface {
	AddItem(userID, itemID string, quantity int64) error
	isVerifiedUserItem(userID, itemID string) (bool, error)
	GetItemsList(userID string) ([]entity.Cart, error)
	RemoveItem(userID, itemID string) error
	UpdateUserItem(userID, itemID string, quantity int64) error
}

// cartService is struct for cart
type cartService struct {
	cartRepository repository.CartRepository
	userService    UserService
	menuService    MenuService
}

// NewCartService constructor for cartservice with injected fields as repo, user service and menu service
func NewCartService(cartRepository repository.CartRepository, userService UserService, menuService MenuService) CartService {
	return &cartService{cartRepository: cartRepository,
		userService: userService,
		menuService: menuService,
	}
}

// AddItem after validating user input (userId/itemId) adds item to cart using cart repository
func (c *cartService) AddItem(userID, itemID string, quantity int64) error {
	isVerified, err := c.isVerifiedUserItem(userID, itemID)
	if err != nil {
		return err
	}
	if isVerified {
		var cart entity.Cart

		cart.UserID = userID
		cart.ItemID = itemID
		cart.Quantity = quantity

		item, err := c.menuService.GetItem(itemID)
		if err != nil {
			return err
		}
		cart.Item = *item

		cart.PricePerUnit = parsePriceStringToFloat(item.Price)
		cart.TotalPrice = calculateTotalPrice(cart.PricePerUnit, cart.Quantity)
		cart.AddedAt = time.Now()

		err = c.cartRepository.AddItem(cart)
		if err != nil {
			return err
		}
	}
	return nil
}

// isVerifiedUserItem verifies if user and dish exist in the respective tables
func (c *cartService) isVerifiedUserItem(userID, itemID string) (bool, error) {
	user, err := c.userService.GetUserById(userID)
	if err != nil {
		return false, err
	}
	item, err := c.menuService.GetItem(itemID)
	if err != nil {
		return false, err
	}
	if user.ID == userID && itemID == item.ID {
		return true, nil
	}
	return false, errors.New("not found")
}

// GetItemsList retrieves items list from the user's cart
func (c *cartService) GetItemsList(userID string) ([]entity.Cart, error) {
	items, err := c.cartRepository.GetCartItems(userID)
	if err != nil {
		return nil, err
	}
	return items, nil
}

// RemoveItem removes entire item from user's cart
func (c *cartService) RemoveItem(userID, itemID string) error {
	err := c.cartRepository.RemoveItem(userID, itemID)
	if err != nil {
		return err
	}
	return nil
}

// UpdateUserItem updates quantity of item in user's cart. if item doesn't exist, it saves it.
func (c *cartService) UpdateUserItem(userID, itemID string, quantity int64) error {
	item, _ := c.cartRepository.GetCartItem(userID, itemID)
	if item == nil {
		err := c.AddItem(userID, itemID, quantity)
		if err != nil {
			return err
		}
	}
	item.Quantity += quantity
	err := c.cartRepository.UpdateCartItems(userID, itemID, item.Quantity)
	if err != nil {
		return err
	}
	return nil
}

// parsePriceStringToFloat parses price string (removing USD) to float
func parsePriceStringToFloat(price string) float64 {
	s := strings.TrimSuffix(price, " USD")
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

// calculateTotalPrice calculates total price of the order
func calculateTotalPrice(price float64, quantity int64) float64 {
	return price * float64(quantity)
}
