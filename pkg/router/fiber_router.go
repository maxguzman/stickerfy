package router

import (
	"net/http"
	"stickerfy/pkg/configs"
	"stickerfy/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

// fiberRouter is an implementation of Router interface
type fiberRouter struct {
	app *fiber.App
}

// NewFiberRouter creates a new fiber router
func NewFiberRouter() Router {
	return &fiberRouter{app: fiber.New(configs.FiberConfig())}
}

// Get is a method for GET HTTP method
func (fr *fiberRouter) Get(path string, f func(*fiber.Ctx) error) fiber.Router {
	return fr.app.Get(path, f)
}

// Post is a method for POST HTTP method
func (fr *fiberRouter) Post(path string, f func(*fiber.Ctx) error) fiber.Router {
	return fr.app.Post(path, f)
}

// Put is a method for PUT HTTP method
func (fr *fiberRouter) Put(path string, f func(*fiber.Ctx) error) fiber.Router {
	return fr.app.Put(path, f)
}

// Delete is a method for DELETE HTTP method
func (fr *fiberRouter) Delete(path string, f func(*fiber.Ctx) error) fiber.Router {
	return fr.app.Delete(path, f)
}

// Group is a method for grouping routes
func (fr *fiberRouter) Group(prefix string, f func(c *fiber.Ctx) error) fiber.Router {
	return fr.app.Group(prefix, f)
}

// Use is a method for adding middleware to the router
func (fr *fiberRouter) Use(f func(c *fiber.Ctx) error) fiber.Router {
	return fr.app.Use(f)
}

// Serve is a method for running the server
func (fr *fiberRouter) Serve() {
	addr, err := utils.URLBuilder("server")
	if err != nil {
		panic(err)
	}
	if err := fr.app.Listen(addr); err != nil {
		panic(err)
	}
}

// ServeWithGracefulShutdown is a method for running the server with graceful shutdown
func (fr *fiberRouter) ServeWithGracefulShutdown() {
	addr, err := utils.URLBuilder("server")
	if err != nil {
		panic(err)
	}
	if err := fr.app.Listen(addr); err != nil {
		panic(err)
	}
}

// Test is a method for testing the server
func (fr *fiberRouter) Test(req *http.Request, msTimeout ...int) (*http.Response, error) {
	return fr.app.Test(req, msTimeout...)
}
