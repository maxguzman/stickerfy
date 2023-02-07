package routes

import (
	"stickerfy/pkg/controllers"
	"stickerfy/pkg/router"
)

// PublicRoutes registers all public routes
func PublicRoutes(router router.Router, productController controllers.ProductController, orderController controllers.OrderController) {
	router.Get("/products", productController.GetAll)
	router.Get("/products/{id}", productController.GetByID)
	router.Post("/products", productController.Post)
	router.Put("/products/{id}", productController.Update)
	router.Delete("/products/{id}", productController.Delete)

	router.Get("/orders", orderController.GetAll)
	router.Post("/orders", orderController.Post)
}
