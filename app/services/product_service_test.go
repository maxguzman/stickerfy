package services_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"stickerfy/app/models"
	"stickerfy/app/services"
	mocks "stickerfy/pkg/mocks/repositories"
)

// TestGetAllProducts tests ProductService.GetAll
func TestGetAllProducts(t *testing.T) {
	mockProductRepository := new(mocks.ProductRepository)
	mockProductRepository.On("GetAll").Return([]models.Product{}, nil)

	productService := services.NewProductService(mockProductRepository)

	products, err := productService.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, []models.Product{}, products)
}

// TestGetProductByID tests ProductService.GetByID
func TestGetProductByID(t *testing.T) {
	mockId := uuid.New()
	mockProductRepository := new(mocks.ProductRepository)
	mockProductRepository.On("GetByID", mockId).Return(models.Product{}, nil)

	productService := services.NewProductService(mockProductRepository)

	product, err := productService.GetByID(mockId)
	assert.Nil(t, err)
	assert.Equal(t, models.Product{}, product)
}

// TestPostProduct tests ProductService.Post
func TestPostProduct(t *testing.T) {
	mockProductRepository := new(mocks.ProductRepository)
	mockProductRepository.On("Post", models.Product{}).Return(nil)

	productService := services.NewProductService(mockProductRepository)

	err := productService.Post(models.Product{})
	assert.Nil(t, err)
}

// TestUpdateProduct tests ProductService.Update
func TestUpdateProduct(t *testing.T) {
	mockProductRepository := new(mocks.ProductRepository)
	mockProductRepository.On("Update", models.Product{}).Return(nil)

	productService := services.NewProductService(mockProductRepository)

	err := productService.Update(models.Product{})
	assert.Nil(t, err)
}

// TestDeleteProduct tests ProductService.Delete
func TestDeleteProduct(t *testing.T) {
	mockProductRepository := new(mocks.ProductRepository)
	mockProductRepository.On("Delete", models.Product{}).Return(nil)

	productService := services.NewProductService(mockProductRepository)

	err := productService.Delete(models.Product{})
	assert.Nil(t, err)
}
