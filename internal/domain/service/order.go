package service

import (
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
	CreateOrder(userEmail string) error
	calculateTotalOrderPrice(carts []entity.Cart) float64
}

func NewOrderService(orderRepository repository.OrderRepository, cartService CartService, userService UserService) OrderService {
	return &orderService{
		orderRepository: orderRepository,
		cartService:     cartService,
		userService:     userService,
	}
}

// TODO: to remove items from cart once they are transferred to order
func (s *orderService) CreateOrder(userEmail string) error {
	user, err := s.userService.GetUserByEmail(userEmail)
	cartItems, err := s.cartService.GetItemsList(user.UserID)
	if err != nil {
		return err
	}
	var order entity.Order

	order.OrderID = nanoid.New()
	order.Items = cartItems
	order.TotalOrderPrice = s.calculateTotalOrderPrice(cartItems)
	order.UserID = user.UserID
	order.AddedAt = time.Now()
	order.Status = 1

	err = s.orderRepository.CreateOrder(order)
	err = s.cartService.RemoveItems(cartItems)
	if err != nil {
		return err
	}
	return nil
}

func (s *orderService) calculateTotalOrderPrice(carts []entity.Cart) float64 {
	var totalPrice float64
	for _, cart := range carts {
		totalPrice += cart.TotalPrice
	}
	return totalPrice
}
