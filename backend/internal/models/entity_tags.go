package models

import (
	"encoding/json"
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
