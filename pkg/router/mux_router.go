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

func (mr *muxRouter) Get(path string, f func(*fiber.Ctx) error) {
	mr.mux.HandleFunc(path, adaptor.FiberHandlerFunc(f)).Methods(http.MethodGet)
}

func (mr *muxRouter) Post(path string, f func(*fiber.Ctx) error) {
	mr.mux.HandleFunc(path, adaptor.FiberHandlerFunc(f)).Methods(http.MethodPost)
}

func (mr *muxRouter) Put(path string, f func(*fiber.Ctx) error) {
	mr.mux.HandleFunc(path, adaptor.FiberHandlerFunc(f)).Methods(http.MethodPut)
}

func (mr *muxRouter) Delete(path string, f func(*fiber.Ctx) error) {
	mr.mux.HandleFunc(path, adaptor.FiberHandlerFunc(f)).Methods(http.MethodDelete)
}

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

func (mr *muxRouter) Use(f func(c *fiber.Ctx) error) {}
