package routes

import (
	"stickerfy/pkg/controllers"
	"stickerfy/pkg/router"
)

// OrderRoutes registers all public routes
func OrderRoutes(httpRouter router.Router, orderController controllers.OrderController) {
	route := httpRouter.Group("/v1")

	route.Get("/orders", orderController.GetAll)
	route.Post("/order", orderController.Post)
}
