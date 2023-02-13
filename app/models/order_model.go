package models

import "github.com/google/uuid"

// Order represents an order
type Order struct {
	ID    uuid.UUID   `json:"id" validate:"required,uuid"`
	Items []OrderItem `json:"items" validate:"required,dive"`
}

// OrderItem represents an item in an order
type OrderItem struct {
	Product  Product `json:"product" validate:"required"`
	Quantity int32   `json:"quantity" validate:"required,min=1"`
}
