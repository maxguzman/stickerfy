package router

import (
	"fmt"
	"net/http"
	"stickerfy/pkg/configs"
	"stickerfy/pkg/utils"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/mux"
)

type muxRouter struct {
	mux *mux.Router
}

// NewMuxRouter creates a new mux router
func NewMuxRouter() Router {
	return &muxRouter{mux: mux.NewRouter()}
}

// Get is a function that adds a new GET route
func (mr *muxRouter) Get(path string, f func(*fiber.Ctx) error) fiber.Router {
	mr.mux.HandleFunc(path, adaptor.FiberHandlerFunc(f)).Methods(http.MethodGet)
	return nil
}

// Post is a function that adds a new POST route
func (mr *muxRouter) Post(path string, f func(*fiber.Ctx) error) fiber.Router {
	mr.mux.HandleFunc(path, adaptor.FiberHandlerFunc(f)).Methods(http.MethodPost)
	return nil
}

// Put is a function that adds a new PUT route
func (mr *muxRouter) Put(path string, f func(*fiber.Ctx) error) fiber.Router {
	mr.mux.HandleFunc(path, adaptor.FiberHandlerFunc(f)).Methods(http.MethodPut)
	return nil
}

// Delete is a function that adds a new DELETE route
func (mr *muxRouter) Delete(path string, f func(*fiber.Ctx) error) fiber.Router {
	mr.mux.HandleFunc(path, adaptor.FiberHandlerFunc(f)).Methods(http.MethodDelete)
	return nil
}

// TODO: Implement Group and Use
// Use is a function that adds middleware to the router
func (mr *muxRouter) Use(f func(c *fiber.Ctx) error) fiber.Router {
	return nil
}

// Group is a function that groups routes
func (mr *muxRouter) Group(prefix string, f func(c *fiber.Ctx) error) fiber.Router {
	return nil
}

// Serve is a function that runs the server
func (mr *muxRouter) Serve() {
	url, err := utils.URLBuilder("server")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Mux server is running on: %s", url)
	server := configs.MuxConfig(mr.mux)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

// TODO: Implement graceful shutdown
// ServeWithGracefulShutdown is a function that runs the server with graceful shutdown
func (mr *muxRouter) ServeWithGracefulShutdown() {
	url, err := utils.URLBuilder("server")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Mux server running on port %s", url)
	server := configs.MuxConfig(mr.mux)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

// Test is a method for testing the server
func (mr *muxRouter) Test(req *http.Request, msTimeout ...int) (*http.Response, error) {
	return nil, nil
}
