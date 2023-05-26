package cache

import (
	"context"
	"time"
)

// Cache is an interface for cache implementations
type Cache interface {
	Set(context.Context, string, interface{}, time.Duration) error
	Get(context.Context, string) (string, error)
	Ping(context.Context) (string, error)
}
