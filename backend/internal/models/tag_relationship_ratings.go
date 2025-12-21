package models

import (
	"errors"
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

// Validate checks if the TagRelationshipRating is valid
func (trr *TagRelationshipRating) Validate() error {
	if trr.ID == uuid.Nil {
		return errors.New("id is required")
	}
	if trr.TagAID == uuid.Nil {
		return errors.New("tag_a_id is required")
	}
	if trr.TagBID == uuid.Nil {
		return errors.New("tag_b_id is required")
	}
	if trr.ContextID == uuid.Nil {
		return errors.New("context_id is required")
	}
	if trr.RatingID == uuid.Nil {
		return errors.New("rating_id is required")
	}
	return nil
}
