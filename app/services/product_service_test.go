package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"stickerfy/app/models"
	"stickerfy/app/services"
)

// MockProductRepository is a mock of ProductRepository
type MockProductRepository struct {
	mock.Mock
}

// FindAll is a mock of ProductRepository.FindAll
func (m *MockProductRepository) GetAll() ([]models.Product, error) {
	args := m.Called()
	return args.Get(0).([]models.Product), args.Error(1)
}

// GetByID is a mock of ProductRepository.GetByID
func (m *MockProductRepository) GetByID(id string) (models.Product, error) {
	args := m.Called(id)
	return args.Get(0).(models.Product), args.Error(1)
}

// Post is a mock of ProductRepository.Post
func (m *MockProductRepository) Post(product models.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

// Update is a mock of ProductRepository.Update
func (m *MockProductRepository) Update(product models.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

// Delete is a mock of ProductRepository.Delete
func (m *MockProductRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// TestFindAllProducts tests ProductService.FindAll
func TestGetAllProducts(t *testing.T) {
	mockProductRepository := new(MockProductRepository)
	mockProductRepository.On("FindAll").Return([]models.Product{}, nil)

	productService := services.NewProductService(mockProductRepository)

	products, err := productService.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, []models.Product{}, products)
}

// TestGetProductByID tests ProductService.GetByID
func TestGetProductByID(t *testing.T) {
	mockProductRepository := new(MockProductRepository)
	mockProductRepository.On("GetByID", "1").Return(models.Product{}, nil)

	productService := services.NewProductService(mockProductRepository)

	product, err := productService.GetByID("1")
	assert.Nil(t, err)
	assert.Equal(t, models.Product{}, product)
}

// TestPostProduct tests ProductService.Post
func TestPostProduct(t *testing.T) {
	mockProductRepository := new(MockProductRepository)
	mockProductRepository.On("Post", models.Product{}).Return(nil)

	productService := services.NewProductService(mockProductRepository)

	err := productService.Post(models.Product{})
	assert.Nil(t, err)
}

// TestUpdateProduct tests ProductService.Update
func TestUpdateProduct(t *testing.T) {
	mockProductRepository := new(MockProductRepository)
	mockProductRepository.On("Update", models.Product{}).Return(nil)

	productService := services.NewProductService(mockProductRepository)

	err := productService.Update(models.Product{})
	assert.Nil(t, err)
}

// TestDeleteProduct tests ProductService.Delete
func TestDeleteProduct(t *testing.T) {
	mockProductRepository := new(MockProductRepository)
	mockProductRepository.On("Delete", "1").Return(nil)

	productService := services.NewProductService(mockProductRepository)

	err := productService.Delete("1")
	assert.Nil(t, err)
}
