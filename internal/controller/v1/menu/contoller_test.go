package menu

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"

	custom_errors "github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/custom-errors"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/service/mocks"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/utils"
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

		pagination := &utils.Pagination{
			Limit:      1,
			Page:       1,
			TotalRows:  1,
			TotalPages: 1,
			Rows:       menu,
		}

		r.Get("/api/v1/restaurants/{restaurant_id}/items", controller.GetMenuByRestaurant)

		mockService.On("GetMenu", 1, 10, 0).Return(pagination, nil)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/restaurants/1/items", nil)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		var result utils.PaginatedResponse
		err := json.Unmarshal(rec.Body.Bytes(), &result)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("should return empty string if menu not found", func(t *testing.T) {
		r := chi.NewRouter()

		mockService := new(mocks.MenuService)
		controller := Controller{s: mockService}
		r.Get("/api/v1/restaurants/{restaurant_id}/items", controller.GetMenuByRestaurant)

		mockService.On("GetMenu", 2, 10, 0).Return(nil, custom_errors.ErrNotFound)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/restaurants/2/items", nil)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		respBody := rec.Body.Bytes()
		assert.JSONEq(t, "{\"error\":\"no items available\",\"data\":\"\"}", string(respBody))
		assert.Equal(t, http.StatusNotFound, rec.Code)

	})
}
