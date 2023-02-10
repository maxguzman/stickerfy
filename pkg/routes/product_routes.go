package routes

import (
	"stickerfy/pkg/controllers"
	"stickerfy/pkg/router"
)

// ProductRoutes registers all public routes
func ProductRoutes(httpRouter router.Router, productController controllers.ProductController) {
	httpRouter.Get("/products", productController.GetAll)
	httpRouter.Get("/product/:id", productController.GetByID)
	httpRouter.Post("/product", productController.Post)
	httpRouter.Put("/product", productController.Update)
	httpRouter.Delete("/product", productController.Delete)
}
