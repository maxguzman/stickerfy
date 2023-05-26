package services

import (
	"context"
	"stickerfy/app/models"
	"stickerfy/app/repositories"

	"github.com/google/uuid"
)

// ProductService is an interface for a product service
type ProductService interface {
	GetAll(ctx context.Context) ([]models.Product, error)
	GetByID(ctx context.Context, id uuid.UUID) (models.Product, error)
	Post(ctx context.Context, product models.Product) error
	Delete(ctx context.Context, product models.Product) error
	Update(ctx context.Context, product models.Product) error
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
func (ps *productService) GetAll(ctx context.Context) ([]models.Product, error) {
	return ps.productRepository.GetAll(ctx)
}

// Get returns a product by id
func (ps *productService) GetByID(ctx context.Context, id uuid.UUID) (models.Product, error) {
	return ps.productRepository.GetByID(ctx, id)
}

// New creates a new product
func (ps *productService) Post(ctx context.Context, product models.Product) error {
	return ps.productRepository.Post(ctx, product)
}

// Delete deletes a product
func (ps *productService) Delete(ctx context.Context, product models.Product) error {
	return ps.productRepository.Delete(ctx, product)
}

// Update updates a product
func (ps *productService) Update(ctx context.Context, product models.Product) error {
	return ps.productRepository.Update(ctx, product)
}
