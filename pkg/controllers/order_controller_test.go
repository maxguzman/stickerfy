package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"stickerfy/app/models"
	"stickerfy/pkg/controllers"
	router "stickerfy/pkg/router"
	"stickerfy/pkg/routes"
	mock_services "stickerfy/test/mocks/services"

	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// TestOrderController_GetAll tests the GetAll method of the OrderController
func TestOrderController_GetAll(t *testing.T) {
	mockOrderService := mock_services.NewOrderService(t)
	mockOrderService.On("GetAll").Return([]models.Order{}, nil)
	orderController := controllers.NewOrderController(mockOrderService)

	fr := router.NewFiberRouter()
	routes.OrderRoutes(fr, orderController)

	req := httptest.NewRequest(http.MethodGet, "/orders", nil)
	req.Header.Set("Content-Type", "application/json")

	res, err := fr.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	mockOrderService.AssertExpectations(t)
}

// TestOrderController_Post tests the Post method of the OrderController
func TestOrderController_Post(t *testing.T) {
	body := new(bytes.Buffer)
	mockOrder := models.Order{
		Items: []models.OrderItem{
			{
				Product: models.Product{
					ID:          uuid.New(),
					Description: "Test Product",
					Price:       10.0,
					ImagePath:   "test.png",
					Title:       "Test Product",
				},
				Quantity: 1,
			},
		},
	}
	err := json.NewEncoder(body).Encode(&mockOrder)
	assert.Nil(t, err)

	mockOrderService := mock_services.NewOrderService(t)
	mockOrderService.On("Post", mockOrder).Return(nil)
	orderController := controllers.NewOrderController(mockOrderService)

	fr := router.NewFiberRouter()
	routes.OrderRoutes(fr, orderController)

	req := httptest.NewRequest(http.MethodPost, "/order", body)
	req.Header.Set("Content-Type", "application/json")

	res, err := fr.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	mockOrderService.AssertExpectations(t)
}
