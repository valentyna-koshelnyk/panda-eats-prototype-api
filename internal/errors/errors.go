package errors

import (
	"errors"
	"github.com/go-chi/render"
	"net/http"
)

var (
	ErrEmptyUserField  = errors.New("username, password and email can't be empty")
	ErrFieldWithSpaces = errors.New("username and password can't have spaces")
	ErrShortPassword   = errors.New("password shorter than 8 characters")
	ErrInvalidEmail    = errors.New("invalid email address")
)

// ErrorResponse has only error field which holds the message to be returned in response
type ErrorResponse struct {
	Error string `json:"error"`
}

// RespondWithError marshals error into JSON as http response
func RespondWithError(w http.ResponseWriter, r *http.Request, errMsg string) {
	render.JSON(w, r, ErrorResponse{Error: errMsg})
}
