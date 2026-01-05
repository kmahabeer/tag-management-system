package models

import (
	"errors"
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

// Validate checks if the TagContextRating is valid
func (tcr *TagContextRating) Validate() error {
	if tcr.ID == uuid.Nil {
		return errors.New("id is required")
	}
	if tcr.TagID == uuid.Nil {
		return errors.New("tag_id is required")
	}
	if tcr.ContextID == uuid.Nil {
		return errors.New("context_id is required")
	}
	if tcr.RatingID == uuid.Nil {
		return errors.New("rating_id is required")
	}
	return nil
}
