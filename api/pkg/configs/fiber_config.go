package configs

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// FiberConfig returns the fiber config
func FiberConfig() fiber.Config {
	writeTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("WRITE_TIMEOUT"))
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("READ_TIMEOUT"))
	idleTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("IDLE_TIMEOUT"))

	return fiber.Config{
		WriteTimeout: time.Second * time.Duration(writeTimeoutSecondsCount),
		ReadTimeout:  time.Second * time.Duration(readTimeoutSecondsCount),
		IdleTimeout:  time.Second * time.Duration(idleTimeoutSecondsCount),
	}
}
