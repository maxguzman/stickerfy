package main

import (
	"stickerfy/app/repositories"
	"stickerfy/app/services"
	"stickerfy/pkg/controllers"
	"stickerfy/pkg/platform/database"
	"stickerfy/pkg/router"
	"stickerfy/pkg/routes"
)

var (
	productRepository repositories.ProductRepository = database.NewMongoProductRepository()
	orderRepository   repositories.OrderRepository   = database.NewMongoOrderRepository()
	productService    services.ProductService        = services.NewProductService(productRepository)
	orderService      services.OrderService          = services.NewOrderService(orderRepository)
	productController controllers.ProductController  = controllers.NewProductController(productService)
	orderController   controllers.OrderController    = controllers.NewOrderController(orderService)
	httpRouter        router.Router                  = router.NewFiberRouter()
)

func main() {
	routes.ProductRoutes(httpRouter, productController)
	routes.OrderRoutes(httpRouter, orderController)
	httpRouter.Serve()
}
