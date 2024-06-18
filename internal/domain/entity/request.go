package utils

// QuantityItemRequest struct for deserialization json body of quantity
type QuantityItemRequest struct {
	Quantity int64 `json:"quantity"`
}
