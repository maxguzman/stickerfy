package configs

import (
	"os"

	"github.com/go-redis/redis/v8"

	"stickerfy/pkg/utils"
)

// RedisConfig returns a redis config
func RedisConfig() *redis.Options {
	redisAddr, err := utils.URLBuilder("redis")
	if err != nil {
		panic(err)
	}
	return &redis.Options{
		Addr:     redisAddr,
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	}
}
