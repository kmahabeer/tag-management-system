package models

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// Tag represents a tag in the database
type Tag struct {
	ID             uuid.UUID       `db:"id" json:"id"`
	Name           string          `db:"name" json:"name"`
	DisplayName    sql.NullString  `db:"display_name" json:"display_name"`
	Metadata       json.RawMessage `db:"metadata" json:"metadata"`
	PartOfSpeechID uuid.UUID       `db:"part_of_speech_id" json:"part_of_speech_id"`
	CreatedAt      time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time       `db:"updated_at" json:"updated_at"`
}
