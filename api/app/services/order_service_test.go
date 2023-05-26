package services_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"stickerfy/app/models"
	"stickerfy/app/services"
	mock_repositories "stickerfy/test/mocks/repositories"
)

// TestGetAllOrders tests OrderService.GetAll
func TestGetAllOrders(t *testing.T) {
	t.Parallel()
	mockOrderRepository := new(mock_repositories.OrderRepository)
	mockOrderRepository.On("GetAll", mock.Anything).Return([]models.Order{}, nil)

	orderService := services.NewOrderService(mockOrderRepository)

	orders, err := orderService.GetAll(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, []models.Order{}, orders)
}

// TestPost tests OrderService.Post
func TestPostOrder(t *testing.T) {
	t.Parallel()
	mockOrderRepository := new(mock_repositories.OrderRepository)
	mockOrderRepository.On("Post", mock.Anything, models.Order{}).Return(nil)

	orderService := services.NewOrderService(mockOrderRepository)

	err := orderService.Post(context.Background(), models.Order{})
	assert.Nil(t, err)
}
