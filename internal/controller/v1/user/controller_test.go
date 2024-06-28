package user

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	ce "github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/custom_errors"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/service/mocks"
)

var (
	user = entity.User{
		Email:    "user@example.com",
		Password: "password1234!",
	}

	userJSON = `{
    "email": "user@example.com",
    "password": "password1234!"
}`
)

func TestController_RegistrationUser(t *testing.T) {
	// Arrange
	t.Run("on registration, return created", func(t *testing.T) {
		r := chi.NewRouter()

		mockService := new(mocks.UserService)
		controller := userController{
			userService: mockService,
		}
		r.Post("/api/v1/auth/signup", controller.RegistrationUser)

		// Act
		mockService.On("CreateUser", user).Return(&user, nil)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/signup", strings.NewReader(userJSON))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		var result entity.User
		_ = json.Unmarshal(w.Body.Bytes(), &result)

		// Assert
		assert.Equal(t, w.Code, http.StatusCreated)
		assert.Equal(t, "{\"data\":\"User registered successfully\"}\n", w.Body.String())
	})

	t.Run("on registration, return error", func(t *testing.T) {
		// Arrange
		r := chi.NewRouter()
		mockService := new(mocks.UserService)
		controller := userController{
			userService: mockService,
		}
		r.Post("/api/v1/auth/signup", controller.RegistrationUser)

		// Act
		mockService.On("CreateUser", user).Return(nil, ce.ErrShortPassword)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/signup", strings.NewReader(userJSON))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		var result entity.User
		_ = json.Unmarshal(w.Body.Bytes(), &result)

		// Assert
		assert.Equal(t, w.Code, http.StatusBadRequest)
		assert.Equal(t, "{\"error\":\"error creating new user\",\"data\":\"\"}\n", w.Body.String())
	})

	t.Run("on registration, return error", func(t *testing.T) {
		// Arrange
		r := chi.NewRouter()
		mockService := new(mocks.UserService)
		controller := userController{userService: mockService}
		r.Post("/api/v1/auth/signup", controller.RegistrationUser)

		// Act
		mockService.On("CreateUser", user).Return(nil, errors.New("invalid request body"))
		req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/signup", strings.NewReader(userJSON))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		var result entity.User
		_ = json.Unmarshal(w.Body.Bytes(), &result)

		//Assert
		assert.Equal(t, w.Code, http.StatusBadRequest)
		assert.Equal(t, "{\"error\":\"error creating new user\",\"data\":\"\"}\n", w.Body.String())

	})
}

func TestController_LoginUser(t *testing.T) {
	t.Run("on login, return OK", func(t *testing.T) {
		// Arrange
		r := chi.NewRouter()

		mockService := new(mocks.UserService)
		controller := userController{
			userService: mockService,
		}
		r.Post("/api/v1/auth/login", controller.LoginUser)

		// Act
		mockService.On("GenerateTokenResponse", user.Email, user.Password).Return("string", nil)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", strings.NewReader(userJSON))
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
		controller := userController{
			userService: mockService,
		}

		r.Post("/api/v1/auth/login", controller.LoginUser)
		// Act
		mockService.On("GenerateTokenResponse", user.Email, user.Password).Return("", errors.New("incorrect password"))
		r.Post("/api/v1/auth/login", controller.LoginUser)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", strings.NewReader(userJSON))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		var result entity.User
		_ = json.Unmarshal(w.Body.Bytes(), &result)

		assert.Equal(t, w.Code, http.StatusBadRequest)
		assert.Equal(t, "{\"error\":\"error with formatting\",\"data\":\"\"}\n{\"data\":\"\"}\n", w.Body.String())
	})

}
