package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type RatingType struct {
	ID           uuid.UUID `db:"id" json:"id"`
	Name         string    `db:"name" json:"name"`
	IsNormalized bool      `db:"is_normalized" json:"is_normalized"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}

// Validate checks if the RatingType is valid
func (rt *RatingType) Validate() error {
	if rt.ID == uuid.Nil {
		return errors.New("id is required")
	}
	if rt.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
