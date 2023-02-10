package routes

import (
	"stickerfy/pkg/router"

	"github.com/gofiber/swagger"
)

// SwaggerRoute func for describe group of API Docs routes.
func SwaggerRoute(httpRouter router.Router) {
	httpRouter.Get("/swagger/*", swagger.HandlerDefault) // get one user by ID
}
