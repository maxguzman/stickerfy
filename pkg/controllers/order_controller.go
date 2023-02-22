package controllers

import (
	"encoding/json"
	"net/http"
	"stickerfy/app/models"
	"stickerfy/app/services"
	"stickerfy/pkg/platform/events"

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
	eventProducer events.EventProducer
}

// NewOrderController creates a new OrderController
func NewOrderController(orderService services.OrderService, eventProducer events.EventProducer) OrderController {
	return &orderController{
		orderService: orderService,
		eventProducer: eventProducer,
	}
}

// GetAll returns all orders
// @Description Get all exists orders.
// @Summary get all exists orders
// @Tags Order
// @Accept json
// @Produce json
// @Success 200 {array} models.Product
// @Router /orders [get]
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
// @Description Create a new order.
// @Summary create a new order
// @Tags Order
// @Accept json
// @Produce json
// @Param order body models.Order true "Order"
// @Success 200 {object} models.Order
// @Router /order [post]
func (oc *orderController) Post(c *fiber.Ctx) error {
	var order models.Order
	order.ID = uuid.New()
	if err := c.BodyParser(&order); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"order": nil,
			"error": true,
			"msg":   "invalid order",
		})
	}
	err := oc.orderService.Post(order)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"order": nil,
			"error": true,
			"msg":   "there was an error creating the order",
		})
	}
	encodedOrder, err := json.Marshal(order)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"order": nil,
			"error": true,
			"msg":   "there was an error encoding the order",
		})
	}
	if err:=oc.eventProducer.Publish(encodedOrder); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"order": nil,
			"error": true,
			"msg":   "there was an error publishing the order",
		})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"order": order,
		"error": false,
		"msg":   nil,
	})
}
