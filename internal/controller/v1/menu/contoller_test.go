package menu

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
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

func TestGetMenuByRestaurant(t *testing.T) {
	t.Run("should return menu", func(t *testing.T) {
		r := chi.NewRouter()
		server := &http.Server{
			Addr:    ":" + "3000",
			Handler: r,
		}
		r.Use(middleware.RequestID)
		log.Info("Starting server on port : 3000")
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}

		mockService := new(mocks.MenuService)
		expectedMenu := []entity.Menu{{RestaurantID: 1, Category: "Extra Large Pizza", Name: "Extra Large 5 Cheese Slice", Description: "Slice.", Price: "3.99 USD"}}
		mockService.On("GetMenuByRestaurant", mock.Anything, mock.Anything).Return(expectedMenu, nil)
		controller := NewController(mockService)
		r.Route("/api/v1/restaurants/{restaurant_id}/items", func(r chi.Router) {
			r.Get("/", controller.GetMenuByRestaurant)
		})

		req, _ := http.NewRequest(http.MethodGet, "/api/v1/restaurants/1/items", nil)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		assert.NotNil(t, rec.Code)
		assert.Equal(t, http.StatusOK, rec.Code)
		expectedJSON, _ := json.Marshal(expectedMenu)
		assert.JSONEq(t, string(expectedJSON), rec.Body.String(), "Expected menu to match")

		mockService.AssertExpectations(t)
	})
}
