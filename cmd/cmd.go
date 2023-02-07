package main

import (
	"stickerfy/app/repositories"
	"stickerfy/app/services"
	"stickerfy/pkg/controllers"
	"stickerfy/pkg/router"
	"stickerfy/pkg/routes"
	"stickerfy/platform/database"
)

var (
	productRepository repositories.ProductRepository = database.NewMongoProductRepository()
	orderRepository   repositories.OrderRepository   = database.NewMongoOrderRepository()
	productService    services.ProductService        = services.NewProductService(productRepository)
	orderService      services.OrderService          = services.NewOrderService(orderRepository)
	productController controllers.ProductController  = controllers.NewProductController(productService)
	orderController   controllers.OrderController    = controllers.NewOrderController(orderService)
	httpRouter        router.Router                  = router.NewMuxRouter()
)

func main() {
	const port string = ":8080"

	routes.PublicRoutes(httpRouter, productController, orderController)

	httpRouter.Serve(port)
}
