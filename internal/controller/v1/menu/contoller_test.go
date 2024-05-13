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

// mocking data : Mockery
// test compare responses

//go:generate mockery

// we need to set up a new server
// and to compare response
// mocking the client - not yet
// when this client is called with these methods and these arguments then return this response
// http (server and route, and api client

//func TestGetMenuByRestaurant(t *testing.T) {
//	r := chi.NewRouter()
//
//	mockService := new(mocks.MenuService)
//	controller := Controller{
//		service: mockService,
//	}
//
//	r.Get("/api/v1/restaurants/{restaurant_id}/items", controller.GetMenuByRestaurant)
//
//	t.Run("should return 200 OK response", func(t *testing.T) {
//		menus := []*entity.Menu{}
//		mockService.On("GetMenu", mock.AnythingOfType("int64")).Return(menus, nil)
//		req := httptest.NewRequest(http.MethodGet, "/api/v1/restaurants/1/items", nil)
//		rec := httptest.NewRecorder()
//		r.ServeHTTP(rec, req)
//		assert.Equal(t, http.StatusOK, rec.Code)
//
//	})
//
//	t.Run("should return menu of the restaurant", func(t *testing.T) {
//		menus := []*entity.Menu{
//			{
//				RestaurantID: 1,
//				Category:     "Extra Large Pizza",
//				Name:         "Extra Large Supreme Slice",
//				Description:  "Slice.",
//				Price:        "3.99 USD",
//			},
//		}
//		mockService.On("GetMenu", mock.AnythingOfType("int64")).Return(menus, nil)
//
//		req := httptest.NewRequest(http.MethodGet, "/api/v1/restaurants/1/items", nil)
//		rec := httptest.NewRecorder()
//
//		r.ServeHTTP(rec, req)
//
//		var result []entity.Menu
//		json.Unmarshal(rec.Body.Bytes(), &result)
//
//		assert.Equal(t, 1, len(result))
//		assert.Equal(t, "1", result[0].RestaurantID)
//		assert.Equal(t, "Extra Large Pizza", result[0].Category)
//		assert.Equal(t, "Slice.", result[0].Description)
//	})
//}

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

		mockService.On("GetMenu", mock.AnythingOfType("int64")).Return(&menus, nil)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/restaurants/1/items", nil)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)
		respBody := rec.Body.Bytes()

		var result []entity.Menu
		err := json.Unmarshal(respBody, &result)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(result))
		if len(result) > 0 {
			assert.Equal(t, int64(1), result[0].RestaurantID)
			assert.Equal(t, "Extra Large Pizza", result[0].Category)
			assert.Equal(t, "Slice.", result[0].Description)
		}
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
