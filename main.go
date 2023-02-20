package main

import (
	"context"
	"os"
	"stickerfy/app/repositories"
	"stickerfy/app/services"
	"stickerfy/pkg/controllers"
	"stickerfy/pkg/middleware"
	"stickerfy/pkg/router"
	"stickerfy/pkg/routes"

	_ "stickerfy/docs"
)

const (
	productsCollection string = "products"
	ordersCollection   string = "orders"
)

var (
	productRepository repositories.ProductRepository = repositories.NewMongoProductRepository(context.Background(), productsCollection)
	orderRepository   repositories.OrderRepository   = repositories.NewMongoOrderRepository(context.Background(), ordersCollection)
	productService    services.ProductService        = services.NewProductService(productRepository)
	orderService      services.OrderService          = services.NewOrderService(orderRepository)
	productController controllers.ProductController  = controllers.NewProductController(productService)
	orderController   controllers.OrderController    = controllers.NewOrderController(orderService)
	httpRouter        router.Router                  = router.NewFiberRouter()
)

// @title Stickerfy API
// @version 1.0
// @description A fun sticker store REST API
// @termsOfService http://swagger.io/terms/
// @contact.name Max Guzman
// @contact.email max.guzman@icloud.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	middleware.FiberMiddleware(httpRouter)
	routes.SwaggerRoute(httpRouter)
	routes.ProductRoutes(httpRouter, productController)
	routes.OrderRoutes(httpRouter, orderController)
	routes.NotFoundRoute(httpRouter)

	if os.Getenv("ENV") == "production" {
		httpRouter.ServeWithGracefulShutdown()
	} else {
		httpRouter.Serve()
	}
}
