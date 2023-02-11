package middleware

import (
	"stickerfy/pkg/router"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// FiberMiddleware provide Fiber's built-in middlewares.
func FiberMiddleware(httpRouter router.Router) {
	httpRouter.Use(logger.New()).
		Use(cors.New())
}
