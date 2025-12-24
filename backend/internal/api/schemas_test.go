package api

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestTagInputValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   TagInput
		wantErr bool
	}{
		{
			name: "valid input",
			input: TagInput{
				Name:           "test tag",
				PartOfSpeechID: uuid.New(),
			},
			wantErr: false,
		},
		{
			name: "empty name",
			input: TagInput{
				Name:           "",
				PartOfSpeechID: uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "nil part of speech id",
			input: TagInput{
				Name:           "test tag",
				PartOfSpeechID: uuid.Nil,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.input.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("TagInput.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEntityInputValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   EntityInput
		wantErr bool
	}{
		{
			name: "valid input",
			input: EntityInput{
				Name: "test entity",
			},
			wantErr: false,
		},
		{
			name: "empty name",
			input: EntityInput{
				Name: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.input.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("EntityInput.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTagJSONMarshaling(t *testing.T) {
	tag := Tag{
		ID:             uuid.MustParse("a123e456-78b9-4cde-8123-456789abcdef"),
		Name:           "dog",
		DisplayName:    stringPtr("Dog"),
		Metadata:       map[string]interface{}{"source": "Label Studio"},
		PartOfSpeechID: uuid.MustParse("046b6c7f-0b8a-43b9-b35d-6489e6daee91"),
		CreatedAt:      time.Date(2025, 10, 7, 12, 0, 0, 0, time.UTC),
		UpdatedAt:      time.Date(2025, 10, 7, 12, 30, 0, 0, time.UTC),
	}

	data, err := json.Marshal(tag)
	if err != nil {
		t.Fatalf("Failed to marshal Tag: %v", err)
	}

	var unmarshaled Tag
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal Tag: %v", err)
	}

	if unmarshaled.ID != tag.ID {
		t.Errorf("ID mismatch: got %v, want %v", unmarshaled.ID, tag.ID)
	}
	if unmarshaled.Name != tag.Name {
		t.Errorf("Name mismatch: got %v, want %v", unmarshaled.Name, tag.Name)
	}
}

func stringPtr(s string) *string {
	return &s
}

func TestLabelStudioExportJSONMarshaling(t *testing.T) {
	export := LabelStudioExport{
		Data: LabelStudioData{
			Image: "https://example.com/image.jpg",
		},
		Annotations: []LabelStudioAnnotation{
			{
				Result: []LabelStudioResult{
					{
						Value: LabelStudioValue{
							Labels: []string{"Red", "Car"},
						},
						FromName: "label",
						ToName:   "image",
						Type:     "choices",
					},
				},
			},
		},
	}

	data, err := json.Marshal(export)
	if err != nil {
		t.Fatalf("Failed to marshal LabelStudioExport: %v", err)
	}

	var unmarshaled LabelStudioExport
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal LabelStudioExport: %v", err)
	}

	if unmarshaled.Data.Image != export.Data.Image {
		t.Errorf("Image mismatch: got %v, want %v", unmarshaled.Data.Image, export.Data.Image)
	}
	if len(unmarshaled.Annotations[0].Result[0].Value.Labels) != 2 {
		t.Errorf("Labels count mismatch")
	}
}
