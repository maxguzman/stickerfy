package database

import (
	"context"
	"stickerfy/pkg/configs"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mongoClient is a wrapper around the mongo.Client
type mongoClient struct {
	*mongo.Client
}

// NewMongoClient creates a new mongo client
func NewMongoClient(ctx context.Context) Client {
	uri := configs.NewMongoConfig().GetURI()
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	return &mongoClient{client}
}

// Close closes the mongo client
func (c *mongoClient) Close(ctx context.Context) error {
	return c.Client.Disconnect(ctx)
}

// Ping checks if the mongo client is alive
func (c *mongoClient) Ping(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return c.Client.Ping(ctx, nil)
}

// Database returns a database
func (c *mongoClient) Database(name string, opts ...*options.DatabaseOptions) Database {
	return &mongoDatabase{c.Client.Database(name, opts...)}
}

type mongoDatabase struct {
	*mongo.Database
}

// Collection returns a collection
func (d *mongoDatabase) Collection(name string) Collection {
	return &mongoCollection{d.Database.Collection(name)}
}

type mongoCollection struct {
	*mongo.Collection
}

// Find returns a cursor
func (c *mongoCollection) Find(ctx context.Context, filter interface{}) (Cursor, error) {
	return c.Collection.Find(ctx, filter)
}

// FindOne returns a single result
func (c *mongoCollection) FindOne(ctx context.Context, filter interface{}) SingleResult {
	return c.Collection.FindOne(ctx, filter)
}

// InsertOne inserts a document
func (c *mongoCollection) InsertOne(ctx context.Context, document interface{}) (interface{}, error) {
	return c.Collection.InsertOne(ctx, document)
}

// UpdateOne updates a document
func (c *mongoCollection) UpdateOne(ctx context.Context, filter, update interface{}) (interface{}, error) {
	return c.Collection.UpdateOne(ctx, filter, update)
}

// DeleteOne deletes a document
func (c *mongoCollection) DeleteOne(ctx context.Context, filter interface{}) (interface{}, error) {
	return c.Collection.DeleteOne(ctx, filter)
}
