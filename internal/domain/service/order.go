package service

import (
	"errors"
	"github.com/aidarkhanov/nanoid"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/repository"
	"time"
)

type orderService struct {
	orderRepository repository.OrderRepository
	cartService     CartService
	userService     UserService
}

type OrderService interface {
	CreateOrder(userID string) error
	calculateTotalOrderPrice(carts []entity.Cart) float64
	GetOrderHistory(userID string) ([]entity.Order, error)
	UpdateOrderStatusShipped(userID, orderID string) error
	UpdateOrderStatusDelivered(userID, orderID string) error
}

// NewOrderService is a struct for order service
func NewOrderService(orderRepository repository.OrderRepository, cartService CartService, userService UserService) OrderService {
	return &orderService{
		orderRepository: orderRepository,
		cartService:     cartService,
		userService:     userService,
	}
}

// CreateOrder imlpements order placement and empties the cart once order is persisted through repository layer
func (s *orderService) CreateOrder(userID string) error {
	cartItems, err := s.cartService.GetItemsList(userID)
	if err != nil {
		return err
	}
	if len(cartItems) == 0 {
		return errors.New("cart is empty")
	}
	var order entity.Order

	order.OrderID = nanoid.New()
	order.Items = cartItems
	order.TotalOrderPrice = s.calculateTotalOrderPrice(cartItems)
	order.UserID = userID
	order.AddedAt = time.Now()
	order.Status = 1

	err = s.orderRepository.CreateOrder(order)
	err = s.cartService.RemoveItems(cartItems)
	if err != nil {
		return err
	}
	return nil
}

// calculateTotalOrderPrice calculates total price of the order based on totalPrice attribute of the cart
func (s *orderService) calculateTotalOrderPrice(carts []entity.Cart) float64 {
	var totalPrice float64
	for _, cart := range carts {
		totalPrice += cart.TotalPrice
	}
	return totalPrice
}

// UpdateOrderStatusShipped updates order status (sets respective enum) as shipped based on user id and order id
func (s *orderService) UpdateOrderStatusShipped(userID, orderID string) error {
	order, err := s.orderRepository.GetOrderInformation(userID, orderID)
	if err != nil {
		return errors.New("order not found")
	}
	order.Status = entity.Shipped
	err = s.orderRepository.UpdateOrderStatus(order)

	return nil
}

// UpdateOrderStatusDelivered updates order status (sets respective enum) as delivered based on user id and order id
func (s *orderService) UpdateOrderStatusDelivered(userID, orderID string) error {
	order, err := s.orderRepository.GetOrderInformation(userID, orderID)
	if err != nil {
		return errors.New("order not found")
	}
	order.Status = entity.Delivered
	err = s.orderRepository.UpdateOrderStatus(order)
	return nil
}

// GetOrderHistory retrieves history of orders based on userID
func (s *orderService) GetOrderHistory(userID string) ([]entity.Order, error) {
	orders, err := s.orderRepository.GetOrderHistory(userID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
