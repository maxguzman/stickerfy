package main

import (
	"context"
	"log"
	"os"
	"stickerfy/app/repositories"
	"stickerfy/app/services"
	"stickerfy/pkg/controllers"
	"stickerfy/pkg/metrics"
	"stickerfy/pkg/middleware"
	"stickerfy/pkg/platform/cache"
	"stickerfy/pkg/platform/events"
	"stickerfy/pkg/router"
	"stickerfy/pkg/routes"
	"stickerfy/pkg/utils"

	_ "stickerfy/docs"
)

const (
	productsCollection string = "products"
	ordersCollection   string = "orders"
)

var (
	ctx               context.Context                = context.Background()
	productRepository repositories.ProductRepository = repositories.NewMongoProductRepository(ctx, productsCollection)
	orderRepository   repositories.OrderRepository   = repositories.NewMongoOrderRepository(ctx, ordersCollection)
	productService    services.ProductService        = services.NewProductService(productRepository)
	orderService      services.OrderService          = services.NewOrderService(orderRepository)
	productCache      cache.Cache                    = cache.NewRedisClient()
	productController controllers.ProductController  = controllers.NewProductController(ctx, productService, productCache)
	orderEvents       events.EventProducer           = events.NewKafkaProducer(os.Getenv("TOPIC_NAME"))
	customMetrics     metrics.Metrics                = metrics.NewPrometheusMetrics()
	orderController   controllers.OrderController    = controllers.NewOrderController(ctx, orderService, orderEvents, customMetrics)
	httpRouter        router.Router                  = router.NewFiberRouter()
)

// @title			Stickerfy API
// @version		1.0
// @description	A fun sticker store REST API
// @termsOfService	https://swagger.io/terms/
// @contact.name	Max Guzman
// @contact.email	max.guzman@icloud.com
// @license.name	Apache 2.0
// @license.url	https://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath		/
// @host			localhost:8000
func main() {
	middleware.FiberMiddleware(httpRouter)
	routes.SwaggerRoute(httpRouter)
	routes.ProductRoutes(httpRouter, productController)
	routes.OrderRoutes(httpRouter, orderController)
	routes.MetricsRoute(httpRouter)
	routes.NotFoundRoute(httpRouter)

	tp := utils.InitTracer(ctx)
	defer func() {
		if err := tp.Shutdown(ctx); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	if os.Getenv("ENV") == "dev" {
		httpRouter.Serve()
	} else {
		httpRouter.ServeWithGracefulShutdown()
	}
}
