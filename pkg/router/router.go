package router

import "github.com/gofiber/fiber/v2"

// Router is an interface for a router
type Router interface {
	Get(path string, f func(*fiber.Ctx) error)
	Post(path string, f func(*fiber.Ctx) error)
	Put(path string, f func(*fiber.Ctx) error)
	Delete(path string, f func(*fiber.Ctx) error)
	Serve(port string)
}
