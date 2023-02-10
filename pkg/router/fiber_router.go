package router

import (
	"fmt"
	"net/http"
	"stickerfy/pkg/configs"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

type fiberRouter struct {
	app *fiber.App
}

// NewFiberRouter creates a new fiber router
func NewFiberRouter() Router {
	return &fiberRouter{app: fiber.New(configs.FiberConfig())}
}

func (fr *fiberRouter) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	fr.app.Get(path, adaptor.HTTPHandlerFunc(f))
}

func (fr *fiberRouter) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	fr.app.Post(path, adaptor.HTTPHandlerFunc(f))
}

func (fr *fiberRouter) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	fr.app.Put(path, adaptor.HTTPHandlerFunc(f))
}

func (fr *fiberRouter) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	fr.app.Delete(path, adaptor.HTTPHandlerFunc(f))
}

func (fr *fiberRouter) Serve(port string) {
	fmt.Printf("Fiber server running on port %s", port)
	if err := fr.app.Listen(":" + port); err != nil {
		panic(err)
	}
}
