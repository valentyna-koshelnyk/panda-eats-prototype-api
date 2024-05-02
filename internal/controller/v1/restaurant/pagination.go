package restaurant

import (
	"encoding/json"
	"fmt"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/restaurant"
	"net/http"
	"strconv"
)

// CustomKey is a custom type for the key to retrieve the page ID param
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

// GetAllRestaurants is a handler function that implements pagination
func GetAllRestaurants(w http.ResponseWriter, r *http.Request) {
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

// GetByCategoryPriceZip is a handler for filtering by category AND/OR price range AND/OR zip
func GetByCategoryPriceZip(w http.ResponseWriter, r *http.Request) {
	pageIDStr := r.URL.Query().Get(string(PageIDKey))
	pageSize := 10

	pageID, err := strconv.Atoi(pageIDStr)
	if err != nil {
		pageID = 1
	}

	category := r.URL.Query().Get("category")
	priceRange := r.URL.Query().Get("price_range")
	zipCode := r.URL.Query().Get("zip_code")

	// Initialize restaurant service
	service := restaurant.NewRestaurantService()

	// Filter restaurants by category, price, and zip before pagination
	allRestaurants := service.FilterByCategoryPriceZip(category, priceRange, zipCode)

	// Calculate pagination details
	totalItems := len(allRestaurants)
	totalPages := (totalItems + pageSize - 1) / pageSize

	// Correct boundary checks for pagination
	startIndex := (pageID - 1) * pageSize
	if startIndex > totalItems {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
	endIndex := startIndex + pageSize
	if endIndex > totalItems {
		endIndex = totalItems
	}

	// Slice the specific page's data
	restaurants := allRestaurants[startIndex:endIndex]

	// Pagination URLs
	var prevPage, nextPage string
	if pageID > 1 {
		prevPage = fmt.Sprintf("%s?%s=%d&category=%s&price_range=%s&zip_code=%s", r.URL.Path, PageIDKey, pageID-1, category, priceRange, zipCode)
	}
	if pageID < totalPages {
		nextPage = fmt.Sprintf("%s?%s=%d&category=%s&price_range=%s&zip_code=%s", r.URL.Path, PageIDKey, pageID+1, category, priceRange, zipCode)
	}

	// Prepare the pagination object
	pagination := map[string]interface{}{
		"page_id":     pageID,
		"page_size":   pageSize,
		"total_items": totalItems,
		"total_pages": totalPages,
		"prev_page":   prevPage,
		"next_page":   nextPage,
	}

	// Prepare the response
	response := map[string]interface{}{
		"data":       restaurants,
		"pagination": pagination,
	}

	// Set content type and encode the response as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func fetchData(pageID, pageSize int) ([]restaurant.Restaurant, int, error) {
	service := restaurant.NewRestaurantService()
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
