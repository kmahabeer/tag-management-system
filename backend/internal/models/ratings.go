package models

import (
	"time"

	"github.com/google/uuid"
)

type Rating struct {
	ID           uuid.UUID `db:"id" json:"id"`
	Name         string    `db:"name" json:"name"`
	Score        int       `db:"score" json:"score"`
	Description  *string   `db:"description" json:"description"`
	RatingTypeID uuid.UUID `db:"rating_type_id" json:"rating_type_id"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}
