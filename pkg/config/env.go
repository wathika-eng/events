package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds configuration values for the application
type Config struct {
	Host            string
	Port            int
	DBUser          string
	DBPassword      string
	DBName          string
	SECRET_KEY      string
	EXPIRATION_TIME int
}

var Envs = InitConfig()

// InitConfig initializes and returns the application configuration
func InitConfig() Config {
	// Load .env file, log a fatal error if it fails
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found or failed to load. Using default environment variables.")
	}

	return Config{
		Host:            getEnv("DB_HOST", "localhost"),
		Port:            getEnvAsInt("DB_PORT", 3306),
		DBUser:          getEnv("DB_USER", "root"),
		DBPassword:      getEnv("DB_PASSWORD", ""),
		DBName:          getEnv("DB_NAME", "testdb"),
		SECRET_KEY:      getEnv("SECRET_KEY", ""),
		EXPIRATION_TIME: getEnvAsInt("EXPIRATION_TIME", 1),
	}
}

// getEnv retrieves an environment variable or returns the fallback string
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// getEnvAsInt retrieves an environment variable as an integer or returns the fallback value
func getEnvAsInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
		log.Printf("Warning: Invalid integer for %s. Using fallback value %d\n", key, fallback)
	}
	return fallback
}
