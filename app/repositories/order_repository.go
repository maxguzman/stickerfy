package repositories

import (
	"stickerfy/app/models"
)

// OrderRepository is an interface for an order repository
type OrderRepository interface {
	GetAll() ([]models.Order, error)
	Post(order models.Order) error
}
