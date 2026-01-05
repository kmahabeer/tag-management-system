package queries

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/kmahabeer/tag-management-system/backend/internal/models"
)

// TagAliasQueries implements TagAliasRepository
type TagAliasQueries struct {
	db *sql.DB
}

// NewTagAliasQueries creates a new TagAliasQueries instance
func NewTagAliasQueries(db *sql.DB) *TagAliasQueries {
	return &TagAliasQueries{db: db}
}

// GetByID retrieves a tag alias by its ID
func (q *TagAliasQueries) GetByID(ctx context.Context, id uuid.UUID) (*models.TagAlias, error) {
	query := `
		SELECT id, name, tag_id, created_at, updated_at
		FROM tag_aliases
		WHERE id = $1
	`

	var tagAlias models.TagAlias
	err := q.db.QueryRowContext(ctx, query, id).Scan(
		&tagAlias.ID,
		&tagAlias.Name,
		&tagAlias.TagID,
		&tagAlias.CreatedAt,
		&tagAlias.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("tag alias not found")
		}
		return nil, fmt.Errorf("failed to get tag alias: %w", err)
	}

	return &tagAlias, nil
}

// List retrieves a list of tag aliases with pagination
func (q *TagAliasQueries) List(ctx context.Context, limit, offset int) ([]*models.TagAlias, error) {
	query := `
		SELECT id, name, tag_id, created_at, updated_at
		FROM tag_aliases
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := q.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list tag aliases: %w", err)
	}
	defer rows.Close()

	var tagAliases []*models.TagAlias
	for rows.Next() {
		var tagAlias models.TagAlias
		err := rows.Scan(
			&tagAlias.ID,
			&tagAlias.Name,
			&tagAlias.TagID,
			&tagAlias.CreatedAt,
			&tagAlias.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan tag alias: %w", err)
		}
		tagAliases = append(tagAliases, &tagAlias)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return tagAliases, nil
}

// Create inserts a new tag alias
func (q *TagAliasQueries) Create(ctx context.Context, tagAlias *models.TagAlias) error {
	query := `
		INSERT INTO tag_aliases (id, name, tag_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := q.db.ExecContext(ctx, query,
		tagAlias.ID,
		tagAlias.Name,
		tagAlias.TagID,
		tagAlias.CreatedAt,
		tagAlias.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create tag alias: %w", err)
	}

	return nil
}

// Update modifies an existing tag alias
func (q *TagAliasQueries) Update(ctx context.Context, tagAlias *models.TagAlias) error {
	query := `
		UPDATE tag_aliases
		SET name = $2, tag_id = $3, updated_at = $4
		WHERE id = $1
	`

	result, err := q.db.ExecContext(ctx, query,
		tagAlias.ID,
		tagAlias.Name,
		tagAlias.TagID,
		tagAlias.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to update tag alias: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("tag alias not found")
	}

	return nil
}

// Delete removes a tag alias by ID
func (q *TagAliasQueries) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM tag_aliases WHERE id = $1`

	result, err := q.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete tag alias: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("tag alias not found")
	}

	return nil
}
