package models

import (
	"github.com/google/uuid"
)

// Order represents an order
type Order struct {
	ID        uuid.UUID   `json:"id" bson:"_id" validate:"required,uuid" format:"uuid"`
	Items     []OrderItem `json:"items" validate:"required,dive"`
}

// OrderItem represents an item in an order
type OrderItem struct {
	Product  Product `json:"product" validate:"required"`
	Quantity int   `json:"quantity" validate:"required,min=1" format:"int32"`
}
