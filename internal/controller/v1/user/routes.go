package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/config"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/repository"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/service"
)

func Routes() chi.Router {
	r := chi.NewRouter()
	db := config.GetDB()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := NewUserController(userService)

	r.Post("/register", userController.RegistrationUser)
	return r
}
