package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Context struct {
	ID                 uuid.UUID      `db:"id" json:"id"`
	Name               string         `db:"name" json:"name"`
	ClassificationType string         `db:"classification_type" json:"classification_type"`
	Description        sql.NullString `db:"description" json:"description"`
	IsActive           bool           `db:"is_active" json:"is_active"`
	CreatedAt          time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt          time.Time      `db:"updated_at" json:"updated_at"`
}

// Validate checks if the Context is valid
func (c *Context) Validate() error {
	if c.ID == uuid.Nil {
		return errors.New("id is required")
	}
	if c.Name == "" {
		return errors.New("name is required")
	}
	if c.ClassificationType == "" {
		return errors.New("classification_type is required")
	}
	return nil
}
