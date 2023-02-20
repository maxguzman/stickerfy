package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"stickerfy/app/models"
	"stickerfy/pkg/controllers"
	"stickerfy/pkg/router"
	"stickerfy/pkg/routes"
	mock_services "stickerfy/test/mocks/services"

	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// TestProductController_GetAll tests the GetAll method of the ProductController
func TestProductController_GetAll(t *testing.T) {
	mockProductService := mock_services.NewProductService(t)
	mockProductService.On("GetAll").Return([]models.Product{}, nil)
	productController := controllers.NewProductController(mockProductService)

	fr := router.NewFiberRouter()
	routes.ProductRoutes(fr, productController)

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	req.Header.Set("Content-Type", "application/json")

	res, err := fr.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	mockProductService.AssertExpectations(t)
}

// TestProductController_Post tests the Post method of the ProductController
func TestProductController_Post(t *testing.T) {
	body := new(bytes.Buffer)
	fakeProduct := models.Product{
		ID:          uuid.New(),
		Description: "Test product",
		Price:       10.0,
		Title:       "Test product",
		ImagePath:   "https://example.com/image.png",
	}
	err := json.NewEncoder(body).Encode(&fakeProduct)
	assert.Nil(t, err)

	mockProductService := mock_services.NewProductService(t)
	mockProductService.On("Post", fakeProduct).Return(nil)
	productController := controllers.NewProductController(mockProductService)

	fr := router.NewFiberRouter()
	routes.ProductRoutes(fr, productController)

	req := httptest.NewRequest(http.MethodPost, "/product", body)
	req.Header.Set("Content-Type", "application/json")

	res, err := fr.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	mockProductService.AssertExpectations(t)
}

// TestProductController_GetByID tests the GetByID method of the ProductController
func TestProductController_GetByID(t *testing.T) {
	fakeID := uuid.New()
	fakeProduct := models.Product{
		ID:          fakeID,
		Description: "Test product",
		Price:       10.0,
		Title:       "Test product",
		ImagePath:   "https://example.com/image.png",
	}
	mockProductService := mock_services.NewProductService(t)
	mockProductService.On("GetByID", fakeID).Return(fakeProduct, nil)
	productController := controllers.NewProductController(mockProductService)

	fr := router.NewFiberRouter()
	routes.ProductRoutes(fr, productController)

	req := httptest.NewRequest(http.MethodGet, "/product/"+fakeID.String(), nil)
	req.Header.Set("Content-Type", "application/json")

	res, err := fr.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	mockProductService.AssertExpectations(t)
}

// TestProductController_Delete tests the Delete method of the ProductController
func TestProductController_Delete(t *testing.T) {
	body := new(bytes.Buffer)
	fakeID := uuid.New()
	fakeProduct := models.Product{
		ID:          fakeID,
		Description: "Test product",
		Price:       10.0,
		Title:       "Test product",
		ImagePath:   "https://example.com/image.png",
	}
	err := json.NewEncoder(body).Encode(&fakeProduct)
	assert.Nil(t, err)

	mockProductService := mock_services.NewProductService(t)
	mockProductService.On("Delete", fakeProduct).Return(nil)
	productController := controllers.NewProductController(mockProductService)

	fr := router.NewFiberRouter()
	routes.ProductRoutes(fr, productController)

	req := httptest.NewRequest(http.MethodDelete, "/product/", body)
	req.Header.Set("Content-Type", "application/json")

	res, err := fr.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	mockProductService.AssertExpectations(t)
}

// TestProductController_Update tests the Update method of the ProductController
func TestProductController_Update(t *testing.T) {
	body := new(bytes.Buffer)
	fakeID := uuid.New()
	fakeProduct := models.Product{
		ID:          fakeID,
		Description: "Test product",
		Price:       10.0,
		Title:       "Test product",
		ImagePath:   "https://example.com/image.png",
	}
	err := json.NewEncoder(body).Encode(&fakeProduct)
	assert.Nil(t, err)

	mockProductService := mock_services.NewProductService(t)
	mockProductService.On("Update", fakeProduct).Return(nil)
	productController := controllers.NewProductController(mockProductService)

	fr := router.NewFiberRouter()
	routes.ProductRoutes(fr, productController)

	req := httptest.NewRequest(http.MethodPut, "/product/", body)
	req.Header.Set("Content-Type", "application/json")

	res, err := fr.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	mockProductService.AssertExpectations(t)
}
