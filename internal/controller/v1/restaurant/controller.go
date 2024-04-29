package restaurant

// GetRestaurantById a handler for retrieving information about the restaurant based on its id
//func GetRestaurantById(w http.ResponseWriter, r *http.Request) {
//	idStr := chi.URLParam(r, "id")
//
//	service := restaurant.NewRestaurantService()
//	id, err := strconv.ParseInt(idStr, 10, 64)
//
//	restaurant, err := service.FindById(id)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusNotFound)
//		return
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	err = json.NewEncoder(w).Encode(restaurant)
//	if err != nil {
//		return
//	}
//}

//func GetByCategoryPriceZip(w http.ResponseWriter, r *http.Request) {
//	// Get the pagination parameters from the request
//	pageIDStr := r.URL.Query().Get("page_id")
//	pageSizeStr := r.URL.Query().Get("page_size")
//
//	// Parse the pagination parameters
//	pageID, err := strconv.Atoi(pageIDStr)
//	if err != nil {
//		pageID = 1
//	}
//
//	pageSize, err := strconv.Atoi(pageSizeStr)
//	if err != nil {
//		pageSize = 10
//	}
//
//	category := r.URL.Query().Get("category")
//	price_range := r.URL.Query().Get("price_range")
//	zip_code := r.URL.Query().Get("zip_code")
//	var service = restaurant.NewRestaurantService()
//	var allRestaurants []restaurant.Restaurant
////	allRestaurants, err = service.FindByCategoryPriceZip(category, price_range, zip_code)
//
//	// Apply pagination to the restaurants
//	startIndex := (pageID - 1) * pageSize
//	endIndex := startIndex + pageSize
//	if endIndex > len(allRestaurants) {
//		endIndex = len(allRestaurants)
//	}
//	restaurants := allRestaurants[startIndex:endIndex]
//
//	// Calculate the pagination details
//	totalItems := len(allRestaurants)
//	totalPages := (totalItems + pageSize - 1) / pageSize
//	var prevPage, nextPage string
//	if pageID > 1 {
//		prevPage = fmt.Sprintf("%s?page_id=%d&page_size=%d", r.URL.Path, pageID-1, pageSize)
//	}
//	if pageID < totalPages {
//		nextPage = fmt.Sprintf("%s?page_id=%d&page_size=%d", r.URL.Path, pageID+1, pageSize)
//	}
//
//	// Prepare the response
//	response := map[string]interface{}{
//		"data": restaurants,
//		"pagination": map[string]interface{}{
//			"page_id":     pageID,
//			"page_size":   pageSize,
//			"total_items": totalItems,
//			"total_pages": totalPages,
//			"prev_page":   prevPage,
//			"next_page":   nextPage,
//		},
//	}
//
//	// Encode the response as JSON
//	w.Header().Set("Content-Type", "application/json")
//	err = json.NewEncoder(w).Encode(response)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//}
