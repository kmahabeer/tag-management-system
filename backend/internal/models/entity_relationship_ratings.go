package models

import (
	"time"

	"github.com/google/uuid"
)

type EntityRelationshipRating struct {
	ID        uuid.UUID `db:"id" json:"id"`
	EntityAID uuid.UUID `db:"entity_a_id" json:"entity_a_id"`
	EntityBID uuid.UUID `db:"entity_b_id" json:"entity_b_id"`
	ContextID uuid.UUID `db:"context_id" json:"context_id"`
	RatingID  uuid.UUID `db:"rating_id" json:"rating_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
