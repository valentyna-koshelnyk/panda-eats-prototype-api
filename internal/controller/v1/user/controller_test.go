package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	ce "github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/custom-errors"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/service/mocks"
)

var (
	user = entity.User{
		Email:    "user@example.com",
		Password: "password1234!",
	}

	emptyUser = entity.User{}

	wrongPassword = entity.User{
		Email:    "user@example.com",
		Password: "pass",
	}
)

func TestController_RegistrationUser(t *testing.T) {
	// Arrange
	t.Run("on registration, return created", func(t *testing.T) {
		r := chi.NewRouter()

		mockService := new(mocks.UserService)
		controller := Controller{
			s: mockService,
		}
		r.Post("/api/v1/auth/signup", controller.RegistrationUser)

		// Act
		mockService.On("CreateUser", user).Return(user, nil)
		reqBody, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/signup", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		var result entity.User
		_ = json.Unmarshal(w.Body.Bytes(), &result)

		// Assert
		assert.Equal(t, w.Code, http.StatusCreated)
		assert.Equal(t, "{\"message\":\"User registered successfully\"}\n", w.Body.String())
	})

	t.Run("on registration, return error", func(t *testing.T) {
		// Arrange
		r := chi.NewRouter()
		mockService := new(mocks.UserService)
		controller := Controller{
			s: mockService,
		}
		r.Post("/api/v1/auth/signup", controller.RegistrationUser)

		// Act
		mockService.On("CreateUser", wrongPassword).Return(entity.User{}, ce.ErrShortPassword)
		reqBody, _ := json.Marshal(wrongPassword)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/signup", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		var result entity.User
		_ = json.Unmarshal(w.Body.Bytes(), &result)

		// Assert
		assert.Equal(t, w.Code, http.StatusBadRequest)
		assert.Equal(t, "{\"error\":\"error creating new user\"}\n", w.Body.String())
	})

	t.Run("on registration, return error", func(t *testing.T) {
		// Arrange
		r := chi.NewRouter()
		mockService := new(mocks.UserService)
		controller := Controller{s: mockService}
		r.Post("/api/v1/auth/signup", controller.RegistrationUser)

		// Act
		mockService.On("CreateUser", emptyUser).Return(entity.User{}, errors.New("invalid request body"))
		reqBody, _ := json.Marshal(emptyUser)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/signup", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		var result entity.User
		_ = json.Unmarshal(w.Body.Bytes(), &result)

		//Assert
		assert.Equal(t, w.Code, http.StatusBadRequest)
		assert.Equal(t, "{\"error\":\"error creating new user\"}\n", w.Body.String())

	})
}

func TestController_LoginUser(t *testing.T) {
	t.Run("on login, return OK", func(t *testing.T) {
		// Arrange
		r := chi.NewRouter()

		mockService := new(mocks.UserService)
		controller := Controller{
			s: mockService,
		}
		r.Post("/api/v1/auth/login", controller.LoginUser)

		// Act
		mockService.On("GenerateTokenResponse", user).Return("string", nil)
		reqBody, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		var result entity.User
		_ = json.Unmarshal(w.Body.Bytes(), &result)

		// Assess
		assert.Equal(t, w.Code, http.StatusOK)
	})

	t.Run("on login, return error", func(t *testing.T) {
		r := chi.NewRouter()
		mockService := new(mocks.UserService)
		controller := Controller{
			s: mockService,
		}

		r.Post("/api/v1/auth/login", controller.LoginUser)
		// Act
		mockService.On("GenerateTokenResponse", wrongPassword).Return("", errors.New("incorrect password"))
		reqBody, _ := json.Marshal(wrongPassword)
		r.Post("/api/v1/auth/login", controller.LoginUser)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		var result entity.User
		_ = json.Unmarshal(w.Body.Bytes(), &result)

		assert.Equal(t, w.Code, http.StatusBadRequest)
		assert.Equal(t, "{\"error\":\"error with formatting\"}\n{}\n", w.Body.String())
	})

}
