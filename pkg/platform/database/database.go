package database

// Client is an interface for database clients
type Client interface {
	Connect() error
	Disconnect() error
}
