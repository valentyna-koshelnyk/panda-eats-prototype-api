package service

import (
	"errors"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/repository"
	"strconv"
	"strings"
)

// CartService is interface for cart service having methods for adding item and verification of user/item
type CartService interface {
	AddItem(userID, itemID string, quantity int64) error
	isVerifiedUserItem(userID, itemID string) (bool, error)
}

// cartService is struct for cart
type cartService struct {
	r           repository.CartRepository
	userService UserService
	menuService MenuService
}

// NewCartService constructor for cartservice with injected fields as repo, user service and menu service
func NewCartService(r repository.CartRepository, userService UserService, menuService MenuService) CartService {
	return &cartService{r: r,
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

		cart.PricePerUnit = ParsePriceStringToFloat(item.Price)
		cart.TotalPrice = CalculateTotalPrice(cart.PricePerUnit, cart.Quantity)

		err = c.r.AddItem(cart)
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

func ParsePriceStringToFloat(price string) float64 {
	s := strings.TrimSuffix(price, " USD")
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func CalculateTotalPrice(price float64, quantity int64) float64 {
	return price * float64(quantity)
}
