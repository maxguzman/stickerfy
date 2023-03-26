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
	mock_events "stickerfy/test/mocks/events"
	mock_metrics "stickerfy/test/mocks/metrics"
	mock_services "stickerfy/test/mocks/services"

	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var fakeOrder = models.Order{
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

// TestOrderController_GetAll tests the GetAll method of the OrderController
func TestOrderController_GetAll(t *testing.T) {
	t.Parallel()

	tests := []struct {
		description        string
		serviceReturn      interface{}
		serviceError       error
		expectedStatusCode int
	}{
		{
			description:        "should return 200 and orders",
			serviceReturn:      []models.Order{fakeOrder},
			serviceError:       nil,
			expectedStatusCode: http.StatusOK,
		},
		{
			description:        "should return 500 and error when getting orders fails",
			serviceReturn:      []models.Order{fakeOrder},
			serviceError:       errors.New("error"),
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			description:        "should return 404 and error when no orders are found",
			serviceReturn:      nil,
			serviceError:       nil,
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			mockOrderService := mock_services.NewOrderService(t)
			mockOrderService.On("GetAll", mock.Anything).Return(test.serviceReturn, test.serviceError)
			orderController := controllers.NewOrderController(context.Background(), mockOrderService, nil, nil)

			fr := router.NewFiberRouter()
			routes.OrderRoutes(fr, orderController)

			req := httptest.NewRequest(http.MethodGet, "/v1/orders", nil)
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
		metricLabel        string
		serviceError       error
		publishError       error
		expectedStatusCode int
	}{
		{
			description:        "should return 201 and order",
			metricLabel:        "orderAdded",
			serviceError:       nil,
			publishError:       nil,
			expectedStatusCode: http.StatusCreated,
		},
		{
			description:        "should return 500 and error when service fails",
			metricLabel:        "orderFailed",
			serviceError:       errors.New("error"),
			publishError:       nil,
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			description:        "should return 500 and error when event publish fails",
			metricLabel:        "",
			serviceError:       nil,
			publishError:       errors.New("error"),
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			body := new(bytes.Buffer)
			err := json.NewEncoder(body).Encode(&fakeOrder)
			assert.Nil(t, err)

			mockOrderService := mock_services.NewOrderService(t)
			mockMetrics := mock_metrics.NewMetrics(t)
			mockOrderEvent := mock_events.NewEventProducer(t)
			mockOrderService.On("Post", mock.Anything, fakeOrder).Return(test.serviceError)
			if test.metricLabel != "" {
				mockMetrics.On("IncrementCounter", test.metricLabel, mock.Anything).Return(nil)
			}
			if test.serviceError == nil {
				mockOrderEvent.On("Publish", mock.Anything).Return(test.publishError)
			}
			orderController := controllers.NewOrderController(context.Background(), mockOrderService, mockOrderEvent, mockMetrics)

			fr := router.NewFiberRouter()
			routes.OrderRoutes(fr, orderController)

			req := httptest.NewRequest(http.MethodPost, "/v1/order", body)
			req.Header.Set("Content-Type", "application/json")

			res, err := fr.Test(req)

			assert.Nil(t, err)
			assert.Equal(t, test.expectedStatusCode, res.StatusCode)

			mockOrderService.AssertExpectations(t)
			mockMetrics.AssertExpectations(t)
			mockOrderEvent.AssertExpectations(t)
		})
	}
}
