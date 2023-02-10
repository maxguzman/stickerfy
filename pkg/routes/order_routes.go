package routes

import (
	"stickerfy/pkg/controllers"
	"stickerfy/pkg/router"
)

// OrderRoutes registers all public routes
func OrderRoutes(httpRouter router.Router, orderController controllers.OrderController) {
	httpRouter.Get("/orders", orderController.GetAll)
	httpRouter.Post("/order", orderController.Post)
}
