package menu

// A Menu represents menu entity which is a menu of the restaurant
type Menu struct {
	RestaurantID int64  `json:"restaurant-id" `
	Category     string `json:"category"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Price        string `json:"price"`
}

//TODO: to add validators to menu
