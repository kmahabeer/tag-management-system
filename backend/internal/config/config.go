package config

import (
	"os"
	"strconv"
	"strings"
)

// Config holds all configuration for the application
type Config struct {
	Server    ServerConfig
	Database  DatabaseConfig
	CORS      CORSConfig
	RateLimit RateLimitConfig
}

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Port string
	Host string
}

// DatabaseConfig holds database-related configuration
type DatabaseConfig struct {
	Host       string
	Port       string
	User       string
	Password   string
	DBName     string
	TestDBName string
	SSLMode    string
}

// CORSConfig holds CORS-related configuration
type CORSConfig struct {
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowedHeaders   []string
	AllowCredentials bool
}

// RateLimitConfig holds rate limiting configuration
type RateLimitConfig struct {
	RequestsPerMinute int
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	return &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", "8080"),
			Host: getEnv("HOST", "localhost"),
		},
		Database: DatabaseConfig{
			Host:       getEnv("DB_HOST", "localhost"),
			Port:       getEnv("DB_PORT", "5432"),
			User:       getEnv("DB_USER", "postgres"),
			Password:   getEnv("DB_PASSWORD", ""),
			DBName:     getEnv("DB_NAME", "app"),
			TestDBName: getEnv("DB_TEST_NAME", "app_test"),
			SSLMode:    getEnv("DB_SSLMODE", "disable"),
		},
		CORS: CORSConfig{
			AllowedOrigins:   parseStringSlice(getEnv("CORS_ALLOWED_ORIGINS", "*"), ","),
			AllowedMethods:   parseStringSlice(getEnv("CORS_ALLOWED_METHODS", "GET,POST,PUT,DELETE,OPTIONS"), ","),
			AllowedHeaders:   parseStringSlice(getEnv("CORS_ALLOWED_HEADERS", "*"), ","),
			AllowCredentials: parseBool(getEnv("CORS_ALLOW_CREDENTIALS", "false")),
		},
		RateLimit: RateLimitConfig{
			RequestsPerMinute: parseInt(getEnv("RATE_LIMIT_REQUESTS_PER_MINUTE", "60")),
		},
	}, nil
}

// parseStringSlice parses a comma-separated string into a slice of strings
func parseStringSlice(value, sep string) []string {
	if value == "" {
		return []string{}
	}
	return strings.Split(value, sep)
}

// parseBool parses a string into a boolean
func parseBool(value string) bool {
	result, _ := strconv.ParseBool(value)
	return result
}

// parseInt parses a string into an int
func parseInt(value string) int {
	result, _ := strconv.Atoi(value)
	return result
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
