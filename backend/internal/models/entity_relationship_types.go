package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type EntityRelationshipType struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// Validate checks if the EntityRelationshipType is valid
func (ert *EntityRelationshipType) Validate() error {
	if ert.ID == uuid.Nil {
		return errors.New("id is required")
	}
	if ert.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
