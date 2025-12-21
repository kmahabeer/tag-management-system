package queries

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/kmahabeer/tag-management-system/backend/internal/models"
)

// PartOfSpeechQueries implements PartOfSpeechRepository
type PartOfSpeechQueries struct {
	db *sql.DB
}

// NewPartOfSpeechQueries creates a new PartOfSpeechQueries instance
func NewPartOfSpeechQueries(db *sql.DB) *PartOfSpeechQueries {
	return &PartOfSpeechQueries{db: db}
}

// GetByID retrieves a part of speech by its ID
func (q *PartOfSpeechQueries) GetByID(ctx context.Context, id uuid.UUID) (*models.PartOfSpeech, error) {
	query := `
		SELECT id, name, description, is_active, created_at, updated_at
		FROM parts_of_speech
		WHERE id = $1
	`

	var partOfSpeech models.PartOfSpeech
	err := q.db.QueryRowContext(ctx, query, id).Scan(
		&partOfSpeech.ID,
		&partOfSpeech.Name,
		&partOfSpeech.Description,
		&partOfSpeech.IsActive,
		&partOfSpeech.CreatedAt,
		&partOfSpeech.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("part of speech not found")
		}
		return nil, fmt.Errorf("failed to get part of speech: %w", err)
	}

	return &partOfSpeech, nil
}

// List retrieves a list of parts of speech with pagination
func (q *PartOfSpeechQueries) List(ctx context.Context, limit, offset int) ([]*models.PartOfSpeech, error) {
	query := `
		SELECT id, name, description, is_active, created_at, updated_at
		FROM parts_of_speech
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := q.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list parts of speech: %w", err)
	}
	defer rows.Close()

	var partsOfSpeech []*models.PartOfSpeech
	for rows.Next() {
		var partOfSpeech models.PartOfSpeech
		err := rows.Scan(
			&partOfSpeech.ID,
			&partOfSpeech.Name,
			&partOfSpeech.Description,
			&partOfSpeech.IsActive,
			&partOfSpeech.CreatedAt,
			&partOfSpeech.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan part of speech: %w", err)
		}
		partsOfSpeech = append(partsOfSpeech, &partOfSpeech)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return partsOfSpeech, nil
}

// Create inserts a new part of speech
func (q *PartOfSpeechQueries) Create(ctx context.Context, partOfSpeech *models.PartOfSpeech) error {
	query := `
		INSERT INTO parts_of_speech (id, name, description, is_active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := q.db.ExecContext(ctx, query,
		partOfSpeech.ID,
		partOfSpeech.Name,
		partOfSpeech.Description,
		partOfSpeech.IsActive,
		partOfSpeech.CreatedAt,
		partOfSpeech.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create part of speech: %w", err)
	}

	return nil
}

// Update modifies an existing part of speech
func (q *PartOfSpeechQueries) Update(ctx context.Context, partOfSpeech *models.PartOfSpeech) error {
	query := `
		UPDATE parts_of_speech
		SET name = $2, description = $3, is_active = $4, updated_at = $5
		WHERE id = $1
	`

	result, err := q.db.ExecContext(ctx, query,
		partOfSpeech.ID,
		partOfSpeech.Name,
		partOfSpeech.Description,
		partOfSpeech.IsActive,
		partOfSpeech.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to update part of speech: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("part of speech not found")
	}

	return nil
}

// Delete removes a part of speech by ID
func (q *PartOfSpeechQueries) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM parts_of_speech WHERE id = $1`

	result, err := q.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete part of speech: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("part of speech not found")
	}

	return nil
}
