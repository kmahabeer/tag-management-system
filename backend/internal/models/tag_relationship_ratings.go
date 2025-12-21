package models

import (
	"time"

	"github.com/google/uuid"
)

type TagRelationshipRating struct {
	ID        uuid.UUID `db:"id" json:"id"`
	TagAID    uuid.UUID `db:"tag_a_id" json:"tag_a_id"`
	TagBID    uuid.UUID `db:"tag_b_id" json:"tag_b_id"`
	ContextID uuid.UUID `db:"context_id" json:"context_id"`
	RatingID  uuid.UUID `db:"rating_id" json:"rating_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
