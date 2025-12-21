package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

type UILayout struct {
	ID           uuid.UUID      `db:"id" json:"id"`
	PurposeTagID *uuid.UUID     `db:"purpose_tag_id" json:"purpose_tag_id"`
	Name         sql.NullString `db:"name" json:"name"`
	CreatedAt    time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time      `db:"updated_at" json:"updated_at"`
}

// Validate checks if the UILayout is valid
func (uil *UILayout) Validate() error {
	if uil.ID == uuid.Nil {
		return errors.New("id is required")
	}
	return nil
}
