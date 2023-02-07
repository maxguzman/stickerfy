package services

import (
	"stickerfy/app/models"
	"stickerfy/app/repositories"
)

// ProductService is an interface for a product service
type ProductService interface {
	GetAll() ([]models.Product, error)
	GetByID(id string) (models.Product, error)
	Post(product models.Product) error
	Delete(id string) error
	Update(product models.Product) error
}

// NewProductService creates a new ProductService
func NewProductService(productRepository repositories.ProductRepository) ProductService {
	return &productService{
		productRepository: productRepository,
	}
}

// productService is a implementation of ProductService
type productService struct {
	productRepository repositories.ProductRepository
}

// GetAll returns all products
func (ps *productService) GetAll() ([]models.Product, error) {
	return ps.productRepository.GetAll()
}

// New creates a new product
func (ps *productService) Post(product models.Product) error {
	return ps.productRepository.Post(product)
}

// Get returns a product by id
func (ps *productService) GetByID(id string) (models.Product, error) {
	return ps.productRepository.GetByID(id)
}

// Delete deletes a product by id
func (ps *productService) Delete(id string) error {
	return ps.productRepository.Delete(id)
}

// Update updates a product
func (ps *productService) Update(product models.Product) error {
	return ps.productRepository.Update(product)
}
