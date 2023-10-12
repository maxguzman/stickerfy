package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"stickerfy/app/models"
	"stickerfy/app/services"
	"stickerfy/pkg/metrics"
	"stickerfy/pkg/platform/events"
	"stickerfy/pkg/utils"

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
	orderService  services.OrderService
	eventProducer events.EventProducer
	customMetrics metrics.Metrics
	context       context.Context
}

// NewOrderController creates a new OrderController
func NewOrderController(ctx context.Context, orderService services.OrderService, eventProducer events.EventProducer, customMetrics metrics.Metrics) OrderController {
	return &orderController{
		orderService:  orderService,
		eventProducer: eventProducer,
		customMetrics: customMetrics,
		context:       ctx,
	}
}

// GetAll returns all orders
// w
// @Summary		get all exists orders
// @Description	Get all exists orders.
// @ID				get-all-orders
// @Tags			orders
// @Accept			json
// @Produce		json
// @Success		200	{object}	utils.HTTPOrders
// @Failure		400	{object}	utils.HTTPError
// @Failure		404	{object}	utils.HTTPError
// @Failure		500	{object}	utils.HTTPError
// @Router			/orders [get]
func (oc *orderController) GetAll(c *fiber.Ctx) error {
	orders, err := oc.orderService.GetAll(oc.context)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "there was an error getting the orders",
		})
	}

	if orders == nil {
		return c.Status(http.StatusNotFound).JSON(utils.HTTPError{
			Code:    http.StatusNotFound,
			Message: "there are no orders",
		})
	}

	return c.Status(http.StatusOK).JSON(utils.HTTPOrders{
		Orders: orders,
	})
}

// Post creates a new order
//
// @Summary		create a new order
// @Description	Create a new order.
// @ID				create-new-order
// @Tags			orders
// @Accept			json
// @Produce		json
// @Param			order	body		models.Order	true	"Order"
// @Success		201		{object}	models.Order
// @Failure		400		{object}	utils.HTTPError
// @Failure		500		{object}	utils.HTTPError
// @Router			/orders [post]
func (oc *orderController) Post(c *fiber.Ctx) error {
	var order models.Order
	order.ID = uuid.New()
	if err := c.BodyParser(&order); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid order",
		})
	}
	err := oc.orderService.Post(oc.context, order)
	if err != nil {
		oc.customMetrics.IncrementCounter("orderFailed", order.ID.String())
		return c.Status(http.StatusInternalServerError).JSON(utils.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "there was an error creating the order",
		})
	}
	encodedOrder, err := json.Marshal(order)
	if err != nil {
		oc.customMetrics.IncrementCounter("orderFailed", order.ID.String())
		return c.Status(http.StatusInternalServerError).JSON(utils.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "there was an error encoding the order",
		})
	}
	if err := oc.eventProducer.Publish(encodedOrder); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "there was an error publishing the order",
		})
	}
	oc.customMetrics.IncrementCounter("orderAdded", order.ID.String())
	return c.Status(http.StatusCreated).JSON(order)
}
