package models

import (
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
