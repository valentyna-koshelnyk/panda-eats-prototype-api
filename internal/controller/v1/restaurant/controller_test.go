package restaurant

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"

	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/service/mocks"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/utils"
)

var (
	restaurant = entity.Restaurant{
		ID:          2,
		Position:    5,
		Name:        "Philly Fresh Cheesesteaks (541-B Graymont Ave)",
		Score:       0,
		Ratings:     0,
		Category:    "Pizza",
		PriceRange:  "$",
		FullAddress: "541-B Graymont Ave, Birmingham, AL, 23204",
		ZipCode:     "23204",
		Lat:         "33.43098",
		Lng:         "-86.8565464",
	}

	restaurants = []entity.Restaurant{
		{ID: 2,
			Position:    5,
			Name:        "Philly Fresh Cheesesteaks (541-B Graymont Ave)",
			Score:       0,
			Ratings:     0,
			Category:    "Pizza",
			PriceRange:  "$",
			FullAddress: "541-B Graymont Ave, Birmingham, AL, 23204",
			ZipCode:     "23204",
			Lat:         "33.43098",
			Lng:         "-86.8565464"},
	}
)

func TestController_GetAll(t *testing.T) {
	// Arrange
	t.Run("on get all, return OK", func(t *testing.T) {
		r := chi.NewRouter()

		mockService := new(mocks.RestaurantService)
		controller := Controller{
			s: mockService,
		}

		r.Get("/api/v1/restaurants", controller.GetAll)

		// Act
		mockService.On("FilterRestaurants", "pizza", "23204", "$").Return(restaurants, nil)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/restaurants?category=pizza&price_range=$&zip_code=23204", nil)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)
		respBody := rec.Body.Bytes()
		var result utils.PaginatedResponse
		_ = json.Unmarshal(respBody, &result)

		// Assert
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("on get all, return error", func(t *testing.T) {
		r := chi.NewRouter()
		mockService := new(mocks.RestaurantService)
		controller := Controller{
			s: mockService,
		}
		r.Get("/api/v1/restaurants", controller.GetAll)

		// Act
		mockService.On("FilterRestaurants", "p", "1", "$").Return(nil, errors.New("restaurant doesn't exist"))
		req := httptest.NewRequest(http.MethodGet, "/api/v1/restaurants?category=p&price_range=$&zip_code=1", nil)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.JSONEq(t, "{\"data\":\"\", \"error\":\"error getting restaurants\"}", rec.Body.String())
	})
}

func TestController_Create(t *testing.T) {
	t.Run("on create, return Created", func(t *testing.T) {
		// Arrange
		r := chi.NewRouter()
		mockService := new(mocks.RestaurantService)
		controller := Controller{
			s: mockService,
		}
		r.Post("/api/v1/restaurants", controller.Create)

		//Act
		mockService.On("CreateRestaurant", entity.Restaurant{}).Return(nil)
		reqBody, _ := json.Marshal(entity.Restaurant{})
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/restaurants", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		r.ServeHTTP(response, req)

		// Assert
		assert.Equal(t, http.StatusCreated, response.Code)
	})

	t.Run("on create, return error", func(t *testing.T) {
		r := chi.NewRouter()
		mockService := new(mocks.RestaurantService)
		controller := Controller{
			s: mockService,
		}
		r.Post("/api/v1/restaurants", controller.Create)
		wrongJSON := []byte(`{
			"ID": 1,
			"Position": 6,
			"Name": "Restaurant",
			"Score": 0.0,
			"Ratings": 0,
			"Category": "American, Cheesesteak, Sandwiches, Alcohol",
			"PriceRange": "$",
			"FullAddress": "541-B Graymont Ave, Birmingham, AL, 35204",
			"ZipCode": "35204",
			"Lat": "33.5098"",
			"Lng": "-86.85464"
		}`)
		mockService.On("CreateRestaurant", restaurant).Return(nil)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/restaurants", bytes.NewBuffer(wrongJSON))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.JSONEq(t, "{\"data\":\"\", \"error\":\"error decoding restaurant\"}", rec.Body.String())
	})
}

func TestController_Update(t *testing.T) {
	t.Run("on update, return NoContent", func(t *testing.T) {
		// Arrange
		r := chi.NewRouter()
		mockService := new(mocks.RestaurantService)
		controller := Controller{
			s: mockService,
		}
		r.Put("/api/v1/restaurants", controller.Update)

		// Act
		mockService.On("UpdateRestaurant", restaurant).Return(nil)
		reqBody, _ := json.Marshal(restaurant)
		req := httptest.NewRequest(http.MethodPut, "/api/v1/restaurants", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		// Assert
		assert.Equal(t, http.StatusNoContent, rec.Code)
	})

	t.Run("on update, return error", func(t *testing.T) {
		// Arrange
		r := chi.NewRouter()
		mockService := new(mocks.RestaurantService)
		controller := Controller{
			s: mockService,
		}
		r.Put("/api/v1/restaurants", controller.Update)

		// Act
		mockService.On("UpdateRestaurant", restaurant).Return(errors.New("error updating restaurant"))
		reqBody, _ := json.Marshal(restaurant)
		req := httptest.NewRequest(http.MethodPut, "/api/v1/restaurants", bytes.NewBuffer(reqBody))
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		// Assert
		var restaurant entity.Restaurant
		_ = json.Unmarshal(rec.Body.Bytes(), &restaurant)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.JSONEq(t, "{\"data\":\"\", \"error\":\"error updating restaurant\"}", rec.Body.String())
	})
}

func TestController_Delete(t *testing.T) {
	t.Run("on delete, return NoContent", func(t *testing.T) {
		// Arrange
		r := chi.NewRouter()
		mockService := new(mocks.RestaurantService)
		controller := Controller{
			s: mockService,
		}
		r.Delete("/api/v1/restaurants/{restaurant_id}", controller.Delete)

		// Act
		mockService.On("DeleteRestaurant", int64(1)).Return(nil)
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/restaurants/1", nil)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		// Assert
		assert.Equal(t, http.StatusNoContent, rec.Code)
	})

	t.Run("on delete, return error", func(t *testing.T) {
		// Arrange
		r := chi.NewRouter()
		mockService := new(mocks.RestaurantService)
		controller := Controller{
			s: mockService,
		}
		r.Delete("/api/v1/restaurants/{restaurant_id}", controller.Delete)

		// Act
		mockService.On("DeleteRestaurant", int64(13)).Return(errors.New("error deleting restaurant"))
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/restaurants/13", nil)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.JSONEq(t, "{\"data\":\"\", \"error\":\"error deleting restaurant\"}", rec.Body.String())
	})
}
