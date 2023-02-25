package middleware

import (
	"os"
	"stickerfy/pkg/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mikhail-bigun/fiberlogrus"
	"github.com/sirupsen/logrus"
)

// FiberMiddleware provide Fiber's built-in middlewares.
func FiberMiddleware(httpRouter router.Router) {
	httpRouter.
		Use(cors.New()).
		Use(setLogger())
}

func setLogger() func(ctx *fiber.Ctx) error {
	if os.Getenv("ENV") == "dev" {
		return logger.New()
	}
	logrusLogger := logrus.New()
	logrusLogger.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})
	logrusLogger.SetLevel(logrus.InfoLevel)
	return fiberlogrus.New(fiberlogrus.Config{
		Logger: logrusLogger,
		Tags: []string{
			fiberlogrus.TagStatus,
			fiberlogrus.TagLatency,
			fiberlogrus.TagMethod,
			fiberlogrus.TagIP,
			fiberlogrus.TagPath,
			fiberlogrus.TagBody,
			fiberlogrus.TagUA,
			fiberlogrus.TagResBody,
		},
	})
}
