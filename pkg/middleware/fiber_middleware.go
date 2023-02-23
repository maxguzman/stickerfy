package middleware

import (
	"stickerfy/pkg/router"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/mikhail-bigun/fiberlogrus"
	"github.com/sirupsen/logrus"
)

// FiberMiddleware provide Fiber's built-in middlewares.
func FiberMiddleware(httpRouter router.Router) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})
	logger.SetLevel(logrus.InfoLevel)

	httpRouter.
		Use(cors.New()).
		Use(fiberlogrus.New(fiberlogrus.Config{
			Logger: logger,
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
		}))
}
