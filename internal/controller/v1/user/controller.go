package user

import (
	"encoding/json"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/service"
)

// userController handles user-related requests
type userController struct {
	userService service.UserService
}

// UserController interface for user registration and login
type UserController interface {
	RegistrationUser(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)
}

// NewUserController creates a new UserController
func NewUserController(userService service.UserService) UserController {
	return &userController{userService: userService}
}

// RegistrationUser handler for a user registration
//
//		@Summary Registration of a new user
//		@Description Creates a new account in the system
//		@Accept json
//	    @Produce  json
//		@Success      201
//	    @Failure 404 {object}  entity.CustomResponse
//		@Param user body entity.User true "User Registration Information"
//		@Router /auth/signup [post]
func (c *userController) RegistrationUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user entity.User

	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("error reading body: %s", err)
		entity.RespondWithJSON(w, r, "", "error creating new user")
		return
	}

	err = json.Unmarshal(data, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("error unmarshalling body: %s", err)
		entity.RespondWithJSON(w, r, "", "error creating new user")
		return
	}

	_, err = c.userService.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("error creating new user: %s", err)
		entity.RespondWithJSON(w, r, "", "error creating new user")
		return
	}
	w.WriteHeader(http.StatusCreated)

	entity.RespondWithJSON(w, r, "User registered successfully", "")
	return
}

// LoginUser handler for login user
//
//		@Summary Login system
//		@Description User authentication handler
//		@Accept json
//	    @Produce  json
//		@Param user body entity.User true "User Login Information"
//		@Success      200  {object}  entity.User
//	    @Failure 404 {object}  entity.CustomResponse
//		@Router /auth/login [post]
func (c *userController) LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user *entity.User

	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("error reading body: %s", err)
		entity.RespondWithJSON(w, r, "", "user not found")
		return
	}

	err = json.Unmarshal(data, &user)

	response, err := c.userService.GenerateTokenResponse(user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("error generating response: %s", err)
		entity.RespondWithJSON(w, r, "", "error with formatting")
	}

	w.WriteHeader(http.StatusOK)
	entity.RespondWithTokenJSON(w, r, response, "")
	return
}
