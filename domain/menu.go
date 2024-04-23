package domain

type Menu struct {
	MenuID       int64  `json:"menu_id" `
	RestaurantID int64  `json:"restaurant_id"`
	Category     string `json:"category"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Price        string `json:"price"`
}

//TODO: to add validators to menu
