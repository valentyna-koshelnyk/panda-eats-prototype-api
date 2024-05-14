package menu

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/service/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMenuByRestaurant(t *testing.T) {

	t.Run("should correctly retrieve and return menu details", func(t *testing.T) {
		r := chi.NewRouter()

		mockService := new(mocks.MenuService)
		controller := Controller{
			service: mockService,
		}

		r.Get("/api/v1/restaurants/{restaurant_id}/items", controller.GetMenuByRestaurant)
		menus := []entity.Menu{
			{RestaurantID: 1, Category: "Extra Large Pizza", Name: "Extra Large Supreme Slice", Description: "Slice.", Price: "3.99 USD"},
		}

		mockService.On("GetMenu", int64(1)).Return(&menus, nil)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/restaurants/1/items", nil)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)
		respBody := rec.Body.Bytes()

		var result []entity.Menu
		err := json.Unmarshal(respBody, &result)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(result))
		assert.Equal(t, int64(1), result[0].RestaurantID)
		assert.Equal(t, "Extra Large Pizza", result[0].Category)
		assert.Equal(t, "Slice.", result[0].Description)
	})

	t.Run("should return empty string if menu not found", func(t *testing.T) {
		r := chi.NewRouter()

		mockService := new(mocks.MenuService)
		controller := Controller{service: mockService}
		r.Get("/api/v1/restaurants/{restaurant_id}/items", controller.GetMenuByRestaurant)

		mockService.On("GetMenu", int64(2)).Return(nil, nil)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/restaurants/2/items", nil)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		respBody := rec.Body.Bytes()
		assert.Equal(t, "", string(respBody))
		assert.Equal(t, http.StatusNotFound, rec.Code)

	})
}
