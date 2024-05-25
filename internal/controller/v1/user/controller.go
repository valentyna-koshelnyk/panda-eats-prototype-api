package user

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/auth"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/service"
	ce "github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/errors"
)

// Controller handles user-related requests
type Controller struct {
	s service.UserService
}

// NewUserController creates a new UserController
func NewUserController(service service.UserService) Controller {
	return Controller{s: service}
}

// RegistrationUser handles user registration by validating and creating a new user
func (c *Controller) RegistrationUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var user *entity.User
		data, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Errorf("error reading body: %s", err)
			ce.RespondWithError(w, r, "invalid request body")
			return
		}
		err = json.Unmarshal(data, &user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Errorf("error unmarshalling body: %s", err)
			ce.RespondWithError(w, r, "invalid request body")
			return
		}
		validate := validator.New()
		err = validate.Struct(user)
		if err != nil {
			var errs validator.ValidationErrors
			errors.As(err, &errs)
			http.Error(w, errs.Error(), http.StatusBadRequest)
			return
		}
		hashedPassword, err := auth.Hash(user.Password)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Errorf("error hashing password: %s", err)
			ce.RespondWithError(w, r, "invalid request body")
			return
		}
		user.Password = hashedPassword
		_, err = c.s.CreateUser(*user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Errorf("error creating new user: %s", err)
			ce.RespondWithError(w, r, "internal server error")
			return
		}
		next.ServeHTTP(w, r)
	})
}
