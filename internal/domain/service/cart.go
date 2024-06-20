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
//
//go:generate mockery --name=CartService
type CartService interface {
	AddItem(userID, itemID string, quantity int64) error
	GetItemsList(userID string) ([]entity.Cart, error)
	RemoveItem(userID, itemID string) error
	RemoveItems(cartItems []entity.Cart) error
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
	// TODO: to replace getUserById to bool isUserPresent, so we needn't to retrieve entire object but bool value
	user, _ := c.userService.GetUserByID(userID)
	if user == nil {
		return errors.New("user not found")
	}

	item, err := c.menuService.GetItem(itemID)
	if err != nil {
		return err
	}
	if item == nil {
		return errors.New("item not found")
	}

	existingCartItem, _ := c.cartRepository.GetCartItem(user.UserID, itemID)
	if existingCartItem != nil {
		return c.UpdateUserItem(user.UserID, itemID, quantity)
	}

	var cart entity.Cart
	cart.Item = *item
	cart.UserID = user.UserID
	cart.ItemID = itemID
	cart.Quantity = quantity
	cart.PricePerUnit = parsePriceStringToFloat(item.Price)
	cart.TotalPrice = calculateTotalPrice(cart.PricePerUnit, cart.Quantity)
	cart.AddedAt = time.Now()

	return c.cartRepository.AddItem(cart)
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
	existingItem, err := c.cartRepository.GetCartItem(userID, itemID)
	if err != nil {
		return err
	}
	if existingItem == nil {
		err = errors.New("item doesn't exist")
		return err
	}
	err = c.cartRepository.RemoveItem(userID, itemID)
	if err != nil {
		return err
	}
	return nil
}

// TODO to find out if dynnamo supports batch removal of items
func (c *cartService) RemoveItems(cartItems []entity.Cart) error {
	for _, cart := range cartItems {
		err := c.cartRepository.RemoveItem(cart.UserID, cart.ItemID)
		if err != nil {
			return err
		}
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
	item.TotalPrice = calculateTotalPrice(item.PricePerUnit, item.Quantity)

	err := c.cartRepository.UpdateCartItems(userID, *item)
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
