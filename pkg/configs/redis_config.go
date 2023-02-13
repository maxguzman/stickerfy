package configs

import (
	"os"

	"github.com/go-redis/redis/v8"

	"stickerfy/pkg/utils"
)

// RedisConfig returns a redis config
func RedisConfig() *redis.Options {
	redisAddr, _ := utils.URLBuilder("redis")
	return &redis.Options{
		Addr:     redisAddr,
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	}
}
