package repositories

import (
	"stickerfy/app/models"

	"stickerfy/pkg/platform/database"

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

// productRepository is a implementation of ProductRepository
type productRepository struct{
	client database.Client
}

// NewProductRepository creates a new ProductRepository
func NewProductRepository(c database.Client) ProductRepository {
	return &productRepository{
		client: c,
	}
}

// FindAll returns all products
func (pr *productRepository) GetAll() ([]models.Product, error) {
	return []models.Product{}, nil
}

// Get returns a product by id
func (pr *productRepository) GetByID(id uuid.UUID) (models.Product, error) {
	return models.Product{}, nil
}

// New creates a new product
func (pr *productRepository) Post(product models.Product) error {
	return nil
}

// Update updates a product
func (pr *productRepository) Update(product models.Product) error {
	return nil
}

// Delete deletes a product by id
func (pr *productRepository) Delete(product models.Product) error {
	return nil
}
