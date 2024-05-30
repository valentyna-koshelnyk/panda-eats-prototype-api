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
	menu := []entity.Menu{
		{RestaurantID: 1, Category: "Extra Large Pizza", Name: "Extra Large Supreme Slice", Description: "Slice.", Price: "3.99 USD"},
	}

	t.Run("on success, return OK", func(t *testing.T) {
		r := chi.NewRouter()

		mockService := new(mocks.MenuService)
		controller := Controller{
			s: mockService,
		}

		var items []entity.Item
		for _, m := range menu {
			items = append(items, m)
		}

		response := entity.NewResponse(items)

		r.Get("/api/v1/restaurants/{restaurant_id}/items", controller.GetMenuByRestaurant)

		mockService.On("GetMenu", int64(1)).Return(response, nil)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/restaurants/1/items", nil)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		var result entity.Response
		err := json.Unmarshal(rec.Body.Bytes(), &result)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("should return empty string if menu not found", func(t *testing.T) {
		r := chi.NewRouter()

		mockService := new(mocks.MenuService)
		controller := Controller{s: mockService}
		r.Get("/api/v1/restaurants/{restaurant_id}/items", controller.GetMenuByRestaurant)

		mockService.On("GetMenu", int64(2)).Return(nil, nil)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/restaurants/2/items", nil)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		respBody := rec.Body.Bytes()
		assert.Equal(t, "{\"error\":\"restaurant data not found\"}\n", string(respBody))
		assert.Equal(t, http.StatusNotFound, rec.Code)

	})
}
