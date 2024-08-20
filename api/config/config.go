package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var ServerAddress string

func LoadConfig() {
    absPath, err := filepath.Abs(filepath.Join("..", "..", ".env"))
    if err != nil {
        log.Fatalf("Error getting absolute path: %v", err)
    }

    err = godotenv.Load(absPath)
    if err != nil {
        log.Fatalf("Error loading .env file from %s", absPath)
    }

    ServerAddress = getEnv("SERVER_ADDRESS", ":3001")
}

func getEnv(key, defaultValue string) string {
    value, exists := os.LookupEnv(key)
    if !exists {
        return defaultValue
    }
    return value
}
