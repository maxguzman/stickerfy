package services

import (
	"stickerfy/app/models"
	"stickerfy/app/repositories"
)

// OrderService is an interface for an order service
type OrderService interface {
	GetAll() ([]models.Order, error)
	Post(order models.Order) error
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
func (os *orderService) GetAll() ([]models.Order, error) {
	return os.orderRepository.GetAll()
}

// Post creates a new order
func (os *orderService) Post(order models.Order) error {
	return os.orderRepository.Post(order)
}
