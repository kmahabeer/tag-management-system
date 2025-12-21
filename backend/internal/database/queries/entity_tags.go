package queries

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/kmahabeer/tag-management-system/backend/internal/models"
)

// EntityTagQueries implements EntityTagRepository
type EntityTagQueries struct {
	db *sql.DB
}

// NewEntityTagQueries creates a new EntityTagQueries instance
func NewEntityTagQueries(db *sql.DB) *EntityTagQueries {
	return &EntityTagQueries{db: db}
}

// GetByID retrieves an entity tag by its ID
func (q *EntityTagQueries) GetByID(ctx context.Context, id uuid.UUID) (*models.EntityTag, error) {
	query := `
		SELECT id, entity_id, tag_id, context_id, metadata, created_at, updated_at
		FROM entity_tags
		WHERE id = $1
	`

	var entityTag models.EntityTag
	err := q.db.QueryRowContext(ctx, query, id).Scan(
		&entityTag.ID,
		&entityTag.EntityID,
		&entityTag.TagID,
		&entityTag.ContextID,
		&entityTag.Metadata,
		&entityTag.CreatedAt,
		&entityTag.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("entity tag not found")
		}
		return nil, fmt.Errorf("failed to get entity tag: %w", err)
	}

	return &entityTag, nil
}

// List retrieves a list of entity tags with pagination
func (q *EntityTagQueries) List(ctx context.Context, limit, offset int) ([]*models.EntityTag, error) {
	query := `
		SELECT id, entity_id, tag_id, context_id, metadata, created_at, updated_at
		FROM entity_tags
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := q.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list entity tags: %w", err)
	}
	defer rows.Close()

	var entityTags []*models.EntityTag
	for rows.Next() {
		var entityTag models.EntityTag
		err := rows.Scan(
			&entityTag.ID,
			&entityTag.EntityID,
			&entityTag.TagID,
			&entityTag.ContextID,
			&entityTag.Metadata,
			&entityTag.CreatedAt,
			&entityTag.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan entity tag: %w", err)
		}
		entityTags = append(entityTags, &entityTag)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return entityTags, nil
}

// Create inserts a new entity tag
func (q *EntityTagQueries) Create(ctx context.Context, entityTag *models.EntityTag) error {
	query := `
		INSERT INTO entity_tags (id, entity_id, tag_id, context_id, metadata, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := q.db.ExecContext(ctx, query,
		entityTag.ID,
		entityTag.EntityID,
		entityTag.TagID,
		entityTag.ContextID,
		entityTag.Metadata,
		entityTag.CreatedAt,
		entityTag.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create entity tag: %w", err)
	}

	return nil
}

// Update modifies an existing entity tag
func (q *EntityTagQueries) Update(ctx context.Context, entityTag *models.EntityTag) error {
	query := `
		UPDATE entity_tags
		SET entity_id = $2, tag_id = $3, context_id = $4, metadata = $5, updated_at = $6
		WHERE id = $1
	`

	result, err := q.db.ExecContext(ctx, query,
		entityTag.ID,
		entityTag.EntityID,
		entityTag.TagID,
		entityTag.ContextID,
		entityTag.Metadata,
		entityTag.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to update entity tag: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("entity tag not found")
	}

	return nil
}

// Delete removes an entity tag by ID
func (q *EntityTagQueries) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM entity_tags WHERE id = $1`

	result, err := q.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete entity tag: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("entity tag not found")
	}

	return nil
}
