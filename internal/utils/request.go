package utils

// AddItemRequest struct for deserialization json body of quantity
type AddItemRequest struct {
	Quantity int64 `json:"quantity"`
}
