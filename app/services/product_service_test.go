package services_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"stickerfy/app/models"
	"stickerfy/app/services"
	mock_repositories "stickerfy/test/mocks/repositories"
)

// TestGetAllProducts tests ProductService.GetAll
func TestGetAllProducts(t *testing.T) {
	t.Parallel()
	mockProductRepository := mock_repositories.NewProductRepository(t)
	mockProductRepository.On("GetAll", mock.Anything).Return([]models.Product{}, nil)

	productService := services.NewProductService(mockProductRepository)

	products, err := productService.GetAll(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, []models.Product{}, products)
}

// TestGetProductByID tests ProductService.GetByID
func TestGetProductByID(t *testing.T) {
	t.Parallel()
	mockId := uuid.New()
	mockProductRepository := mock_repositories.NewProductRepository(t)
	mockProductRepository.On("GetByID", mock.Anything, mockId).Return(models.Product{}, nil)

	productService := services.NewProductService(mockProductRepository)

	product, err := productService.GetByID(context.Background(), mockId)
	assert.Nil(t, err)
	assert.Equal(t, models.Product{}, product)
}

// TestPostProduct tests ProductService.Post
func TestPostProduct(t *testing.T) {
	t.Parallel()
	mockProductRepository := mock_repositories.NewProductRepository(t)
	mockProductRepository.On("Post", mock.Anything, models.Product{}).Return(nil)

	productService := services.NewProductService(mockProductRepository)

	err := productService.Post(context.Background(), models.Product{})
	assert.Nil(t, err)
}

// TestUpdateProduct tests ProductService.Update
func TestUpdateProduct(t *testing.T) {
	t.Parallel()
	mockProductRepository := mock_repositories.NewProductRepository(t)
	mockProductRepository.On("Update", mock.Anything, models.Product{}).Return(nil)

	productService := services.NewProductService(mockProductRepository)

	err := productService.Update(context.Background(), models.Product{})
	assert.Nil(t, err)
}

// TestDeleteProduct tests ProductService.Delete
func TestDeleteProduct(t *testing.T) {
	t.Parallel()
	mockProductRepository := mock_repositories.NewProductRepository(t)
	mockProductRepository.On("Delete", mock.Anything, models.Product{}).Return(nil)

	productService := services.NewProductService(mockProductRepository)

	err := productService.Delete(context.Background(), models.Product{})
	assert.Nil(t, err)
}
