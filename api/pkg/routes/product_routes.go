package routes

import (
	"stickerfy/pkg/controllers"
	"stickerfy/pkg/router"
)

// ProductRoutes registers all public routes
func ProductRoutes(httpRouter router.Router, productController controllers.ProductController) {
	httpRouter.Get("/products", productController.GetAll)
	httpRouter.Get("/products/:id", productController.GetByID)
	httpRouter.Post("/products", productController.Post)
	httpRouter.Put("/products", productController.Update)
	httpRouter.Delete("/products", productController.Delete)
}
