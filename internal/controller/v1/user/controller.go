package user

import (
	"encoding/json"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	ce "github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/custom-errors"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/service"
	"io"
	"net/http"
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
func (c *Controller) RegistrationUser(w http.ResponseWriter, r *http.Request) {
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

	_, err = c.s.CreateUser(*user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("error creating new user: %s", err)
		ce.RespondWithError(w, r, err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)

	render.JSON(w, r, "User registered successfully")
	return
}

func (c *Controller) LoginUser(w http.ResponseWriter, r *http.Request) {
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

	response, err := c.s.GenerateTokenResponse(*user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("error generating response: %s", err)
		ce.RespondWithError(w, r, err.Error())
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, response)
	return
}
