package configs

import (
	"os"

	"github.com/go-redis/redis/v8"
)

// RedisConfig returns a redis config
func RedisConfig() *redis.Options {
	return &redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	}
}
