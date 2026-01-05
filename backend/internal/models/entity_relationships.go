package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type EntityRelationship struct {
	ID                 uuid.UUID `db:"id" json:"id"`
	EntityAID          uuid.UUID `db:"entity_a_id" json:"entity_a_id"`
	EntityBID          uuid.UUID `db:"entity_b_id" json:"entity_b_id"`
	RelationshipTypeID uuid.UUID `db:"relationship_type_id" json:"relationship_type_id"`
	CreatedAt          time.Time `db:"created_at" json:"created_at"`
	UpdatedAt          time.Time `db:"updated_at" json:"updated_at"`
}

// Validate checks if the EntityRelationship is valid
func (er *EntityRelationship) Validate() error {
	if er.ID == uuid.Nil {
		return errors.New("id is required")
	}
	if er.EntityAID == uuid.Nil {
		return errors.New("entity_a_id is required")
	}
	if er.EntityBID == uuid.Nil {
		return errors.New("entity_b_id is required")
	}
	if er.RelationshipTypeID == uuid.Nil {
		return errors.New("relationship_type_id is required")
	}
	if er.EntityAID == er.EntityBID {
		return errors.New("entity_a_id and entity_b_id cannot be the same")
	}
	return nil
}
