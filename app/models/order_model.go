package models

// Order represents an order
type Order struct {
	OrderID string      `json:"order_id"`
	Items   []OrderItem `json:"items"`
}

// OrderItem represents an item in an order
type OrderItem struct {
	Product  Product `json:"product"`
	Quantity int32   `json:"quantity"`
}
