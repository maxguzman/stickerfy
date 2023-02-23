package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"stickerfy/app/models"
	"stickerfy/pkg/controllers"
	"stickerfy/pkg/router"
	"stickerfy/pkg/routes"
	mock_events "stickerfy/test/mocks/events"
	mock_services "stickerfy/test/mocks/services"

	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestOrderController_GetAll tests the GetAll method of the OrderController
func TestOrderController_GetAll(t *testing.T) {
	t.Parallel()

	tests := []struct {
		description        string
		serviceError       error
		expectedStatusCode int
	}{
		{
			description:        "should return 200 and orders",
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
			mockOrderService := mock_services.NewOrderService(t)
			mockOrderService.On("GetAll").Return([]models.Order{}, test.serviceError)
			orderController := controllers.NewOrderController(mockOrderService, nil)

			fr := router.NewFiberRouter()
			routes.OrderRoutes(fr, orderController)

			req := httptest.NewRequest(http.MethodGet, "/orders", nil)
			req.Header.Set("Content-Type", "application/json")

			res, err := fr.Test(req)

			assert.Nil(t, err)
			assert.Equal(t, test.expectedStatusCode, res.StatusCode)

			mockOrderService.AssertExpectations(t)
		})
	}
}

// TestOrderController_Post tests the Post method of the OrderController
func TestOrderController_Post(t *testing.T) {
	t.Parallel()

	tests := []struct {
		description        string
		serviceError       error
		expectedStatusCode int
	}{
		{
			description:        "should return 201 and order",
			serviceError:       nil,
			expectedStatusCode: http.StatusCreated,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
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
			body := new(bytes.Buffer)
			err := json.NewEncoder(body).Encode(&mockOrder)
			assert.Nil(t, err)

			mockOrderService := mock_services.NewOrderService(t)
			mockOrderEvent := mock_events.NewEventProducer(t)
			mockOrderService.On("Post", mockOrder).Return(test.serviceError)
			mockOrderEvent.On("Publish", mock.Anything, mock.Anything).Return(nil)
			orderController := controllers.NewOrderController(mockOrderService, mockOrderEvent)

			fr := router.NewFiberRouter()
			routes.OrderRoutes(fr, orderController)

			req := httptest.NewRequest(http.MethodPost, "/order", body)
			req.Header.Set("Content-Type", "application/json")

			res, err := fr.Test(req)

			assert.Nil(t, err)
			assert.Equal(t, test.expectedStatusCode, res.StatusCode)

			mockOrderService.AssertExpectations(t)
			mockOrderEvent.AssertExpectations(t)
		})
	}
}
