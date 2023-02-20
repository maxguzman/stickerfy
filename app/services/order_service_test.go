package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"stickerfy/app/models"
	"stickerfy/app/services"
	mock_repositories "stickerfy/test/mocks/repositories"
)

// TestGetAllOrders tests OrderService.GetAll
func TestGetAllOrders(t *testing.T) {
	mockOrderRepository := new(mock_repositories.OrderRepository)
	mockOrderRepository.On("GetAll").Return([]models.Order{}, nil)

	orderService := services.NewOrderService(mockOrderRepository)

	orders, err := orderService.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, []models.Order{}, orders)
}

// TestPost tests OrderService.Post
func TestPostOrder(t *testing.T) {
	mockOrderRepository := new(mock_repositories.OrderRepository)
	mockOrderRepository.On("Post", models.Order{}).Return(nil)

	orderService := services.NewOrderService(mockOrderRepository)

	err := orderService.Post(models.Order{})
	assert.Nil(t, err)
}
