package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds configuration values for the application
type Config struct {
	HOST            string
	PORT            string
	DB_USER         string
	DB_PASSWORD     string
	DB_NAME         string
	SECRET_KEY      string
	EXPIRATION_TIME string
}

var Envs = initConfig()

// InitConfig initializes and returns the application configuration
func initConfig() Config {
	// Load .env file, log a fatal error if it fails
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found or failed to load. Using default environment variables.")
	}

	return Config{
		HOST:            getEnv("DB_HOST", "localhost"),
		PORT:            getEnv("DB_PORT", "3306"),
		DB_USER:         getEnv("DB_USER", "root"),
		DB_PASSWORD:     getEnv("DB_PASSWORD", ""),
		DB_NAME:         getEnv("DB_NAME", "testdb"),
		SECRET_KEY:      getEnv("SECRET_KEY", ""),
		EXPIRATION_TIME: getEnv("EXPIRATION_TIME", "1"),
	}
}

// getEnv retrieves an environment variable or returns the fallback string
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
