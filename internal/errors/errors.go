package errors

import (
	"github.com/go-chi/render"
	"net/http"
)

// ErrorResponse has only error field which holds the message to be returned in response
type ErrorResponse struct {
	Error string `json:"error"`
}

// RespondWithError marshals error into JSON as http response
func RespondWithError(w http.ResponseWriter, r *http.Request, errMsg string) {
	render.JSON(w, r, ErrorResponse{Error: errMsg})
}
