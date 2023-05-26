package repositories

import (
	"context"
	"stickerfy/app/models"
)

// OrderRepository is an interface for an order repository
type OrderRepository interface {
	GetAll(ctx context.Context) ([]models.Order, error)
	Post(ctx context.Context, order models.Order) error
}
