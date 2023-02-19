package repositories

import (
	"stickerfy/app/models"

	"github.com/google/uuid"
)

// ProductRepository is an interface for a product repository
type ProductRepository interface {
	GetAll() ([]models.Product, error)
	GetByID(id uuid.UUID) (models.Product, error)
	Post(product models.Product) error
	Update(product models.Product) error
	Delete(product models.Product) error
}
