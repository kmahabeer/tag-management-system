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
	Security  SecurityConfig
	Logging   LoggingConfig
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

// SecurityConfig holds security-related configuration
type SecurityConfig struct {
	ContentSecurityPolicy   string
	StrictTransportSecurity string
	XContentTypeOptions     string
	XFrameOptions           string
	XXSSProtection          string
	ReferrerPolicy          string
	PermissionsPolicy       string
	ServerHeader            string
	RemoveXPoweredBy        bool
	CacheControl            string
}

// LoggingConfig holds logging-related configuration
type LoggingConfig struct {
	Level  string
	Format string
	Output string
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
		Security: SecurityConfig{
			ContentSecurityPolicy:   getEnv("SECURITY_CONTENT_SECURITY_POLICY", "default-src 'self'"),
			StrictTransportSecurity: getEnv("SECURITY_STRICT_TRANSPORT_SECURITY", "max-age=31536000; includeSubDomains"),
			XContentTypeOptions:     getEnv("SECURITY_X_CONTENT_TYPE_OPTIONS", "nosniff"),
			XFrameOptions:           getEnv("SECURITY_X_FRAME_OPTIONS", "DENY"),
			XXSSProtection:          getEnv("SECURITY_X_XSS_PROTECTION", "1; mode=block"),
			ReferrerPolicy:          getEnv("SECURITY_REFERRER_POLICY", "strict-origin-when-cross-origin"),
			PermissionsPolicy:       getEnv("SECURITY_PERMISSIONS_POLICY", ""),
			ServerHeader:            getEnv("SECURITY_SERVER_HEADER", ""),
			RemoveXPoweredBy:        parseBool(getEnv("SECURITY_REMOVE_X_POWERED_BY", "true")),
			CacheControl:            getEnv("SECURITY_CACHE_CONTROL", "no-cache, no-store, must-revalidate"),
		},
		Logging: LoggingConfig{
			Level:  getEnv("LOG_LEVEL", "info"),
			Format: getEnv("LOG_FORMAT", "json"),
			Output: getEnv("LOG_OUTPUT", "stdout"),
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
