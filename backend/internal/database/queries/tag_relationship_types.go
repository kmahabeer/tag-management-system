package queries

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/kmahabeer/tag-management-system/backend/internal/models"
)

// TagRelationshipTypeQueries implements TagRelationshipTypeRepository
type TagRelationshipTypeQueries struct {
	db *sql.DB
}

// NewTagRelationshipTypeQueries creates a new TagRelationshipTypeQueries instance
func NewTagRelationshipTypeQueries(db *sql.DB) *TagRelationshipTypeQueries {
	return &TagRelationshipTypeQueries{db: db}
}

// GetByID retrieves a tag relationship type by its ID
func (q *TagRelationshipTypeQueries) GetByID(ctx context.Context, id uuid.UUID) (*models.TagRelationshipType, error) {
	query := `
		SELECT id, name, created_at, updated_at
		FROM tag_relationship_types
		WHERE id = $1
	`

	var tagRelationshipType models.TagRelationshipType
	err := q.db.QueryRowContext(ctx, query, id).Scan(
		&tagRelationshipType.ID,
		&tagRelationshipType.Name,
		&tagRelationshipType.CreatedAt,
		&tagRelationshipType.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("tag relationship type not found")
		}
		return nil, fmt.Errorf("failed to get tag relationship type: %w", err)
	}

	return &tagRelationshipType, nil
}

// List retrieves a list of tag relationship types with pagination
func (q *TagRelationshipTypeQueries) List(ctx context.Context, limit, offset int) ([]*models.TagRelationshipType, error) {
	query := `
		SELECT id, name, created_at, updated_at
		FROM tag_relationship_types
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := q.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list tag relationship types: %w", err)
	}
	defer rows.Close()

	var tagRelationshipTypes []*models.TagRelationshipType
	for rows.Next() {
		var tagRelationshipType models.TagRelationshipType
		err := rows.Scan(
			&tagRelationshipType.ID,
			&tagRelationshipType.Name,
			&tagRelationshipType.CreatedAt,
			&tagRelationshipType.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan tag relationship type: %w", err)
		}
		tagRelationshipTypes = append(tagRelationshipTypes, &tagRelationshipType)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return tagRelationshipTypes, nil
}

// Create inserts a new tag relationship type
func (q *TagRelationshipTypeQueries) Create(ctx context.Context, tagRelationshipType *models.TagRelationshipType) error {
	query := `
		INSERT INTO tag_relationship_types (id, name, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
	`

	_, err := q.db.ExecContext(ctx, query,
		tagRelationshipType.ID,
		tagRelationshipType.Name,
		tagRelationshipType.CreatedAt,
		tagRelationshipType.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create tag relationship type: %w", err)
	}

	return nil
}

// Update modifies an existing tag relationship type
func (q *TagRelationshipTypeQueries) Update(ctx context.Context, tagRelationshipType *models.TagRelationshipType) error {
	query := `
		UPDATE tag_relationship_types
		SET name = $2, updated_at = $3
		WHERE id = $1
	`

	result, err := q.db.ExecContext(ctx, query,
		tagRelationshipType.ID,
		tagRelationshipType.Name,
		tagRelationshipType.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to update tag relationship type: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("tag relationship type not found")
	}

	return nil
}

// Delete removes a tag relationship type by ID
func (q *TagRelationshipTypeQueries) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM tag_relationship_types WHERE id = $1`

	result, err := q.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete tag relationship type: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("tag relationship type not found")
	}

	return nil
}
