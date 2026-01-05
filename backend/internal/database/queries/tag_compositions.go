package queries

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/kmahabeer/tag-management-system/backend/internal/models"
)

// TagCompositionQueries implements TagCompositionRepository
type TagCompositionQueries struct {
	db *sql.DB
}

// NewTagCompositionQueries creates a new TagCompositionQueries instance
func NewTagCompositionQueries(db *sql.DB) *TagCompositionQueries {
	return &TagCompositionQueries{db: db}
}

// GetByID retrieves a tag composition by its ID
func (q *TagCompositionQueries) GetByID(ctx context.Context, id uuid.UUID) (*models.TagComposition, error) {
	query := `
		SELECT id, base_tag_id, component_tag_id, position, created_at, updated_at
		FROM tag_compositions
		WHERE id = $1
	`

	var tagComposition models.TagComposition
	err := q.db.QueryRowContext(ctx, query, id).Scan(
		&tagComposition.ID,
		&tagComposition.BaseTagID,
		&tagComposition.ComponentTagID,
		&tagComposition.Position,
		&tagComposition.CreatedAt,
		&tagComposition.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("tag composition not found")
		}
		return nil, fmt.Errorf("failed to get tag composition: %w", err)
	}

	return &tagComposition, nil
}

// List retrieves a list of tag compositions with pagination
func (q *TagCompositionQueries) List(ctx context.Context, limit, offset int) ([]*models.TagComposition, error) {
	query := `
		SELECT id, base_tag_id, component_tag_id, position, created_at, updated_at
		FROM tag_compositions
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := q.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list tag compositions: %w", err)
	}
	defer rows.Close()

	var tagCompositions []*models.TagComposition
	for rows.Next() {
		var tagComposition models.TagComposition
		err := rows.Scan(
			&tagComposition.ID,
			&tagComposition.BaseTagID,
			&tagComposition.ComponentTagID,
			&tagComposition.Position,
			&tagComposition.CreatedAt,
			&tagComposition.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan tag composition: %w", err)
		}
		tagCompositions = append(tagCompositions, &tagComposition)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return tagCompositions, nil
}

// Create inserts a new tag composition
func (q *TagCompositionQueries) Create(ctx context.Context, tagComposition *models.TagComposition) error {
	query := `
		INSERT INTO tag_compositions (id, base_tag_id, component_tag_id, position, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := q.db.ExecContext(ctx, query,
		tagComposition.ID,
		tagComposition.BaseTagID,
		tagComposition.ComponentTagID,
		tagComposition.Position,
		tagComposition.CreatedAt,
		tagComposition.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create tag composition: %w", err)
	}

	return nil
}

// Update modifies an existing tag composition
func (q *TagCompositionQueries) Update(ctx context.Context, tagComposition *models.TagComposition) error {
	query := `
		UPDATE tag_compositions
		SET base_tag_id = $2, component_tag_id = $3, position = $4, updated_at = $5
		WHERE id = $1
	`

	result, err := q.db.ExecContext(ctx, query,
		tagComposition.ID,
		tagComposition.BaseTagID,
		tagComposition.ComponentTagID,
		tagComposition.Position,
		tagComposition.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to update tag composition: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("tag composition not found")
	}

	return nil
}

// Delete removes a tag composition by ID
func (q *TagCompositionQueries) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM tag_compositions WHERE id = $1`

	result, err := q.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete tag composition: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("tag composition not found")
	}

	return nil
}
