package restaurant

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/service/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestController(t *testing.T) {
	// Arrange
	restaurants := []entity.Restaurant{

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

	restaurant := entity.Restaurant{
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
		Lng:         "-86.8565464"}

	t.Run("should return restaurants filtered by zip, price range, category", func(t *testing.T) {
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
		var result []entity.Restaurant
		_ = json.Unmarshal(respBody, &result)

		// Assert
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Pizza", result[0].Category)
	})
	t.Run("should return empty string since restaurant not found", func(t *testing.T) {
		r := chi.NewRouter()
		mockService := new(mocks.RestaurantService)
		controller := Controller{
			s: mockService,
		}
		r.Get("/api/v1/restaurants", controller.GetAll)

		// Act
		mockService.On("FilterRestaurants", "pizza", "23204", "$").Return(nil, nil)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/restaurants?category=pizza&price_range=$&zip_code=23204", nil)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		respBody := rec.Body.Bytes()

		// Assert
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.Equal(t, "", string(respBody))
	})
	t.Run("should return restaurant created message ", func(t *testing.T) {
		// Arrange
		r := chi.NewRouter()
		mockService := new(mocks.RestaurantService)
		controller := Controller{
			s: mockService,
		}
		r.Post("/api/v1/restaurants", controller.Create)

		//Act
		mockService.On("CreateRestaurant", mock.AnythingOfType("entity.Restaurant")).Return(nil)
		reqBody, _ := json.Marshal(restaurant)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/restaurants", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		r.ServeHTTP(response, req)

		// Assert
		assert.Equal(t, http.StatusCreated, response.Code)
		assert.JSONEq(t, `"Restaurant created"`, response.Body.String())
	})

	t.Run("should return error decoding restaurant", func(t *testing.T) {
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
				"Lat": 33.5098,
				"Lng": -86.85464
}`)
		// Act
		mockService.On("CreateRestaurant", mock.AnythingOfType("entity.Restaurant")).Return(nil)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/restaurants", bytes.NewBuffer(wrongJSON))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Errorf(t, errors.New("Error decoding restaurant"), rec.Body.String())
	})
	t.Run("should return restaurant updated successfully", func(t *testing.T) {
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

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `"Restaurant updated successfully"`, rec.Body.String())
	})

	t.Run("should return error updating restaurant", func(t *testing.T) {
		// Arrange
		r := chi.NewRouter()
		mockService := new(mocks.RestaurantService)
		controller := Controller{
			s: mockService,
		}
		r.Put("/api/v1/restaurants", controller.Update)

		// Act
		mockService.On("UpdateRestaurant", restaurant).Return(errors.New("error"))
		reqBody, _ := json.Marshal(restaurant)
		req := httptest.NewRequest(http.MethodPut, "/api/v1/restaurants", bytes.NewBuffer(reqBody))
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		// Assert
		var errorResponse map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &errorResponse)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, "error", errorResponse["error"])
	})

	t.Run("should delete restaurant successfully", func(t *testing.T) {
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
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `"Restaurant deleted successfully"`, rec.Body.String())
	})

	t.Run("should return error deleting restaurant", func(t *testing.T) {
		r := chi.NewRouter()
		mockService := new(mocks.RestaurantService)
		controller := Controller{
			s: mockService,
		}
		r.Delete("/api/v1/restaurants/{restaurant_id}", controller.Delete)

		// Act
		mockService.On("DeleteRestaurant", int64(13)).Return(errors.New("Error deleting restaurant"))
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/restaurants/13", nil)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		// Assert
		var errorResponse map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &errorResponse)
		assert.NoError(t, err)
		assert.Equal(t, "Error deleting restaurant", errorResponse["error"])
	})
}
