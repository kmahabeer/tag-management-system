package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type TagAlias struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	TagID     uuid.UUID `db:"tag_id" json:"tag_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// Validate checks if the TagAlias is valid
func (ta *TagAlias) Validate() error {
	if ta.ID == uuid.Nil {
		return errors.New("id is required")
	}
	if ta.Name == "" {
		return errors.New("name is required")
	}
	if ta.TagID == uuid.Nil {
		return errors.New("tag_id is required")
	}
	return nil
}
