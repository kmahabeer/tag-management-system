package queries

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/kmahabeer/tag-management-system/backend/internal/models"
)

// TagRelationshipQueries implements TagRelationshipRepository
type TagRelationshipQueries struct {
	db *sql.DB
}

// NewTagRelationshipQueries creates a new TagRelationshipQueries instance
func NewTagRelationshipQueries(db *sql.DB) *TagRelationshipQueries {
	return &TagRelationshipQueries{db: db}
}

// GetByID retrieves a tag relationship by its ID
func (q *TagRelationshipQueries) GetByID(ctx context.Context, id uuid.UUID) (*models.TagRelationship, error) {
	query := `
		SELECT id, tag_a_id, tag_b_id, relationship_type_id, description, created_at, updated_at
		FROM tag_relationships
		WHERE id = $1
	`

	var tagRelationship models.TagRelationship
	err := q.db.QueryRowContext(ctx, query, id).Scan(
		&tagRelationship.ID,
		&tagRelationship.TagAID,
		&tagRelationship.TagBID,
		&tagRelationship.RelationshipTypeID,
		&tagRelationship.Description,
		&tagRelationship.CreatedAt,
		&tagRelationship.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("tag relationship not found")
		}
		return nil, fmt.Errorf("failed to get tag relationship: %w", err)
	}

	return &tagRelationship, nil
}

// List retrieves a list of tag relationships with pagination
func (q *TagRelationshipQueries) List(ctx context.Context, limit, offset int) ([]*models.TagRelationship, error) {
	query := `
		SELECT id, tag_a_id, tag_b_id, relationship_type_id, description, created_at, updated_at
		FROM tag_relationships
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := q.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list tag relationships: %w", err)
	}
	defer rows.Close()

	var tagRelationships []*models.TagRelationship
	for rows.Next() {
		var tagRelationship models.TagRelationship
		err := rows.Scan(
			&tagRelationship.ID,
			&tagRelationship.TagAID,
			&tagRelationship.TagBID,
			&tagRelationship.RelationshipTypeID,
			&tagRelationship.Description,
			&tagRelationship.CreatedAt,
			&tagRelationship.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan tag relationship: %w", err)
		}
		tagRelationships = append(tagRelationships, &tagRelationship)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return tagRelationships, nil
}

// Create inserts a new tag relationship
func (q *TagRelationshipQueries) Create(ctx context.Context, tagRelationship *models.TagRelationship) error {
	query := `
		INSERT INTO tag_relationships (id, tag_a_id, tag_b_id, relationship_type_id, description, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := q.db.ExecContext(ctx, query,
		tagRelationship.ID,
		tagRelationship.TagAID,
		tagRelationship.TagBID,
		tagRelationship.RelationshipTypeID,
		tagRelationship.Description,
		tagRelationship.CreatedAt,
		tagRelationship.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create tag relationship: %w", err)
	}

	return nil
}

// Update modifies an existing tag relationship
func (q *TagRelationshipQueries) Update(ctx context.Context, tagRelationship *models.TagRelationship) error {
	query := `
		UPDATE tag_relationships
		SET tag_a_id = $2, tag_b_id = $3, relationship_type_id = $4, description = $5, updated_at = $6
		WHERE id = $1
	`

	result, err := q.db.ExecContext(ctx, query,
		tagRelationship.ID,
		tagRelationship.TagAID,
		tagRelationship.TagBID,
		tagRelationship.RelationshipTypeID,
		tagRelationship.Description,
		tagRelationship.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to update tag relationship: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("tag relationship not found")
	}

	return nil
}

// Delete removes a tag relationship by ID
func (q *TagRelationshipQueries) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM tag_relationships WHERE id = $1`

	result, err := q.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete tag relationship: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("tag relationship not found")
	}

	return nil
}
