package user

import (
	"github.com/go-chi/chi/v5"
)

// Routes for user registration and login
func Routes(c UserController) chi.Router {
	r := chi.NewRouter()
	r.Post("/signup", c.RegistrationUser)
	r.Post("/login", c.LoginUser)
	return r
}
