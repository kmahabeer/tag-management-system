package database

import (
	"fmt"
	"os"
	"path/filepath"
)

// Migrate runs database migrations
func (db *DB) Migrate(migrationsDir string) error {
	// Get all SQL files in the migrations directory
	files, err := filepath.Glob(filepath.Join(migrationsDir, "*.sql"))
	if err != nil {
		return fmt.Errorf("failed to list migration files: %w", err)
	}

	// Sort files by name (assuming they are numbered)
	// filepath.Glob returns them in lexical order, which should work for numbered files

	for _, file := range files {
		if err := db.runMigrationFile(file); err != nil {
			return fmt.Errorf("failed to run migration %s: %w", file, err)
		}
	}

	return nil
}

// runMigrationFile executes a single SQL migration file
func (db *DB) runMigrationFile(filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read migration file %s: %w", filename, err)
	}

	// Execute the SQL
	_, err = db.Exec(string(content))
	if err != nil {
		return fmt.Errorf("failed to execute migration %s: %w", filename, err)
	}

	return nil
}
