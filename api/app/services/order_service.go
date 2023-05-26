package services

import (
	"context"
	"stickerfy/app/models"
	"stickerfy/app/repositories"
)

// OrderService is an interface for an order service
type OrderService interface {
	GetAll(ctx context.Context) ([]models.Order, error)
	Post(ctx context.Context, order models.Order) error
}

// orderService is a implementation of OrderService
type orderService struct {
	orderRepository repositories.OrderRepository
}

// NewOrderService creates a new OrderService
func NewOrderService(orderRepository repositories.OrderRepository) OrderService {
	return &orderService{
		orderRepository: orderRepository,
	}
}

// GetAll returns all orders
func (os *orderService) GetAll(ctx context.Context) ([]models.Order, error) {
	return os.orderRepository.GetAll(ctx)
}

// Post creates a new order
func (os *orderService) Post(ctx context.Context, order models.Order) error {
	return os.orderRepository.Post(ctx, order)
}
