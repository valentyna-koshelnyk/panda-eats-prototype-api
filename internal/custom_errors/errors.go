package custom_errors

import (
	"github.com/go-chi/render"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func RespondWithError(w http.ResponseWriter, r *http.Request, errMsg string) {
	render.JSON(w, r, ErrorResponse{Error: errMsg})
}
