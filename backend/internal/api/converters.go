package api

import (
	"database/sql"
	"encoding/json"

	"github.com/kmahabeer/tag-management-system/backend/internal/models"
)

// Null handling utilities

// stringToNullString converts a *string to sql.NullString
func stringToNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: *s, Valid: true}
}

// nullStringToStringPtr converts sql.NullString to *string
func nullStringToStringPtr(ns sql.NullString) *string {
	if !ns.Valid {
		return nil
	}
	return &ns.String
}

// metadataToRawMessage converts any to json.RawMessage, handling nil
func metadataToRawMessage(metadata any) json.RawMessage {
	if metadata == nil {
		return nil
	}
	data, _ := json.Marshal(metadata)
	return data
}

// rawMessageToAny converts json.RawMessage to any, handling null
func rawMessageToAny(raw json.RawMessage) any {
	if len(raw) == 0 || string(raw) == "null" {
		return nil
	}
	var result any
	json.Unmarshal(raw, &result)
	return result
}

// Tag conversions

// TagToAPI converts a database Tag model to API Tag schema
func TagToAPI(t models.Tag) Tag {
	tag := Tag{
		ID:             t.ID,
		Name:           t.Name,
		Metadata:       rawMessageToAny(t.Metadata),
		PartOfSpeechID: t.PartOfSpeechID,
		CreatedAt:      t.CreatedAt,
		UpdatedAt:      t.UpdatedAt,
	}

	tag.DisplayName = nullStringToStringPtr(t.DisplayName)

	// Note: Embedding field not in DB model yet, so omitted

	return tag
}

// TagInputToDB converts an API TagInput schema to database Tag model
func TagInputToDB(t TagInput) models.Tag {
	tag := models.Tag{
		Name:           t.Name,
		PartOfSpeechID: t.PartOfSpeechID,
	}

	tag.DisplayName = stringToNullString(t.DisplayName)
	tag.Metadata = metadataToRawMessage(t.Metadata)

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
		Metadata:  rawMessageToAny(e.Metadata),
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}

	entity.Location = nullStringToStringPtr(e.Location)

	// Note: Embedding field not in DB model yet, so omitted

	return entity
}

// EntityInputToDB converts an API EntityInput schema to database Entity model
func EntityInputToDB(e EntityInput) models.Entity {
	entity := models.Entity{
		Name:      e.Name,
		IsPrimary: e.IsPrimary,
	}

	entity.Location = stringToNullString(e.Location)
	entity.Metadata = metadataToRawMessage(e.Metadata)

	// Note: Embedding field not in DB model yet, so omitted

	return entity
}
