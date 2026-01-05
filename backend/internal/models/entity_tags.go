package models

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type EntityTag struct {
	ID        uuid.UUID       `db:"id" json:"id"`
	EntityID  uuid.UUID       `db:"entity_id" json:"entity_id"`
	TagID     uuid.UUID       `db:"tag_id" json:"tag_id"`
	ContextID uuid.UUID       `db:"context_id" json:"context_id"`
	Metadata  json.RawMessage `db:"metadata" json:"metadata"`
	CreatedAt time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt time.Time       `db:"updated_at" json:"updated_at"`
}

// Validate checks if the EntityTag is valid
func (et *EntityTag) Validate() error {
	if et.ID == uuid.Nil {
		return errors.New("id is required")
	}
	if et.EntityID == uuid.Nil {
		return errors.New("entity_id is required")
	}
	if et.TagID == uuid.Nil {
		return errors.New("tag_id is required")
	}
	if et.ContextID == uuid.Nil {
		return errors.New("context_id is required")
	}
	return nil
}
