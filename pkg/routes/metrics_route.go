package routes

import (
	"stickerfy/pkg/router"

	"github.com/gofiber/adaptor/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// MetricsRoute func for describe group of metrics routes.
func MetricsRoute(httpRouter router.Router) {
	httpRouter.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))
}
