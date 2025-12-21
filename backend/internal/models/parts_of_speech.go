package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

type PartOfSpeech struct {
	ID          uuid.UUID      `db:"id" json:"id"`
	Name        string         `db:"name" json:"name"`
	Description sql.NullString `db:"description" json:"description"`
	IsActive    bool           `db:"is_active" json:"is_active"`
	CreatedAt   time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at" json:"updated_at"`
}

// Validate checks if the PartOfSpeech is valid
func (pos *PartOfSpeech) Validate() error {
	if pos.ID == uuid.Nil {
		return errors.New("id is required")
	}
	if pos.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
