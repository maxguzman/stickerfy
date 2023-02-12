package repositories

import (
	"stickerfy/app/models"
	"stickerfy/pkg/platform/database"
)

// OrderRepository is an interface for an order repository
type OrderRepository interface {
	GetAll() ([]models.Order, error)
	Post(order models.Order) error
}

// orderRepository is a implementation of OrderRepository
type orderRepository struct{
	client database.Client
}

// NewOrderRepository creates a new OrderRepository
func NewOrderRepository(c database.Client) OrderRepository {
	return &orderRepository{
		client: c,
	}
}

// GetAll returns all orders
func (or *orderRepository) GetAll() ([]models.Order, error) {
	return []models.Order{}, nil
}

// Post creates a new order
func (or *orderRepository) Post(order models.Order) error {
	return nil
}
