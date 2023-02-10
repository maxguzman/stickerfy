package database

import (
	"stickerfy/app/models"
	"stickerfy/app/repositories"

	"github.com/google/uuid"
)

// ProductRepositoryImpl is a implementation of ProductRepository
type mongoProductRepository struct{}

// NewMongoProductRepository creates a new ProductRepository
func NewMongoProductRepository() repositories.ProductRepository {
	return &mongoProductRepository{}
}

// FindAll returns all products
func (pr *mongoProductRepository) GetAll() ([]models.Product, error) {
	return []models.Product{}, nil
}

// Get returns a product by id
func (pr *mongoProductRepository) GetByID(id uuid.UUID) (models.Product, error) {
	return models.Product{}, nil
}

// New creates a new product
func (pr *mongoProductRepository) Post(product models.Product) error {
	return nil
}

// Update updates a product
func (pr *mongoProductRepository) Update(product models.Product) error {
	return nil
}

// Delete deletes a product by id
func (pr *mongoProductRepository) Delete(product models.Product) error {
	return nil
}

// MongoOrderRepository is a implementation of OrderRepository
type mongoOrderRepository struct{}

// NewMongoOrderRepository creates a new OrderRepository
func NewMongoOrderRepository() repositories.OrderRepository {
	return &mongoOrderRepository{}
}

// GetAll returns all orders
func (or *mongoOrderRepository) GetAll() ([]models.Order, error) {
	return []models.Order{}, nil
}

// Post creates a new order
func (or *mongoOrderRepository) Post(order models.Order) error {
	return nil
}

