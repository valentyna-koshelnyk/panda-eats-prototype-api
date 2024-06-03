package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/auth"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/config"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/repository"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/service"
)

func Routes() chi.Router {
	r := chi.NewRouter()
	db := config.GetDB()

	userRepository := repository.NewUserRepository(db)
	authService := auth.NewAuthService()
	tokenService := auth.NewTokenService()
	userService := service.NewUserService(userRepository, authService, tokenService)
	userController := NewUserController(userService)

	r.Post("/signup", userController.RegistrationUser)
	r.Post("/login", userController.LoginUser)
	return r
}
