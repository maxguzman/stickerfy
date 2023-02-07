package routes

import (
	"stickerfy/pkg/controllers"
	"stickerfy/pkg/router"
)

// PublicRoutes registers all public routes
func PublicRoutes(httpRouter router.Router, productController controllers.ProductController, orderController controllers.OrderController) {
	httpRouter.Get("/products", productController.GetAll)
	httpRouter.Get("/products/{id}", productController.GetByID)
	httpRouter.Post("/products", productController.Post)
	httpRouter.Put("/products/{id}", productController.Update)
	httpRouter.Delete("/products/{id}", productController.Delete)

	httpRouter.Get("/orders", orderController.GetAll)
	httpRouter.Post("/orders", orderController.Post)
}
