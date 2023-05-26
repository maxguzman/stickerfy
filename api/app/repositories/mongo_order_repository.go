package repositories

import (
	"context"
	"stickerfy/app/models"
	"stickerfy/pkg/configs"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mongoOrderRepository is a implementation of OrderRepository
type mongoOrderRepository struct {
	*mongo.Client
	collection string
}

// NewMongoOrderRepository creates a new OrderRepository
func NewMongoOrderRepository(ctx context.Context, collection string) OrderRepository {
	uri := configs.NewMongoConfig().GetURI()
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	return &mongoOrderRepository{client, collection}
}

// GetAll returns all orders
func (or *mongoOrderRepository) GetAll(ctx context.Context) ([]models.Order, error) {
	var orders []models.Order
	col := or.getCollection()
	cursor, err := col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
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
func (or *mongoOrderRepository) Post(ctx context.Context, order models.Order) error {
	col := or.getCollection()
	_, err := col.InsertOne(ctx, order)
	if err != nil {
		return err
	}
	return nil
}

func (or *mongoOrderRepository) getCollection() *mongo.Collection {
	return or.Database(configs.NewMongoConfig().GetDatabase()).Collection(or.collection)
}
