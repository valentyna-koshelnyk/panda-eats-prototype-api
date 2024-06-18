package entity

// QuantityItemRequest struct for deserialization json body of quantity
type QuantityItemRequest struct {
	Quantity int64 `json:"quantity"`
}

// OrderIDRequest struct for deserialization json body of order ID
type OrderIDRequest struct {
	OrderID string `json:"order_id"`
}
