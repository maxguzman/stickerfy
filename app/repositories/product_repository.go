package repositories

import "stickerfy/app/models"

// ProductRepository is an interface for a product repository
type ProductRepository interface {
	GetAll() ([]models.Product, error)
	GetByID(id string) (models.Product, error)
	Post(product models.Product) error
	Delete(id string) error
	Update(product models.Product) error
}
