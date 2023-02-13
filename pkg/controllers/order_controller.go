package controllers

import (
	"net/http"
	"stickerfy/app/models"
	"stickerfy/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// OrderController is an interface for an order controller
type OrderController interface {
	GetAll(*fiber.Ctx) error
	Post(*fiber.Ctx) error
}

// orderController is a implementation of OrderController
type orderController struct {
	orderService services.OrderService
}

// NewOrderController creates a new OrderController
func NewOrderController(orderService services.OrderService) OrderController {
	return &orderController{
		orderService: orderService,
	}
}

// GetAll returns all orders
func (oc *orderController) GetAll(c *fiber.Ctx) error {
	orders, err := oc.orderService.GetAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"orders": nil,
			"error":  true,
			"msg":    "there where no orders found",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"orders": orders,
		"error":  false,
		"msg":    nil,
	})
}

// Post creates a new order
func (oc *orderController) Post(c *fiber.Ctx) error {
	var order models.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"order": nil,
			"error": true,
			"msg":   "invalid order",
		})
	}
	order.ID = uuid.New()
	err := oc.orderService.Post(order)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"order": nil,
			"error": true,
			"msg":   "there was an error creating the order",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"order": order,
		"error": false,
		"msg":   nil,
	})
}
