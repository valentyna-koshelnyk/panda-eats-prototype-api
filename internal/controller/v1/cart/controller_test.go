package cart

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/service/mocks"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	userID   = "1"
	itemID   = "5"
	quantity = 3
)

var (
	cart = []entity.Cart{
		{
			UserID: userID,
			ItemID: itemID,
			Item: entity.Menu{
				ID:           itemID,
				RestaurantID: 3,
				Name:         "chicken",
				Description:  "delicious chicken",
				Price:        "2.50 USD",
			},
			PricePerUnit: 2.5,
			TotalPrice:   7.5,
		},
	}
	emptyCart = []entity.Cart{}
)

func TestCartController_AddItem(t *testing.T) {
	// Arrange
	t.Run("on add item, return created", func(t *testing.T) {
		// Arrange
		r := chi.NewRouter()
		mockService := new(mocks.CartService)
		controller := Controller{
			cartService: mockService,
		}

		r.Post("/api/v1/cart/item/{item_id}", controller.AddItem)

		// Act
		mockService.On("AddItem", userID, itemID, int64(quantity)).Return(nil)
		reqBody, _ := json.Marshal(utils.QuantityItemRequest{Quantity: int64(quantity)})
		req := httptest.NewRequest(http.MethodPost, "/api/v1/cart/item/5", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		// Assert
		assert.Equal(t, http.StatusCreated, rec.Code)
	})

	t.Run("on add item, return error", func(t *testing.T) {
		// Arrange
		r := chi.NewRouter()
		mockService := new(mocks.CartService)
		controller := Controller{
			cartService: mockService,
		}

		r.Post("/api/v1/cart/item/", controller.AddItem)

		// Act

		mockService.On("AddItem", userID, itemID, int64(quantity)).Return(errors.New("not found"))
		reqBody, _ := json.Marshal(utils.QuantityItemRequest{Quantity: int64(quantity)})
		req := httptest.NewRequest(http.MethodPost, "/api/v1/cart/item/5", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		// Assert
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}

func TestController_GetCartItems(t *testing.T) {
	t.Run("on get items, return OK", func(t *testing.T) {
		// Arrange
		r := chi.NewRouter()
		mockService := new(mocks.CartService)
		controller := Controller{
			cartService: mockService,
		}
		r.Get("/api/v1/cart/", controller.GetCartItems)

		// Act
		mockService.On("GetItemsList", userID).Return(cart, nil)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/cart/", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		// Assert
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("on get items, return no content", func(t *testing.T) {
		// Arrange
		r := chi.NewRouter()
		mockService := new(mocks.CartService)
		controller := Controller{
			cartService: mockService,
		}
		r.Get("/api/v1/cart/", controller.GetCartItems)
		// Act
		mockService.On("GetItemsList", userID).Return(emptyCart, nil)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/cart/", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		//Assert
		assert.Equal(t, http.StatusNoContent, rec.Code)
	})
}

func TestController_UpdateItem(t *testing.T) {
	t.Run("on update item, return updated", func(t *testing.T) {
		// Arrange
		r := chi.NewRouter()
		mockService := new(mocks.CartService)
		controller := Controller{
			cartService: mockService,
		}

		r.Patch("/api/v1/cart/item/{item_id}", controller.UpdateItem)

		// Act
		mockService.On("UpdateUserItem", userID, itemID, int64(quantity)).Return(nil)
		reqBody, _ := json.Marshal(utils.QuantityItemRequest{Quantity: int64(quantity)})
		req := httptest.NewRequest(http.MethodPatch, "/api/v1/cart/item/5", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		// Assert
		assert.Equal(t, http.StatusCreated, rec.Code)
	})

	t.Run("on update item, return error", func(t *testing.T) {
		// Arrange
		r := chi.NewRouter()
		mockService := new(mocks.CartService)
		controller := Controller{
			cartService: mockService,
		}

		r.Post("/api/v1/cart/item/{item_id}", controller.UpdateItem)

		// Act
		mockService.On("UpdateUserItem", userID, itemID, int64(quantity)).Return(errors.New("not allowed"))
		req := httptest.NewRequest(http.MethodPatch, "/api/v1/cart/item/5", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		// Assert
		assert.Equal(t, http.StatusMethodNotAllowed, rec.Code)
	})
}

func TestController_RemoveItem(t *testing.T) {
	t.Run("on remove item, return OK", func(t *testing.T) {
		// Arrange
		r := chi.NewRouter()
		mockService := new(mocks.CartService)
		controller := Controller{
			cartService: mockService,
		}

		r.Delete("/api/v1/cart/item/{item_id}", controller.RemoveItem)

		// Act
		mockService.On("RemoveItem", userID, itemID).Return(nil)
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/cart/item/5", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		// Assert
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("on remove item, return error", func(t *testing.T) {
		// Arrange
		r := chi.NewRouter()
		mockService := new(mocks.CartService)
		controller := Controller{
			cartService: mockService,
		}

		r.Delete("/api/v1/cart/item/{item_id}", controller.RemoveItem)

		// Act
		mockService.On("RemoveItem", userID, itemID).Return(errors.New("error"))
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/cart/item/5", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

}
