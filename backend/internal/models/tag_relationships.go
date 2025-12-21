package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type TagRelationship struct {
	ID                 uuid.UUID      `db:"id" json:"id"`
	TagAID             uuid.UUID      `db:"tag_a_id" json:"tag_a_id"`
	TagBID             uuid.UUID      `db:"tag_b_id" json:"tag_b_id"`
	RelationshipTypeID uuid.UUID      `db:"relationship_type_id" json:"relationship_type_id"`
	Description        sql.NullString `db:"description" json:"description"`
	CreatedAt          time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt          time.Time      `db:"updated_at" json:"updated_at"`
}
