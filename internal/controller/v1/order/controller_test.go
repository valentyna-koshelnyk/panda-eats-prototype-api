package order

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/service/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	userID  = "50aa4686-bb62-4202-b2ce-471df794adea"
	orderID = "jQ8pMyTsrfs7ZkvmDj6y8"
)

var orders = []entity.Order{
	{OrderID: orderID,
		UserID: userID,
		Items: []entity.Cart{{
			UserID: userID,
			ItemID: "12",
			Item: entity.Menu{
				ID:           "12",
				RestaurantID: 3,
				Name:         "chicken",
				Description:  "delicious chicken",
				Price:        "2.50 USD",
			},
			PricePerUnit: 2.5,
			TotalPrice:   7.5,
		},
		},
		TotalOrderPrice: 124.50,
		Status:          entity.InProcess,
	},
}

func TestOrderController_CreateOrder(t *testing.T) {
	t.Run("on create, return created", func(t *testing.T) {
		// Arrange
		r := chi.NewRouter()
		mockOrderService := new(mocks.OrderService)

		controller := orderController{
			orderService: mockOrderService,
		}
		r.Post("/api/v1/cart/order", controller.CreateOrder)

		// Act
		mockOrderService.On("CreateOrder", userID).Return(nil)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/cart/order", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		// Assert
		assert.Equal(t, http.StatusCreated, rec.Code)
	})

	t.Run("on create, return error", func(t *testing.T) {
		// Arrange
		r := chi.NewRouter()
		mockOrderService := new(mocks.OrderService)

		controller := orderController{
			orderService: mockOrderService,
		}
		r.Post("/api/v1/cart/order", controller.CreateOrder)

		// Act
		mockOrderService.On("CreateOrder", userID).Return(errors.New("user doesn't exist"))
		req := httptest.NewRequest(http.MethodPost, "/api/v1/cart/order", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func TestOrderController_UpdateOrderStatusShipped(t *testing.T) {
	t.Run("on update, return accepted", func(t *testing.T) {
		// Arrange
		r := chi.NewRouter()
		mockOrderService := new(mocks.OrderService)
		controller := orderController{
			orderService: mockOrderService,
		}
		r.Patch("/api/v1/cart/order/shipped", controller.UpdateOrderStatusShipped)
		// Act
		mockOrderService.On("UpdateOrderStatusShipped", userID, orderID).Return(nil)
		reqBody, _ := json.Marshal(entity.OrderIDRequest{orderID})
		req := httptest.NewRequest(http.MethodPatch, "/api/v1/cart/order/shipped", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)
		// Assert
		assert.Equal(t, http.StatusAccepted, rec.Code)
	})
	t.Run("on update, return bad request", func(t *testing.T) {
		// Arrange
		r := chi.NewRouter()
		mockOrderService := new(mocks.OrderService)
		controller := orderController{
			orderService: mockOrderService,
		}
		r.Patch("/api/v1/cart/order/shipped", controller.UpdateOrderStatusShipped)
		// Act
		mockOrderService.On("UpdateOrderStatusShipped", userID, orderID).Return(errors.New("order not found"))
		reqBody, _ := json.Marshal(entity.OrderIDRequest{orderID})
		req := httptest.NewRequest(http.MethodPatch, "/api/v1/cart/order/shipped", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)
		// Assert
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func TestOrderController_UpdateOrderStatusDelivered(t *testing.T) {
	t.Run("on update, return accepted", func(t *testing.T) {
		// Arrange
		r := chi.NewRouter()
		mockOrderService := new(mocks.OrderService)
		controller := orderController{
			orderService: mockOrderService,
		}
		r.Patch("/api/v1/cart/order", controller.UpdateOrderStatusDelivered)
		// Act
		mockOrderService.On("UpdateOrderStatusDelivered", userID, orderID).Return(nil)
		reqBody, _ := json.Marshal(entity.OrderIDRequest{orderID})
		req := httptest.NewRequest(http.MethodPatch, "/api/v1/cart/order", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)
		// Assert
		assert.Equal(t, http.StatusAccepted, rec.Code)
	})
	t.Run("on update, return bad request", func(t *testing.T) {
		// Arrange
		r := chi.NewRouter()
		mockOrderService := new(mocks.OrderService)
		controller := orderController{
			orderService: mockOrderService,
		}
		r.Patch("/api/v1/cart/order/delivery", controller.UpdateOrderStatusDelivered)
		// Act
		mockOrderService.On("UpdateOrderStatusDelivered", userID, orderID).Return(errors.New("order not found"))
		reqBody, _ := json.Marshal(entity.OrderIDRequest{OrderID: orderID})
		req := httptest.NewRequest(http.MethodPatch, "/api/v1/cart/order/delivery", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)
		// Assert
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func TestOrderController_GetOrdersHistory(t *testing.T) {
	t.Run("on get, return OK", func(t *testing.T) {
		// Arrange
		r := chi.NewRouter()
		mockOrderService := new(mocks.OrderService)
		controller := orderController{
			orderService: mockOrderService,
		}
		r.Get("/api/v1/cart/order/orders", controller.GetOrdersHistory)
		// Act
		mockOrderService.On("GetOrderHistory", userID).Return(orders, nil)
		reqBody, _ := json.Marshal(entity.OrderIDRequest{OrderID: orderID})
		req := httptest.NewRequest(http.MethodGet, "/api/v1/cart/order/orders", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		// Assert
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("on get, return no content ", func(t *testing.T) {
		// Arrange
		r := chi.NewRouter()
		mockOrderService := new(mocks.OrderService)
		controller := orderController{
			orderService: mockOrderService,
		}
		r.Get("/api/v1/cart/order/orders", controller.GetOrdersHistory)
		// Act
		mockOrderService.On("GetOrderHistory", userID).Return([]entity.Order{}, nil)
		reqBody, _ := json.Marshal(entity.OrderIDRequest{OrderID: orderID})
		req := httptest.NewRequest(http.MethodGet, "/api/v1/cart/order/orders", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		// Assert
		assert.Equal(t, http.StatusNoContent, rec.Code)
	})
}
