package models

import (
	"errors"
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

// Validate checks if the EntityRelationshipRating is valid
func (err *EntityRelationshipRating) Validate() error {
	if err.ID == uuid.Nil {
		return errors.New("id is required")
	}
	if err.EntityAID == uuid.Nil {
		return errors.New("entity_a_id is required")
	}
	if err.EntityBID == uuid.Nil {
		return errors.New("entity_b_id is required")
	}
	if err.EntityAID == err.EntityBID {
		return errors.New("entity_a_id cannot be equal to entity_b_id")
	}
	if err.ContextID == uuid.Nil {
		return errors.New("context_id is required")
	}
	if err.RatingID == uuid.Nil {
		return errors.New("rating_id is required")
	}
	return nil
}
