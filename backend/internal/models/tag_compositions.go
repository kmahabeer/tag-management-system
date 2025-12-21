package models

import (
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
