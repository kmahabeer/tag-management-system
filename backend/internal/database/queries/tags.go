package queries

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/kmahabeer/tag-management-system/backend/internal/models"
)

// TagQueries implements TagRepository
type TagQueries struct {
	db *sql.DB
}

// NewTagQueries creates a new TagQueries instance
func NewTagQueries(db *sql.DB) *TagQueries {
	return &TagQueries{db: db}
}

// GetByID retrieves a tag by its ID
func (q *TagQueries) GetByID(ctx context.Context, id uuid.UUID) (*models.Tag, error) {
	query := `
		SELECT id, name, display_name, metadata, part_of_speech_id, created_at, updated_at
		FROM tags
		WHERE id = $1
	`

	var tag models.Tag
	err := q.db.QueryRowContext(ctx, query, id).Scan(
		&tag.ID,
		&tag.Name,
		&tag.DisplayName,
		&tag.Metadata,
		&tag.PartOfSpeechID,
		&tag.CreatedAt,
		&tag.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("tag not found")
		}
		return nil, fmt.Errorf("failed to get tag: %w", err)
	}

	return &tag, nil
}

// List retrieves a list of tags with pagination
func (q *TagQueries) List(ctx context.Context, limit, offset int) ([]*models.Tag, error) {
	query := `
		SELECT id, name, display_name, metadata, part_of_speech_id, created_at, updated_at
		FROM tags
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := q.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list tags: %w", err)
	}
	defer rows.Close()

	var tags []*models.Tag
	for rows.Next() {
		var tag models.Tag
		err := rows.Scan(
			&tag.ID,
			&tag.Name,
			&tag.DisplayName,
			&tag.Metadata,
			&tag.PartOfSpeechID,
			&tag.CreatedAt,
			&tag.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan tag: %w", err)
		}
		tags = append(tags, &tag)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return tags, nil
}

// Create inserts a new tag
func (q *TagQueries) Create(ctx context.Context, tag *models.Tag) error {
	query := `
		INSERT INTO tags (id, name, display_name, metadata, part_of_speech_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := q.db.ExecContext(ctx, query,
		tag.ID,
		tag.Name,
		tag.DisplayName,
		tag.Metadata,
		tag.PartOfSpeechID,
		tag.CreatedAt,
		tag.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create tag: %w", err)
	}

	return nil
}

// Update modifies an existing tag
func (q *TagQueries) Update(ctx context.Context, tag *models.Tag) error {
	query := `
		UPDATE tags
		SET name = $2, display_name = $3, metadata = $4, part_of_speech_id = $5, updated_at = $6
		WHERE id = $1
	`

	result, err := q.db.ExecContext(ctx, query,
		tag.ID,
		tag.Name,
		tag.DisplayName,
		tag.Metadata,
		tag.PartOfSpeechID,
		tag.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to update tag: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("tag not found")
	}

	return nil
}

// Delete removes a tag by ID
func (q *TagQueries) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM tags WHERE id = $1`

	result, err := q.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete tag: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("tag not found")
	}

	return nil
}
