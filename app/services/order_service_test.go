package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"stickerfy/app/models"
	"stickerfy/app/services"
)

// MockOrderRepository is a mock of OrderRepository
type MockOrderRepository struct {
	mock.Mock
}

// GetAll is a mock of OrderRepository.GetAll
func (m *MockOrderRepository) GetAll() ([]models.Order, error) {
	args := m.Called()
	return args.Get(0).([]models.Order), args.Error(1)
}

// Post is a mock of OrderRepository.Post
func (m *MockOrderRepository) Post(order models.Order) error {
	args := m.Called(order)
	return args.Error(0)
}

// TestGetAllOrders tests OrderService.GetAll
func TestGetAllOrders(t *testing.T) {
	mockOrderRepository := new(MockOrderRepository)
	mockOrderRepository.On("GetAll").Return([]models.Order{}, nil)

	orderService := services.NewOrderService(mockOrderRepository)

	orders, err := orderService.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, []models.Order{}, orders)
}

// TestPost tests OrderService.Post
func TestPostOrder(t *testing.T) {
	mockOrderRepository := new(MockOrderRepository)
	mockOrderRepository.On("Post", models.Order{}).Return(nil)

	orderService := services.NewOrderService(mockOrderRepository)

	err := orderService.Post(models.Order{})
	assert.Nil(t, err)
}
