package api

import (
	"database/sql"
	"encoding/json"

	"github.com/kmahabeer/tag-management-system/backend/internal/models"
)

// Tag conversions

// TagToAPI converts a database Tag model to API Tag schema
func TagToAPI(t models.Tag) Tag {
	tag := Tag{
		ID:             t.ID,
		Name:           t.Name,
		Metadata:       t.Metadata,
		PartOfSpeechID: t.PartOfSpeechID,
		CreatedAt:      t.CreatedAt,
		UpdatedAt:      t.UpdatedAt,
	}

	if t.DisplayName.Valid {
		tag.DisplayName = &t.DisplayName.String
	}

	// Note: Embedding field not in DB model yet, so omitted

	return tag
}

// TagInputToDB converts an API TagInput schema to database Tag model
func TagInputToDB(t TagInput) models.Tag {
	tag := models.Tag{
		Name:           t.Name,
		PartOfSpeechID: t.PartOfSpeechID,
	}

	if t.DisplayName != nil {
		tag.DisplayName = sql.NullString{String: *t.DisplayName, Valid: true}
	}

	if t.Metadata != nil {
		metadata, _ := json.Marshal(t.Metadata)
		tag.Metadata = metadata
	}

	// Note: Embedding field not in DB model yet, so omitted

	return tag
}

// Entity conversions

// EntityToAPI converts a database Entity model to API Entity schema
func EntityToAPI(e models.Entity) Entity {
	entity := Entity{
		ID:        e.ID,
		Name:      e.Name,
		IsPrimary: e.IsPrimary,
		Metadata:  e.Metadata,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}

	if e.Location.Valid {
		entity.Location = &e.Location.String
	}

	// Note: Embedding field not in DB model yet, so omitted

	return entity
}

// EntityInputToDB converts an API EntityInput schema to database Entity model
func EntityInputToDB(e EntityInput) models.Entity {
	entity := models.Entity{
		Name:      e.Name,
		IsPrimary: e.IsPrimary,
	}

	if e.Location != nil {
		entity.Location = sql.NullString{String: *e.Location, Valid: true}
	}

	if e.Metadata != nil {
		metadata, _ := json.Marshal(e.Metadata)
		entity.Metadata = metadata
	}

	// Note: Embedding field not in DB model yet, so omitted

	return entity
}
