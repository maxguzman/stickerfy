package repositories

import (
	"context"
	"stickerfy/app/models"

	"stickerfy/pkg/configs"
	"stickerfy/pkg/platform/database"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

const productsCollection = "products"

// ProductRepository is an interface for a product repository
type ProductRepository interface {
	GetAll() ([]models.Product, error)
	GetByID(id uuid.UUID) (models.Product, error)
	Post(product models.Product) error
	Update(product models.Product) error
	Delete(product models.Product) error
}

// productRepository is a implementation of ProductRepository
type productRepository struct {
	client database.Client
}

// NewProductRepository creates a new ProductRepository
func NewProductRepository(c database.Client) ProductRepository {
	return &productRepository{
		client: c,
	}
}

// FindAll returns all products
func (pr *productRepository) GetAll() ([]models.Product, error) {
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
func (pr *productRepository) GetByID(id uuid.UUID) (models.Product, error) {
	var product models.Product
	col := pr.getCollection()
	err := col.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&product)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

// New creates a new product
func (pr *productRepository) Post(product models.Product) error {
	col := pr.getCollection()
	_, err := col.InsertOne(context.TODO(), product)
	if err != nil {
		return err
	}
	return nil
}

// Update updates a product
func (pr *productRepository) Update(product models.Product) error {
	col := pr.getCollection()
	_, err := col.UpdateOne(context.TODO(), bson.M{"_id": product.ID}, bson.M{"$set": product})
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a product by id
func (pr *productRepository) Delete(product models.Product) error {
	col := pr.getCollection()
	_, err := col.DeleteOne(context.TODO(), bson.M{"_id": product.ID})
	if err != nil {
		return err
	}
	return nil
}

func (pr *productRepository) getCollection()database.Collection {
	mc := configs.NewMongoConfig()
	return pr.client.Database(mc.GetDatabase()).Collection(productsCollection)
}
