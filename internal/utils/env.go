package utils

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func MustGetEnv(name string) string {
	value := os.Getenv(name)
	if value == "" {
		panic("missing required environment variable " + name)
	}
	return value
}
