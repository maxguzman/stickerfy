package database_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"stickerfy/app/models"
	"stickerfy/platform/database"
)

// MockMongoProductRepository is a mock of MongoProductRepository
type MockMongoProductRepository struct {
	mock.Mock
}

// FindAll is a mock of MongoProductRepository.FindAll
func (m *MockMongoProductRepository) GetAll(collection string, result interface{}) error {
	args := m.Called(collection, result)
	return args.Error(0)
}

// GetByID is a mock of MongoProductRepository.GetByID
func (m *MockMongoProductRepository) GetByID(collection string, id string, result interface{}) error {
	args := m.Called(collection, id, result)
	return args.Error(0)
}

// Post is a mock of MongoProductRepository.Post
func (m *MockMongoProductRepository) Post(collection string, product models.Product) error {
	args := m.Called(collection, product)
	return args.Error(0)
}

// Update is a mock of MongoProductRepository.Update
func (m *MockMongoProductRepository) Update(collection string, id string, product models.Product) error {
	args := m.Called(collection, id, product)
	return args.Error(0)
}

// Delete is a mock of MongoProductRepository.Delete
func (m *MockMongoProductRepository) Delete(collection string, id string) error {
	args := m.Called(collection, id)
	return args.Error(0)
}

// TestFindAllProducts tests MongoProductRepository.FindAll
func TestFindAllProducts(t *testing.T) {
	mockMongoProductRepository := new(MockMongoProductRepository)
	mockMongoProductRepository.On("FindAll", "products", &[]models.Product{}).Return(nil)

	mongoProductRepository := database.NewMongoProductRepository()

	products, err := mongoProductRepository.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, []models.Product{}, products)
}