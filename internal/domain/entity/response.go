package entity

import (
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/utils"
	"net/http"

	"github.com/go-chi/render"
)

type CustomResponse struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}
type TokenData struct {
	Token string `json:"token"`
}

type PaginatedResponse struct {
	APIVersion string        `json:"apiVersion"`
	Data       PaginatedData `json:"data"`
}

type PaginatedData struct {
	Limit      int   `json:"limit,omitempty"`
	Offset     int   `json:"offset,omitempty"`
	TotalRows  int64 `json:"total_rows"`
	TotalPages int   `json:"total_pages"`
	Items      any   `json:"items"`
}

const (
	DefaultAPIVersion = "1.0"
	DefaultStartIndex = 1
)

func NewPaginatedResponse(pagination *utils.Pagination) *PaginatedResponse {
	return &PaginatedResponse{
		APIVersion: DefaultAPIVersion,
		Data: PaginatedData{
			Limit:      pagination.Limit,
			Offset:     pagination.Page,
			TotalRows:  pagination.TotalRows,
			TotalPages: pagination.TotalPages,
			Items:      pagination.Rows,
		},
	}
}

// RespondWithJSON message into JSON as http response
func RespondWithJSON(w http.ResponseWriter, r *http.Request, msg string, err string) {
	render.JSON(w, r, CustomResponse{Data: msg, Error: err})
}

func RespondWithTokenJSON(w http.ResponseWriter, r *http.Request, msg string, err string) {
	render.JSON(w, r, CustomResponse{
		Data: TokenData{
			Token: msg,
		},
		Error: err,
	})
}
