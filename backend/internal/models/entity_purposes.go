package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type EntityPurpose struct {
	ID           uuid.UUID `db:"id" json:"id"`
	EntityID     uuid.UUID `db:"entity_id" json:"entity_id"`
	PurposeTagID uuid.UUID `db:"purpose_tag_id" json:"purpose_tag_id"`
	IsPrimary    bool      `db:"is_primary" json:"is_primary"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}

// Validate checks if the EntityPurpose is valid
func (ep *EntityPurpose) Validate() error {
	if ep.ID == uuid.Nil {
		return errors.New("id is required")
	}
	if ep.EntityID == uuid.Nil {
		return errors.New("entity_id is required")
	}
	if ep.PurposeTagID == uuid.Nil {
		return errors.New("purpose_tag_id is required")
	}
	return nil
}
