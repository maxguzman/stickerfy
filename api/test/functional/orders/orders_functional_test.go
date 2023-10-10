package functional_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"stickerfy/app/models"
	"stickerfy/pkg/controllers"
	"stickerfy/pkg/router"
	"stickerfy/pkg/routes"
	mock_events "stickerfy/test/mocks/events"
	mock_metrics "stickerfy/test/mocks/metrics"
	mock_services "stickerfy/test/mocks/services"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type apiOrderFeature struct {
	fr     router.Router
	resp   *http.Response
	orders []models.Order
}

func (a *apiOrderFeature) theFollowingOrders(givenOrders *godog.DocString) error {
	if err := json.Unmarshal([]byte(givenOrders.Content), &a.orders); err != nil {
		return err
	}
	return nil
}

func (a *apiOrderFeature) iSendRequestToWithBody(method, path string, body *godog.DocString) error {
	var expectedBody interface{}
	if err := json.Unmarshal([]byte(body.Content), &expectedBody); err != nil {
		return err
	}
	t := &testing.T{}
	mockOrderService := mock_services.NewOrderService(t)
	mockOrderService.On("GetAll", mock.Anything).Return(a.orders, nil)
	mockOrderService.On("Post", mock.Anything, mock.Anything).Return(nil)
	mockOrderEvent := mock_events.NewEventProducer(t)
	mockOrderEvent.On("Publish", mock.Anything).Return(nil)
	mockMetrics := mock_metrics.NewMetrics(t)
	mockMetrics.On("IncrementCounter", mock.Anything, mock.Anything).Return(nil)
	orderController := controllers.NewOrderController(context.Background(), mockOrderService, mockOrderEvent, mockMetrics)
	routes.OrderRoutes(a.fr, orderController)

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(&expectedBody)
	if err != nil {
		return err
	}
	req := httptest.NewRequest(method, path, b)
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.fr.Test(req)
	if err != nil {
		return err
	}
	a.resp = resp

	mockOrderService.AssertExpectations(t)
	return nil
}

func (a *apiOrderFeature) theResponseCodeShouldBe(code int) error {
	return assertExpectedAndActual(
		assert.Equal, code, a.resp.StatusCode,
		"Expected %d code, but got is %d", code, a.resp.StatusCode,
	)
}

func (a *apiOrderFeature) theResponseShouldMatchJson(body *godog.DocString) error {
	var expected, actual interface{}
	if err := json.Unmarshal([]byte(body.Content), &expected); err != nil {
		return err
	}
	b, err := io.ReadAll(a.resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, &actual); err != nil {
		return err
	}
	return assertExpectedAndActual(
		assert.Equal, expected, actual,
		"Expected %d JSON, but got is %d", expected, actual,
	)
}

func TestOrderFeaturesp(t *testing.T) {
	api := &apiOrderFeature{}
	suite := godog.TestSuite{
		ScenarioInitializer: func(sc *godog.ScenarioContext) {
			sc.Before(func(ctx context.Context, s *godog.Scenario) (context.Context, error) {
				api.fr = router.NewFiberRouter()
				api.orders = []models.Order{}
				return ctx, nil
			})
			sc.Given(`^the following orders$`, api.theFollowingOrders)
			sc.When(`^I send "(GET|POST)" request to "([^"]*)" with body:$`, api.iSendRequestToWithBody)
			sc.Then(`^the response code should be (\d+)$`, api.theResponseCodeShouldBe)
			sc.Then(`^the response should match json:$`, api.theResponseShouldMatchJson)
		},
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
		},
	}

	if execFunctional := os.Getenv("EXEC_FUNCTIONAL"); execFunctional == "" {
		t.Skip("skipping functional tests")
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func assertExpectedAndActual(a expectedAndActualAssertion, expected, actual interface{}, msgAndArgs ...interface{}) error {
	var t asserter
	a(&t, expected, actual, msgAndArgs...)
	return t.err
}

type expectedAndActualAssertion func(t assert.TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool

type asserter struct {
	err error
}

// Errorf is used by the called assertion to report an error
func (a *asserter) Errorf(format string, args ...interface{}) {
	a.err = fmt.Errorf(format, args...)
}
