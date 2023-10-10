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
	"strconv"
	"testing"

	"stickerfy/app/models"
	"stickerfy/pkg/controllers"
	"stickerfy/pkg/router"
	"stickerfy/pkg/routes"
	mock_cache "stickerfy/test/mocks/cache"
	mock_services "stickerfy/test/mocks/services"

	"github.com/cucumber/godog"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type apiProductFeature struct {
	fr       router.Router
	resp     *http.Response
	products []models.Product
}

func (a *apiProductFeature) theFollowingProducts(givenProducts *godog.Table) error {
	head := givenProducts.Rows[0].Cells
	for i := 1; i < len(givenProducts.Rows); i++ {
		product := models.Product{}
		for n, cell := range givenProducts.Rows[i].Cells {
			switch head[n].Value {
			case "id":
				id, _ := uuid.Parse(cell.Value)
				product.ID = id
			case "title":
				product.Title = cell.Value
			case "description":
				product.Description = cell.Value
			case "price":
				price, _ := strconv.ParseFloat(cell.Value, 64)
				product.Price = price
			case "image_path":
				product.ImagePath = cell.Value
			default:
				return fmt.Errorf("unexpected column name: %s", head[n].Value)
			}
		}
		a.products = append(a.products, product)
	}
	return nil
}

func (a *apiProductFeature) iSendRequestTo(method, path string) error {
	t := &testing.T{}

	mockProductCache := mock_cache.NewCache(t)
	mockProductCache.On("Get", mock.Anything, mock.Anything).Return(mock.Anything, redis.Nil)
	mockProductCache.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	mockProductService := mock_services.NewProductService(t)
	if len(a.products) > 0 {
		mockProductService.On("GetAll", mock.Anything).Return(a.products, nil)
		mockProductService.On("GetByID", mock.Anything, mock.Anything).Return(a.products[0], nil)
	} else {
		mockProductService.On("GetAll", mock.Anything).Return(nil, nil)
	}
	productController := controllers.NewProductController(context.Background(), mockProductService, mockProductCache)
	routes.ProductRoutes(a.fr, productController)

	req := httptest.NewRequest(method, path, http.NoBody)
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.fr.Test(req)
	if err != nil {
		return err
	}
	a.resp = resp

	mockProductCache.AssertExpectations(t)
	mockProductService.AssertExpectations(t)
	return nil
}

func (a *apiProductFeature) iSendRequestToWithBody(method, path string, body *godog.DocString) error {
	var expected interface{}
	if err := json.Unmarshal([]byte(body.Content), &expected); err != nil {
		return err
	}
	t := &testing.T{}
	mockProductService := mock_services.NewProductService(t)
	mockProductService.On("Post", mock.Anything, mock.Anything).Return(nil)
	mockProductService.On("Update", mock.Anything, mock.Anything).Return(nil)
	mockProductService.On("Delete", mock.Anything, mock.Anything).Return(nil)
	productController := controllers.NewProductController(context.Background(), mockProductService, nil)
	routes.ProductRoutes(a.fr, productController)

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(&expected)
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

	mockProductService.AssertExpectations(t)
	return nil
}

func (a *apiProductFeature) theResponseCodeShouldBe(expectedCode int) error {
	return assertExpectedAndActual(
		assert.Equal, expectedCode, a.resp.StatusCode,
		"Expected %d code, but got is %d", expectedCode, a.resp.StatusCode,
	)
}

func (a *apiProductFeature) theResponseShouldMatchJson(body *godog.DocString) error {
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

func TestProductFeatures(t *testing.T) {
	api := &apiProductFeature{}
	suite := godog.TestSuite{
		ScenarioInitializer: func(sc *godog.ScenarioContext) {
			sc.Before(func(ctx context.Context, s *godog.Scenario) (context.Context, error) {
				api.fr = router.NewFiberRouter()
				api.products = []models.Product{}
				return ctx, nil
			})
			sc.Given(`^the following products$`, api.theFollowingProducts)
			sc.When(`^I send "([^"]*)" request to "([^"]*)"$`, api.iSendRequestTo)
			sc.When(`^I send "(POST|PUT|DELETE)" request to "([^"]*)" with body:$`, api.iSendRequestToWithBody)
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
