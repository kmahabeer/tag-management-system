package models

import (
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
