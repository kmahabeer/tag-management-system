package models

import (
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
