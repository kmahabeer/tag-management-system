package models

import (
	"database/sql"
	"errors"
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

// Validate checks if the TagRelationship is valid
func (tr *TagRelationship) Validate() error {
	if tr.ID == uuid.Nil {
		return errors.New("id is required")
	}
	if tr.TagAID == uuid.Nil {
		return errors.New("tag_a_id is required")
	}
	if tr.TagBID == uuid.Nil {
		return errors.New("tag_b_id is required")
	}
	if tr.TagAID == tr.TagBID {
		return errors.New("tag_a_id cannot be equal to tag_b_id")
	}
	if tr.RelationshipTypeID == uuid.Nil {
		return errors.New("relationship_type_id is required")
	}
	return nil
}
