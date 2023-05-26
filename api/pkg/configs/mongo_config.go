package configs

import (
	"os"
	"stickerfy/pkg/utils"
)

// MongoConfig is a configuration for a mongo client
type MongoConfig struct {
	URI        string
	Database   string
}

// NewMongoConfig creates a new mongo config
func NewMongoConfig() *MongoConfig {
	uri, _ := utils.URLBuilder("mongo")
	return &MongoConfig{
		URI:      uri,
		Database: os.Getenv("MONGO_DATABASE"),
	}
}

// GetURI returns the mongo uri
func (c *MongoConfig) GetURI() string {
	return c.URI
}

// GetDatabase returns the mongo database
func (c *MongoConfig) GetDatabase() string {
	return c.Database
}
