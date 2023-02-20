package router

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Router is an interface for a router
type Router interface {
	Get(path string, f func(*fiber.Ctx) error) fiber.Router
	Post(path string, f func(*fiber.Ctx) error) fiber.Router
	Put(path string, f func(*fiber.Ctx) error) fiber.Router
	Delete(path string, f func(*fiber.Ctx) error) fiber.Router
	Group(prefix string, f func(*fiber.Ctx) error) fiber.Router
	Use(f func(*fiber.Ctx) error) fiber.Router
	Serve()
	ServeWithGracefulShutdown()
	Test(req *http.Request, msTimeout ...int) (*http.Response, error)
}
