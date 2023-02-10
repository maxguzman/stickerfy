package services

import (
	"stickerfy/app/models"
	"stickerfy/app/repositories"

	"github.com/google/uuid"
)

// ProductService is an interface for a product service
type ProductService interface {
	GetAll() ([]models.Product, error)
	GetByID(id uuid.UUID) (models.Product, error)
	Post(product models.Product) error
	Delete(product models.Product) error
	Update(product models.Product) error
}

// productService is a implementation of ProductService
type productService struct {
	productRepository repositories.ProductRepository
}

// NewProductService creates a new ProductService
func NewProductService(productRepository repositories.ProductRepository) ProductService {
	return &productService{
		productRepository: productRepository,
	}
}

// GetAll returns all products
func (ps *productService) GetAll() ([]models.Product, error) {
	return ps.productRepository.GetAll()
}

// Get returns a product by id
func (ps *productService) GetByID(id uuid.UUID) (models.Product, error) {
	return ps.productRepository.GetByID(id)
}

// New creates a new product
func (ps *productService) Post(product models.Product) error {
	return ps.productRepository.Post(product)
}

// Delete deletes a product
func (ps *productService) Delete(product models.Product) error {
	return ps.productRepository.Delete(product)
}

// Update updates a product
func (ps *productService) Update(product models.Product) error {
	return ps.productRepository.Update(product)
}
