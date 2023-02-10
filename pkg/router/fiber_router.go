package router

import (
	"stickerfy/pkg/configs"

	"github.com/gofiber/fiber/v2"
)

type fiberRouter struct {
	app *fiber.App
}

// NewFiberRouter creates a new fiber router
func NewFiberRouter() Router {
	return &fiberRouter{app: fiber.New(configs.FiberConfig())}
}

func (fr *fiberRouter) Get(path string, f func(*fiber.Ctx) error) {
	fr.app.Get(path, f)
}

func (fr *fiberRouter) Post(path string, f func(*fiber.Ctx) error) {
	fr.app.Post(path, f)
}

func (fr *fiberRouter) Put(path string, f func(*fiber.Ctx) error) {
	fr.app.Put(path, f)
}

func (fr *fiberRouter) Delete(path string, f func(*fiber.Ctx) error) {
	fr.app.Delete(path, f)
}

func (fr *fiberRouter) Serve(port string) {
	if err := fr.app.Listen(":" + port); err != nil {
		panic(err)
	}
}
