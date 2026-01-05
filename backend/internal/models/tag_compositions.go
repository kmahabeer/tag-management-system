package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type TagComposition struct {
	ID             uuid.UUID `db:"id" json:"id"`
	BaseTagID      uuid.UUID `db:"base_tag_id" json:"base_tag_id"`
	ComponentTagID uuid.UUID `db:"component_tag_id" json:"component_tag_id"`
	Position       int       `db:"position" json:"position"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
}

// Validate checks if the TagComposition is valid
func (tc *TagComposition) Validate() error {
	if tc.ID == uuid.Nil {
		return errors.New("id is required")
	}
	if tc.BaseTagID == uuid.Nil {
		return errors.New("base_tag_id is required")
	}
	if tc.ComponentTagID == uuid.Nil {
		return errors.New("component_tag_id is required")
	}
	if tc.Position <= 0 {
		return errors.New("position must be greater than 0")
	}
	return nil
}
