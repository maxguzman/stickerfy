package services_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"stickerfy/app/models"
	"stickerfy/app/services"
	mock_repositories "stickerfy/test/mocks/repositories"
)

// TestGetAllProducts tests ProductService.GetAll
func TestGetAllProducts(t *testing.T) {
	mockProductRepository := mock_repositories.NewProductRepository(t)
	mockProductRepository.On("GetAll").Return([]models.Product{}, nil)

	productService := services.NewProductService(mockProductRepository)

	products, err := productService.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, []models.Product{}, products)
}

// TestGetProductByID tests ProductService.GetByID
func TestGetProductByID(t *testing.T) {
	mockId := uuid.New()
	mockProductRepository := mock_repositories.NewProductRepository(t)
	mockProductRepository.On("GetByID", mockId).Return(models.Product{}, nil)

	productService := services.NewProductService(mockProductRepository)

	product, err := productService.GetByID(mockId)
	assert.Nil(t, err)
	assert.Equal(t, models.Product{}, product)
}

// TestPostProduct tests ProductService.Post
func TestPostProduct(t *testing.T) {
	mockProductRepository := mock_repositories.NewProductRepository(t)
	mockProductRepository.On("Post", models.Product{}).Return(nil)

	productService := services.NewProductService(mockProductRepository)

	err := productService.Post(models.Product{})
	assert.Nil(t, err)
}

// TestUpdateProduct tests ProductService.Update
func TestUpdateProduct(t *testing.T) {
	mockProductRepository := mock_repositories.NewProductRepository(t)
	mockProductRepository.On("Update", models.Product{}).Return(nil)

	productService := services.NewProductService(mockProductRepository)

	err := productService.Update(models.Product{})
	assert.Nil(t, err)
}

// TestDeleteProduct tests ProductService.Delete
func TestDeleteProduct(t *testing.T) {
	mockProductRepository := mock_repositories.NewProductRepository(t)
	mockProductRepository.On("Delete", models.Product{}).Return(nil)

	productService := services.NewProductService(mockProductRepository)

	err := productService.Delete(models.Product{})
	assert.Nil(t, err)
}
