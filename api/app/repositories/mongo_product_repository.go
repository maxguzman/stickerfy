package repositories

import (
	"context"
	"stickerfy/app/models"

	"stickerfy/pkg/configs"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mongoProductRepository is a implementation of ProductRepository
type mongoProductRepository struct {
	*mongo.Client
	collection string
}

// NewMongoProductRepository creates a new ProductRepository
func NewMongoProductRepository(ctx context.Context, collection string) ProductRepository {
	uri := configs.NewMongoConfig().GetURI()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return &mongoProductRepository{client, collection}
}

// FindAll returns all products
func (pr *mongoProductRepository) GetAll(ctx context.Context) ([]models.Product, error) {
	var products []models.Product
	col := pr.getCollection()
	cursor, err := col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product models.Product
		err := cursor.Decode(&product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

// Get returns a product by id
func (pr *mongoProductRepository) GetByID(ctx context.Context, id uuid.UUID) (models.Product, error) {
	var product models.Product
	col := pr.getCollection()
	err := col.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

// New creates a new product
func (pr *mongoProductRepository) Post(ctx context.Context, product models.Product) error {
	col := pr.getCollection()
	_, err := col.InsertOne(ctx, product)
	if err != nil {
		return err
	}
	return nil
}

// Update updates a product
func (pr *mongoProductRepository) Update(ctx context.Context, product models.Product) error {
	col := pr.getCollection()
	_, err := col.UpdateOne(ctx, bson.M{"_id": product.ID}, bson.M{"$set": product})
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a product by id
func (pr *mongoProductRepository) Delete(ctx context.Context, product models.Product) error {
	col := pr.getCollection()
	_, err := col.DeleteOne(ctx, bson.M{"_id": product.ID})
	if err != nil {
		return err
	}
	return nil
}

func (pr *mongoProductRepository) getCollection() *mongo.Collection {
	return pr.Database(configs.NewMongoConfig().Database).Collection(pr.collection)
}
