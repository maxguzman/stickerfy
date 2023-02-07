package router

import (
	"fmt"
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

type fiberRouter struct {
	app *fiber.App
}

// NewFiberRouter creates a new fiber router
func NewFiberRouter() Router {
	return &fiberRouter{app: fiber.New()}
}

func (fr *fiberRouter) Get(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	fr.app.Get(uri, adaptor.HTTPHandlerFunc(f))
}

func (fr *fiberRouter) Post(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	fr.app.Post(uri, adaptor.HTTPHandlerFunc(f))
}

func (fr *fiberRouter) Put(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	fr.app.Put(uri, adaptor.HTTPHandlerFunc(f))
}

func (fr *fiberRouter) Delete(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	fr.app.Delete(uri, adaptor.HTTPHandlerFunc(f))
}

func (fr *fiberRouter) Serve(port string) {
	fmt.Printf("Fiber server running on port %s", port)
	fr.app.Listen(port)
}
