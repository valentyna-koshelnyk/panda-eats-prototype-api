package utils

import (
	"github.com/go-chi/render"
	"net/http"
)

type CustomResponse struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

func NewCustomResponse(message string, err string) *CustomResponse {
	return &CustomResponse{
		Message: message,
		Error:   err,
	}
}

type PaginatedResponse struct {
	APIVersion string `json:"apiVersion"`
	Data       Data   `json:"data"`
}

type Data struct {
	StartIndex   int    `json:"start_index"`
	ItemsCount   int    `json:"items_count"`
	ItemsPerPage int    `json:"items_per_page"`
	Items        []Item `json:"items"`
}

type Item interface{}

const (
	DefaultAPIVersion = "1.0"
	DefaultStartIndex = 1
)

func NewPaginatedResponse(items []Item) *PaginatedResponse {
	itemsCount := len(items)
	return &PaginatedResponse{
		APIVersion: DefaultAPIVersion,
		Data: Data{
			StartIndex:   DefaultStartIndex,
			ItemsCount:   itemsCount,
			ItemsPerPage: itemsCount,
			Items:        items,
		},
	}
}

// RespondWithJSON message into JSON as http response
func RespondWithJSON(w http.ResponseWriter, r *http.Request, msg string, err string) {
	render.JSON(w, r, CustomResponse{Message: msg, Error: err})
}
