package routes

import (
	"stickerfy/pkg/controllers"
	"stickerfy/pkg/router"
)

// ProductRoutes registers all public routes
func ProductRoutes(httpRouter router.Router, productController controllers.ProductController) {
	route := httpRouter.Group("/v1")

	route.Get("/products", productController.GetAll)
	route.Get("/product/:id", productController.GetByID)
	route.Post("/product", productController.Post)
	route.Put("/product", productController.Update)
	route.Delete("/product", productController.Delete)
}
