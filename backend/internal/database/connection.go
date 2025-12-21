package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/kmahabeer/tag-management-system/backend/internal/config"
	_ "github.com/lib/pq"
)

// DB wraps sql.DB for database operations
type DB struct {
	*sql.DB
}

// NewDB creates a new database connection with connection pooling
func NewDB(cfg *config.DatabaseConfig) (*DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &DB{db}, nil
}

// InitDB initializes the database connection and runs migrations
func InitDB(cfg *config.DatabaseConfig, migrationsDir string) (*DB, error) {
	db, err := NewDB(cfg)
	if err != nil {
		return nil, err
	}

	// Run migrations
	if err := db.Migrate(migrationsDir); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

// Close closes the database connection
func (db *DB) Close() error {
	return db.DB.Close()
}
