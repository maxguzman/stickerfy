package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client is an interface for database clients
type Client interface {
	Close(context.Context) error
	Ping(context.Context) error
	Database(name string, opts ...*options.DatabaseOptions) Database
}

// Database is an interface for databases
type Database interface {
	Collection(name string) Collection
}

// Collection is an interface for collections
type Collection interface {
	Find(context.Context, interface{}) (Cursor, error)
	FindOne(context.Context, interface{}) SingleResult
	InsertOne(context.Context, interface{}) (interface{}, error)
	UpdateOne(context.Context, interface{}, interface{}) (interface{}, error)
	DeleteOne(context.Context, interface{}) (interface{}, error)
}

// SingleResult is an interface for single results
type SingleResult interface {
	Decode(interface{}) error
}

// Cursor is an interface for cursors
type Cursor interface {
	Next(context.Context) bool
	Decode(interface{}) error
	Close(context.Context) error
}
