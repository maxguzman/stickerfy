package models

import "github.com/google/uuid"

// Product represents a product
type Product struct {
	ID          uuid.UUID `json:"id" bson:"_id" validate:"required,uuid"`
	ImagePath   string    `json:"image_path" validate:"required,lte=255"`
	Title       string    `json:"title" validate:"required,lte=255"`
	Description string    `json:"description" validate:"required,lte=255"`
	Price       float32   `json:"price" validate:"required,min=0"`
}
