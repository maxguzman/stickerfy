package configs

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// FiberConfig returns the fiber config
func FiberConfig() fiber.Config {
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
	}
}
