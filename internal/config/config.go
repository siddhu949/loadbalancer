package config

import (
	"log"
)

// Config structure
type Config struct {
	ServerPort string
	LogLevel   string
}

// DefaultConfig holds the current configuration
var DefaultConfig = Config{
	ServerPort: "8080",
	LogLevel:   "INFO",
}

// LoadConfig loads configuration from file or environment
func LoadConfig() {
	log.Println("Loading configuration...")
	// TODO: Load config from file or environment variables
}
