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

var fakeProduct = models.Product{
	ID:          uuid.New(),
	Title:       "fake title",
	Description: "fake description",
	Price:       1,
	ImagePath:   "fake_image.png",
}

// TestProductController_GetAll tests the GetAll method of the ProductController
func TestProductController_GetAll(t *testing.T) {
	t.Parallel()

	fakeProducts := []models.Product{fakeProduct, fakeProduct}
	encodedProducts, _ := json.Marshal(fakeProducts)

	tests := []struct {
		description        string
		getCacheError      error
		serviceProducts    []models.Product
		serviceError       error
		setCacheError      error
		expectedStatusCode int
	}{
		{
			description:        "should return 200 and products when no cache hit",
			getCacheError:      redis.Nil,
			serviceProducts:    fakeProducts,
			serviceError:       nil,
			setCacheError:      nil,
			expectedStatusCode: http.StatusOK,
		},
		{
			description:        "should return 500 and error when no cache hit and getting from service fails",
			getCacheError:      redis.Nil,
			serviceProducts:    fakeProducts,
			serviceError:       errors.New("error"),
			setCacheError:      nil,
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			description:        "should return 404 and error when no cache hit and no products",
			getCacheError:      redis.Nil,
			serviceProducts:    nil,
			serviceError:       nil,
			setCacheError:      nil,
			expectedStatusCode: http.StatusNotFound,
		},
		{
			description:        "should return 500 and error when no cache hit and set cache fails",
			getCacheError:      redis.Nil,
			serviceProducts:    fakeProducts,
			serviceError:       nil,
			setCacheError:      errors.New("error"),
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			description:        "should return 500 and error when getting from cache fails",
			getCacheError:      errors.New("error"),
			serviceProducts:    nil,
			serviceError:       nil,
			setCacheError:      nil,
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			description:        "should return 200 and products when cache hit",
			getCacheError:      nil,
			serviceProducts:    nil,
			serviceError:       nil,
			setCacheError:      nil,
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			mockProductService := mock_services.NewProductService(t)
			mockProductCache := mock_cache.NewCache(t)
			mockProductCache.On("Get", mock.Anything, mock.Anything).Return(string(encodedProducts), test.getCacheError)
			if test.getCacheError == redis.Nil {
				mockProductService.On("GetAll", mock.Anything).Return(test.serviceProducts, test.serviceError)
			}
			if test.serviceError == nil && test.serviceProducts != nil {
				mockProductCache.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(test.setCacheError)
			}
			productController := controllers.NewProductController(context.Background(), mockProductService, mockProductCache)

			fr := router.NewFiberRouter()
			routes.ProductRoutes(fr, productController)

			req := httptest.NewRequest(http.MethodGet, "/v1/products", nil)
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
		{
			description:        "should return 400 and error",
			id:                 "wrong-id",
			serviceError:       errors.New("error"),
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			mockProductService := mock_services.NewProductService(t)
			if test.id != "wrong-id" {
				mockProductService.On("GetByID", mock.Anything, mock.Anything).Return(models.Product{}, test.serviceError)
			}
			productController := controllers.NewProductController(context.Background(), mockProductService, nil)

			fr := router.NewFiberRouter()
			routes.ProductRoutes(fr, productController)

			req := httptest.NewRequest(http.MethodGet, "/v1/product/"+test.id, nil)
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
			err := json.NewEncoder(body).Encode(&fakeProduct)
			assert.Nil(t, err)

			mockProductService := mock_services.NewProductService(t)
			mockProductService.On("Post", mock.Anything, mock.Anything).Return(test.serviceError)
			productController := controllers.NewProductController(context.Background(), mockProductService, nil)

			fr := router.NewFiberRouter()
			routes.ProductRoutes(fr, productController)

			req := httptest.NewRequest(http.MethodPost, "/v1/product", body)
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
		product            interface{}
		serviceError       error
		expectedStatusCode int
	}{
		{
			description:        "should return 200",
			product:            fakeProduct,
			serviceError:       nil,
			expectedStatusCode: http.StatusOK,
		},
		{
			description:        "should return 500 and error",
			product:            fakeProduct,
			serviceError:       errors.New("error"),
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			body := new(bytes.Buffer)
			err := json.NewEncoder(body).Encode(&test.product)
			assert.Nil(t, err)
			mockProductService := mock_services.NewProductService(t)
			mockProductService.On("Update", mock.Anything, mock.Anything).Return(test.serviceError)
			productController := controllers.NewProductController(context.Background(), mockProductService, nil)

			fr := router.NewFiberRouter()
			routes.ProductRoutes(fr, productController)

			req := httptest.NewRequest(http.MethodPut, "/v1/product/", body)
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
			description:        "should return 500 and error when service returns error",
			serviceError:       errors.New("error"),
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			body := new(bytes.Buffer)
			err := json.NewEncoder(body).Encode(&fakeProduct)
			assert.Nil(t, err)

			mockProductService := mock_services.NewProductService(t)
			mockProductService.On("Delete", mock.Anything, mock.Anything).Return(test.serviceError)
			productController := controllers.NewProductController(context.Background(), mockProductService, nil)

			fr := router.NewFiberRouter()
			routes.ProductRoutes(fr, productController)

			req := httptest.NewRequest(http.MethodDelete, "/v1/product/", body)
			req.Header.Set("Content-Type", "application/json")

			res, err := fr.Test(req)

			assert.Nil(t, err)
			assert.Equal(t, test.expectedStatusCode, res.StatusCode)

			mockProductService.AssertExpectations(t)
		})
	}
}
