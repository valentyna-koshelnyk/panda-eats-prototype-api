package repository

import (
	"encoding/json"
	"fmt"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/restaurant"
	"net/http"
	"strconv"
)

type CustomKey string

// PageIDKey is the key for the page ID parameter
const PageIDKey CustomKey = "page_id"

// Pagination struct to hold pagination details
type Pagination struct {
	PageID     int    `json:"page_id"`
	PageSize   int    `json:"page_size"`
	TotalItems int    `json:"total_items"`
	TotalPages int    `json:"total_pages"`
	PrevPage   string `json:"prev_page,omitempty"`
	NextPage   string `json:"next_page,omitempty"`
}

// PaginateHandler is a handler function that implements pagination
func PaginateHandler(w http.ResponseWriter, r *http.Request) {
	pageIDStr := r.URL.Query().Get(string(PageIDKey))
	pageSize := 10

	pageID, err := strconv.Atoi(pageIDStr)
	if err != nil {
		pageID = 1
	}
	items, totalItems, err := fetchData(pageID, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	totalPages := (totalItems + pageSize - 1) / pageSize
	var prevPage, nextPage string
	if pageID > 1 {
		prevPage = fmt.Sprintf("%s?%s=%d", r.URL.Path, PageIDKey, pageID-1)
	}
	if pageID < totalPages {
		nextPage = fmt.Sprintf("%s?%s=%d", r.URL.Path, PageIDKey, pageID+1)
	}

	pagination := Pagination{
		PageID:     pageID,
		PageSize:   pageSize,
		TotalItems: totalItems,
		TotalPages: totalPages,
		PrevPage:   prevPage,
		NextPage:   nextPage,
	}

	// Render the response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"data":       items,
		"pagination": pagination,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func fetchData(pageID, pageSize int) ([]restaurant.Restaurant, int, error) {
	service := &restaurant.RestaurantService{}
	restaurants, err := service.FindAll()
	if err != nil {
		return nil, 0, err
	}

	startIndex := (pageID - 1) * pageSize
	endIndex := startIndex + pageSize
	if endIndex > len(restaurants) {
		endIndex = len(restaurants)
	}
	paginatedRestaurants := make([]restaurant.Restaurant, 0, endIndex-startIndex)
	for i := startIndex; i < endIndex; i++ {
		paginatedRestaurants = append(paginatedRestaurants, restaurants[i])
	}

	return paginatedRestaurants, len(restaurants), nil
}
