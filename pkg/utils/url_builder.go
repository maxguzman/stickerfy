package utils

import (
	"fmt"
	"os"
)

// URLBuilder is a function to build connection url
func URLBuilder(n string) (string, error) {
	switch n {
	case "redis":
		return fmt.Sprintf(
			"%s:%s",
			os.Getenv("REDIS_HOST"),
			os.Getenv("REDIS_PORT"),
		), nil
	case "mongo":
		return fmt.Sprintf(
			"mongodb://%s:%s@%s:%s",
			os.Getenv("MONGO_USER"),
			os.Getenv("MONGO_PASSWORD"),
			os.Getenv("MONGO_HOST"),
			os.Getenv("MONGO_PORT"),
		), nil
	case "server":
		return fmt.Sprintf(
			"%s:%s",
			os.Getenv("SERVER_HOST"),
			os.Getenv("SERVER_PORT"),
		), nil
	default:
		return "", fmt.Errorf("connection name '%v' is not supported", n)
	}
}
