package utils

import "stickerfy/app/models"

// HTTPError is a struct for http errors
type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// HTTPProducts is a struct for http products responses
type HTTPProducts struct {
	Products []models.Product `json:"products"`
}

// HTTPOrders is a struct for http orders responses
type HTTPOrders struct {
	Orders []models.Order `json:"orders"`
}
