package controllers

import (
	"encoding/json"
	"net/http"
	"stickerfy/app/models"
	"stickerfy/app/services"
)

// OrderController is an interface for an order controller
type OrderController interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
}

// NewOrderController creates a new OrderController
func NewOrderController(orderService services.OrderService) OrderController {
	return &orderController{
		orderService: orderService,
	}
}

// orderController is a implementation of OrderController
type orderController struct {
	orderService services.OrderService
}

// GetAll returns all orders
func (oc *orderController) GetAll(w http.ResponseWriter, r *http.Request) {
	orders, err := oc.orderService.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(orders); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// Post creates a new order
func (oc *orderController) Post(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := oc.orderService.Post(order)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
