package menu

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/service/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMenuByRestaurant(t *testing.T) {
	r := chi.NewRouter()

	mockService := new(mocks.MenuService)
	controller := Controller{
		service: mockService,
	}

	r.Get("/api/v1/restaurants/{restaurant_id}/items", controller.GetMenuByRestaurant)

	t.Run("should return 200 OK response", func(t *testing.T) {
		menus := []entity.Menu{}
		mockService.On("GetMenu", mock.AnythingOfType("int64")).Return(&menus, nil)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/restaurants/1/items", nil)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("should return menu of the restaurant", func(t *testing.T) {
		menus := []entity.Menu{
			{
				RestaurantID: 1,
				Category:     "Extra Large Pizza",
				Name:         "Extra Large Supreme Slice",
				Description:  "Slice.",
				Price:        "3.99 USD",
			},
		}
		mockService.On("GetMenu", mock.AnythingOfType("int64")).Return(&menus, nil)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/restaurants/1/items", nil)
		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		respBody := rec.Body.Bytes()
		t.Logf("Response Body: %s", respBody)

		var result []entity.Menu
		err := json.Unmarshal(respBody, &result)
		if err != nil {
			t.Errorf("Failed to unmarshal response body: %v", err)
		}

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, 1, len(result))
		if len(result) > 0 {
			assert.Equal(t, int64(1), result[0].RestaurantID)
			assert.Equal(t, "Extra Large Pizza", result[0].Category)
			assert.Equal(t, "Slice.", result[0].Description)
		}
	})
}
