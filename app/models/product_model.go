package models

// Product represents a product
type Product struct {
	ProductID   string  `json:"productId"`
	ImagePath   string  `json:"imagePath"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}
