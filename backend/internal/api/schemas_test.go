package api

import (
	"testing"

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
