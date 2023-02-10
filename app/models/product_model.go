package models

// Product represents a product
type Product struct {
	ProductID   string  `json:"product_id"`
	ImagePath   string  `json:"image_path"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}
