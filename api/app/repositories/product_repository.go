package repositories

import (
	"context"
	"stickerfy/app/models"

	"github.com/google/uuid"
)

// ProductRepository is an interface for a product repository
type ProductRepository interface {
	GetAll(ctx context.Context) ([]models.Product, error)
	GetByID(ctx context.Context, id uuid.UUID) (models.Product, error)
	Post(ctx context.Context, product models.Product) error
	Update(ctx context.Context, product models.Product) error
	Delete(ctx context.Context, product models.Product) error
}
