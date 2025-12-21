package models

import (
	"time"

	"github.com/google/uuid"
)

type UIField struct {
	ID            uuid.UUID `db:"id" json:"id"`
	UILayoutID    uuid.UUID `db:"ui_layout_id" json:"ui_layout_id"`
	UIGroupID     uuid.UUID `db:"ui_group_id" json:"ui_group_id"`
	ContextID     uuid.UUID `db:"context_id" json:"context_id"`
	CategoryTagID uuid.UUID `db:"category_tag_id" json:"category_tag_id"`
	SortOrder     int       `db:"sort_order" json:"sort_order"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}
