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

const productsCollection = "products"

// productRepository is a implementation of ProductRepository
type mongoProductRepository struct {
	*mongo.Client
}

// NewProductRepository creates a new ProductRepository
func NewMongoProductRepository(ctx context.Context) ProductRepository {
		uri := configs.NewMongoConfig().GetURI()
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	return &mongoProductRepository{client}
}

// FindAll returns all products
func (pr *mongoProductRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	col := pr.getCollection()
	cursor, err := col.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
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
func (pr *mongoProductRepository) GetByID(id uuid.UUID) (models.Product, error) {
	var product models.Product
	col := pr.getCollection()
	err := col.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&product)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

// New creates a new product
func (pr *mongoProductRepository) Post(product models.Product) error {
	col := pr.getCollection()
	_, err := col.InsertOne(context.TODO(), product)
	if err != nil {
		return err
	}
	return nil
}

// Update updates a product
func (pr *mongoProductRepository) Update(product models.Product) error {
	col := pr.getCollection()
	_, err := col.UpdateOne(context.TODO(), bson.M{"_id": product.ID}, bson.M{"$set": product})
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a product by id
func (pr *mongoProductRepository) Delete(product models.Product) error {
	col := pr.getCollection()
	_, err := col.DeleteOne(context.TODO(), bson.M{"_id": product.ID})
	if err != nil {
		return err
	}
	return nil
}

func (pr *mongoProductRepository) getCollection() *mongo.Collection {
	return pr.Database(configs.NewMongoConfig().Database).Collection(productsCollection)
}