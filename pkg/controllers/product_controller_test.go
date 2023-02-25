package controllers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"

	"stickerfy/app/models"
	"stickerfy/pkg/controllers"
	"stickerfy/pkg/router"
	"stickerfy/pkg/routes"
	mock_cache "stickerfy/test/mocks/cache"
	mock_services "stickerfy/test/mocks/services"

	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestProductController_GetAll tests the GetAll method of the ProductController
func TestProductController_GetAll(t *testing.T) {
	t.Parallel()

	tests := []struct {
		description        string
		cacheError         error
		serviceError       error
		expectedStatusCode int
	}{
		{
			description:        "should return 200 and products",
			cacheError:         redis.Nil,
			serviceError:       nil,
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			mockProductService := mock_services.NewProductService(t)
			mockProductCache := mock_cache.NewCache(t)

			encodedProducts, _ := json.Marshal([]models.Product{
				{
					ID:          uuid.New(),
					Title:       "test",
					Description: "test",
					Price:       1,
					ImagePath:   "image.png",
				},
			})

			mockProductCache.On("Get", mock.Anything, mock.Anything).Return(string(encodedProducts), test.cacheError)
			mockProductService.On("GetAll", mock.Anything).Return([]models.Product{}, test.serviceError)
			mockProductCache.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
			productController := controllers.NewProductController(context.Background(), mockProductService, mockProductCache)

			fr := router.NewFiberRouter()
			routes.ProductRoutes(fr, productController)

			req := httptest.NewRequest(http.MethodGet, "/products", nil)
			req.Header.Set("Content-Type", "application/json")

			res, err := fr.Test(req)

			assert.Nil(t, err)
			assert.Equal(t, test.expectedStatusCode, res.StatusCode)

			mockProductService.AssertExpectations(t)
			mockProductCache.AssertExpectations(t)
		})
	}
}

// TestProductController_GetByID tests the GetByID method of the ProductController
func TestProductController_GetByID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		description        string
		id                 string
		serviceError       error
		expectedStatusCode int
	}{
		{
			description:        "should return 200 and product",
			id:                 "00000000-0000-0000-0000-000000000000",
			serviceError:       nil,
			expectedStatusCode: http.StatusOK,
		},
		{
			description:        "should return 500 and error",
			id:                 "00000000-0000-0000-0000-000000000000",
			serviceError:       errors.New("error"),
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			mockProductService := mock_services.NewProductService(t)
			mockProductService.On("GetByID", mock.Anything, mock.Anything).Return(models.Product{}, test.serviceError)
			productController := controllers.NewProductController(context.Background(), mockProductService, nil)

			fr := router.NewFiberRouter()
			routes.ProductRoutes(fr, productController)

			req := httptest.NewRequest(http.MethodGet, "/product/"+test.id, nil)
			req.Header.Set("Content-Type", "application/json")

			res, err := fr.Test(req)

			assert.Nil(t, err)
			assert.Equal(t, test.expectedStatusCode, res.StatusCode)

			mockProductService.AssertExpectations(t)
		})
	}
}

// TestProductController_Post tests the Post method of the ProductController
func TestProductController_Post(t *testing.T) {
	t.Parallel()

	tests := []struct {
		description        string
		serviceError       error
		expectedStatusCode int
	}{
		{
			description:        "should return 200",
			serviceError:       nil,
			expectedStatusCode: http.StatusOK,
		},
		{
			description:        "should return 500 and error",
			serviceError:       errors.New("error"),
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
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
			mockProductService.On("Post", mock.Anything, mock.Anything).Return(test.serviceError)
			productController := controllers.NewProductController(context.Background(), mockProductService, nil)

			fr := router.NewFiberRouter()
			routes.ProductRoutes(fr, productController)

			req := httptest.NewRequest(http.MethodPost, "/product", body)
			req.Header.Set("Content-Type", "application/json")

			res, err := fr.Test(req)

			assert.Nil(t, err)
			assert.Equal(t, test.expectedStatusCode, res.StatusCode)

			mockProductService.AssertExpectations(t)
		})
	}
}

// TestProductController_Delete tests the Delete method of the ProductController
func TestProductController_Delete(t *testing.T) {
	t.Parallel()

	tests := []struct {
		description        string
		serviceError       error
		expectedStatusCode int
	}{
		{
			description:        "should return 200",
			serviceError:       nil,
			expectedStatusCode: http.StatusOK,
		},
		{
			description:        "should return 500 and error",
			serviceError:       errors.New("error"),
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
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
			mockProductService.On("Delete", mock.Anything, mock.Anything).Return(test.serviceError)
			productController := controllers.NewProductController(context.Background(), mockProductService, nil)

			fr := router.NewFiberRouter()
			routes.ProductRoutes(fr, productController)

			req := httptest.NewRequest(http.MethodDelete, "/product/", body)
			req.Header.Set("Content-Type", "application/json")

			res, err := fr.Test(req)

			assert.Nil(t, err)
			assert.Equal(t, test.expectedStatusCode, res.StatusCode)

			mockProductService.AssertExpectations(t)
		})
	}
}

// TestProductController_Update tests the Update method of the ProductController
func TestProductController_Update(t *testing.T) {
	t.Parallel()

	tests := []struct {
		description        string
		serviceError       error
		expectedStatusCode int
	}{
		{
			description:        "should return 200",
			serviceError:       nil,
			expectedStatusCode: http.StatusOK,
		},
		{
			description:        "should return 500 and error",
			serviceError:       errors.New("error"),
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
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
			mockProductService.On("Update", mock.Anything, fakeProduct).Return(test.serviceError)
			productController := controllers.NewProductController(context.Background(), mockProductService, nil)

			fr := router.NewFiberRouter()
			routes.ProductRoutes(fr, productController)

			req := httptest.NewRequest(http.MethodPut, "/product/", body)
			req.Header.Set("Content-Type", "application/json")

			res, err := fr.Test(req)

			assert.Nil(t, err)
			assert.Equal(t, test.expectedStatusCode, res.StatusCode)

			mockProductService.AssertExpectations(t)
		})
	}
}
