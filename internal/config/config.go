package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// global configuration variable for the application
var AppConfig *Config

// load configuration from .env file and system environment variables
func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	AppConfig = &Config{
		DB: DBConfig{
			DSN:             getEnvOrDefault("DB_DSN", ""),
			MaxOpenConns:    getEnvAsIntOrDefault("DB_MAX_CONNS", 25),
			MaxIdleConns:    getEnvAsIntOrDefault("DB_MAX_IDLE_CONNS", 25),
			ConnMaxLifetime: getEnvAsDurationOrDefault("DB_CONN_MAX_LIFETIME", 10*time.Minute),
		},
		Server: ServerConfig{
			Port:        getEnvOrDefault("PORT", ""),
			Environment: getEnvOrDefault("ENV", ""),
		},
		Auth: AuthConfig{
			ProxyURL: getEnvOrDefault("AUTH_PROXY_URL", ""),
		},
	}

	validateConfig()
}

// returns the environment variable value or the default value if it is not set.
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// returns the environment variable value as an integer or the default value if it is not set.
func getEnvAsIntOrDefault(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intValue
}

// returns the environment variable value as a duration or the default value if it is not set.
func getEnvAsDurationOrDefault(key string, defaultValue time.Duration) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	duration, err := time.ParseDuration(value)
	if err != nil {
		return defaultValue
	}
	return duration
}

// validates the required configuration values
func validateConfig() {
	// DB configuration validation
	if AppConfig.DB.DSN == "" {
		log.Fatal("DB_DSN is not set in the environment")
	}
	if AppConfig.DB.MaxOpenConns <= 0 {
		log.Fatal("DB_MAX_CONNS must be greater than 0")
	}
	if AppConfig.DB.MaxIdleConns <= 0 {
		log.Fatal("DB_MAX_IDLE_CONNS must be greater than 0")
	}
	if AppConfig.DB.ConnMaxLifetime <= 0 {
		log.Fatal("DB_CONN_MAX_LIFETIME must be greater than 0")
	}

	// server configuration validation
	if AppConfig.Server.Port == "" {
		log.Fatal("PORT is not set in the environment")
	}
	if AppConfig.Server.Environment == "" {
		log.Fatal("ENV is not set in the environment")
	}

	// authentication configuration validation
	if AppConfig.Auth.ProxyURL == "" {
		log.Fatal("AUTH_PROXY_URL is not set in the environment")
	}

	log.Println("Config loaded successfully")

	// test config print
	log.Printf("AppConfig: \n"+
		"  DB: \n"+
		"    DSN: %s\n"+
		"    MaxOpenConns: %d\n"+
		"    MaxIdleConns: %d\n"+
		"    ConnMaxLifetime: %s\n"+
		"  Server: \n"+
		"    Port: %s\n"+
		"    Environment: %s\n"+
		"  Auth: \n"+
		"    ProxyURL: %s\n",
		AppConfig.DB.DSN,
		AppConfig.DB.MaxOpenConns,
		AppConfig.DB.MaxIdleConns,
		AppConfig.DB.ConnMaxLifetime,
		AppConfig.Server.Port,
		AppConfig.Server.Environment,
		AppConfig.Auth.ProxyURL)
}
