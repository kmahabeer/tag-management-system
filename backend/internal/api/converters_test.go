package api

import (
	"database/sql"
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/kmahabeer/tag-management-system/backend/internal/models"
)

func TestTagConversions(t *testing.T) {
	// Test DB to API conversion
	dbTag := models.Tag{
		ID:             uuid.MustParse("a123e456-78b9-4cde-8123-456789abcdef"),
		Name:           "test tag",
		DisplayName:    sql.NullString{String: "Test Tag", Valid: true},
		Metadata:       json.RawMessage(`{"source": "test"}`),
		PartOfSpeechID: uuid.MustParse("046b6c7f-0b8a-43b9-b35d-6489e6daee91"),
		CreatedAt:      time.Date(2025, 10, 7, 12, 0, 0, 0, time.UTC),
		UpdatedAt:      time.Date(2025, 10, 7, 12, 30, 0, 0, time.UTC),
	}

	apiTag := TagToAPI(dbTag)

	if apiTag.ID != dbTag.ID {
		t.Errorf("ID mismatch: got %v, want %v", apiTag.ID, dbTag.ID)
	}
	if apiTag.Name != dbTag.Name {
		t.Errorf("Name mismatch: got %v, want %v", apiTag.Name, dbTag.Name)
	}
	if apiTag.DisplayName == nil || *apiTag.DisplayName != "Test Tag" {
		t.Errorf("DisplayName mismatch: got %v, want Test Tag", apiTag.DisplayName)
	}

	// Test API to DB conversion
	apiInput := TagInput{
		Name:           "new tag",
		DisplayName:    stringPtr("New Tag"),
		Metadata:       map[string]interface{}{"source": "api"},
		PartOfSpeechID: uuid.MustParse("046b6c7f-0b8a-43b9-b35d-6489e6daee91"),
	}

	dbTagFromAPI := TagInputToDB(apiInput)

	if dbTagFromAPI.Name != apiInput.Name {
		t.Errorf("Name mismatch: got %v, want %v", dbTagFromAPI.Name, apiInput.Name)
	}
	if !dbTagFromAPI.DisplayName.Valid || dbTagFromAPI.DisplayName.String != "New Tag" {
		t.Errorf("DisplayName mismatch: got %v, want New Tag", dbTagFromAPI.DisplayName)
	}
}

func TestEntityConversions(t *testing.T) {
	// Test DB to API conversion
	dbEntity := models.Entity{
		ID:        uuid.MustParse("b123e456-78b9-4cde-8123-456789abcdef"),
		Name:      "test entity",
		Location:  sql.NullString{String: "test location", Valid: true},
		IsPrimary: true,
		Metadata:  json.RawMessage(`{"type": "image"}`),
		CreatedAt: time.Date(2025, 10, 7, 12, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2025, 10, 7, 12, 30, 0, 0, time.UTC),
	}

	apiEntity := EntityToAPI(dbEntity)

	if apiEntity.ID != dbEntity.ID {
		t.Errorf("ID mismatch: got %v, want %v", apiEntity.ID, dbEntity.ID)
	}
	if apiEntity.Name != dbEntity.Name {
		t.Errorf("Name mismatch: got %v, want %v", apiEntity.Name, dbEntity.Name)
	}
	if apiEntity.Location == nil || *apiEntity.Location != "test location" {
		t.Errorf("Location mismatch: got %v, want test location", apiEntity.Location)
	}

	// Test API to DB conversion
	apiInput := EntityInput{
		Name:      "new entity",
		Location:  stringPtr("new location"),
		IsPrimary: false,
		Metadata:  map[string]interface{}{"type": "document"},
	}

	dbEntityFromAPI := EntityInputToDB(apiInput)

	if dbEntityFromAPI.Name != apiInput.Name {
		t.Errorf("Name mismatch: got %v, want %v", dbEntityFromAPI.Name, apiInput.Name)
	}
	if !dbEntityFromAPI.Location.Valid || dbEntityFromAPI.Location.String != "new location" {
		t.Errorf("Location mismatch: got %v, want new location", dbEntityFromAPI.Location)
	}
}
