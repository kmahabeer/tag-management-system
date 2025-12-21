package models

import (
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Entity represents an entity in the database
type Entity struct {
	ID        uuid.UUID       `db:"id" json:"id"`
	Name      string          `db:"name" json:"name"`
	Location  sql.NullString  `db:"location" json:"location"`
	IsPrimary bool            `db:"is_primary" json:"is_primary"`
	Metadata  json.RawMessage `db:"metadata" json:"metadata"`
	CreatedAt time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt time.Time       `db:"updated_at" json:"updated_at"`
}

// Validate checks if the Entity is valid
func (e *Entity) Validate() error {
	if e.ID == uuid.Nil {
		return errors.New("id is required")
	}
	if e.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
