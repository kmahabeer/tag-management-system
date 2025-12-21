package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type TagRelationshipType struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// Validate checks if the TagRelationshipType is valid
func (trt *TagRelationshipType) Validate() error {
	if trt.ID == uuid.Nil {
		return errors.New("id is required")
	}
	if trt.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
