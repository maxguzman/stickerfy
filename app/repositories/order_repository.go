package repositories

import (
	"context"
	"stickerfy/app/models"
	"stickerfy/pkg/configs"
	"stickerfy/pkg/platform/database"

	"go.mongodb.org/mongo-driver/bson"
)

const ordersCollection = "orders"

// OrderRepository is an interface for an order repository
type OrderRepository interface {
	GetAll() ([]models.Order, error)
	Post(order models.Order) error
}

// orderRepository is a implementation of OrderRepository
type orderRepository struct{
	client database.Client
}

// NewOrderRepository creates a new OrderRepository
func NewOrderRepository(c database.Client) OrderRepository {
	return &orderRepository{
		client: c,
	}
}

// GetAll returns all orders
func (or *orderRepository) GetAll() ([]models.Order, error) {
	var orders []models.Order
	col := or.getCollection()
	cursor, err := col.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var order models.Order
		err := cursor.Decode(&order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

// Post creates a new order
func (or *orderRepository) Post(order models.Order) error {
	col := or.getCollection()
	_, err := col.InsertOne(context.TODO(), order)
	if err != nil {
		return err
	}
	return nil
}

func (or *orderRepository) getCollection() database.Collection {
	mc := configs.NewMongoConfig()
	return or.client.Database(mc.GetDatabase()).Collection(ordersCollection)
}
