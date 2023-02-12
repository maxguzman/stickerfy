package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"stickerfy/pkg/configs"
)

// RedisClient is an implementation of the Cache interface
type RedisClient struct {
	Client *redis.Client
}

// NewRedisClient instantiates the Redis client
func NewRedisClient() Cache {
	return &RedisClient{
		Client: redis.NewClient(configs.RedisConfig()),
	}
}

// Set a key-value pair in Redis
func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.Client.Set(ctx, key, value, expiration).Err()
}

// Get gets a value from Redis
func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

// Ping pings Redis
func (r *RedisClient) Ping(ctx context.Context) (string, error) {
	return r.Client.Ping(ctx).Result()
}