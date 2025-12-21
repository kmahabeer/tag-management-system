package models

import (
	"time"

	"github.com/google/uuid"
)

type TagContextRating struct {
	ID        uuid.UUID  `db:"id" json:"id"`
	TagID     uuid.UUID  `db:"tag_id" json:"tag_id"`
	ContextID uuid.UUID  `db:"context_id" json:"context_id"`
	RatingID  uuid.UUID  `db:"rating_id" json:"rating_id"`
	UserID    *uuid.UUID `db:"user_id" json:"user_id"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
}
