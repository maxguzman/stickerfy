package router

import "github.com/gofiber/fiber/v2"

// Router is an interface for a router
type Router interface {
	Get(uri string, f func(*fiber.Ctx) error)
	Post(uri string, f func(*fiber.Ctx) error)
	Put(uri string, f func(*fiber.Ctx) error)
	Delete(uri string, f func(*fiber.Ctx) error)
	Serve(port string)
}
