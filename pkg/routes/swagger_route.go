package routes

import (
	"stickerfy/pkg/router"

	"github.com/gofiber/swagger"
)

// SwaggerRoute func for describe group of API Docs routes.
func SwaggerRoute(httpRouter router.Router) {
	route := httpRouter.Group("/swagger")

	route.Get("*", swagger.HandlerDefault)
}
