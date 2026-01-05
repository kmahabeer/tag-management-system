package queries

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/kmahabeer/tag-management-system/backend/internal/models"
)

// EntityRelationshipQueries implements EntityRelationshipRepository
type EntityRelationshipQueries struct {
	db *sql.DB
}

// NewEntityRelationshipQueries creates a new EntityRelationshipQueries instance
func NewEntityRelationshipQueries(db *sql.DB) *EntityRelationshipQueries {
	return &EntityRelationshipQueries{db: db}
}

// GetByID retrieves an entity relationship by its ID
func (q *EntityRelationshipQueries) GetByID(ctx context.Context, id uuid.UUID) (*models.EntityRelationship, error) {
	query := `
		SELECT id, entity_a_id, entity_b_id, relationship_type_id, created_at, updated_at
		FROM entity_relationships
		WHERE id = $1
	`

	var entityRelationship models.EntityRelationship
	err := q.db.QueryRowContext(ctx, query, id).Scan(
		&entityRelationship.ID,
		&entityRelationship.EntityAID,
		&entityRelationship.EntityBID,
		&entityRelationship.RelationshipTypeID,
		&entityRelationship.CreatedAt,
		&entityRelationship.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("entity relationship not found")
		}
		return nil, fmt.Errorf("failed to get entity relationship: %w", err)
	}

	return &entityRelationship, nil
}

// List retrieves a list of entity relationships with pagination
func (q *EntityRelationshipQueries) List(ctx context.Context, limit, offset int) ([]*models.EntityRelationship, error) {
	query := `
		SELECT id, entity_a_id, entity_b_id, relationship_type_id, created_at, updated_at
		FROM entity_relationships
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := q.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list entity relationships: %w", err)
	}
	defer rows.Close()

	var entityRelationships []*models.EntityRelationship
	for rows.Next() {
		var entityRelationship models.EntityRelationship
		err := rows.Scan(
			&entityRelationship.ID,
			&entityRelationship.EntityAID,
			&entityRelationship.EntityBID,
			&entityRelationship.RelationshipTypeID,
			&entityRelationship.CreatedAt,
			&entityRelationship.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan entity relationship: %w", err)
		}
		entityRelationships = append(entityRelationships, &entityRelationship)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return entityRelationships, nil
}

// Create inserts a new entity relationship
func (q *EntityRelationshipQueries) Create(ctx context.Context, entityRelationship *models.EntityRelationship) error {
	query := `
		INSERT INTO entity_relationships (id, entity_a_id, entity_b_id, relationship_type_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := q.db.ExecContext(ctx, query,
		entityRelationship.ID,
		entityRelationship.EntityAID,
		entityRelationship.EntityBID,
		entityRelationship.RelationshipTypeID,
		entityRelationship.CreatedAt,
		entityRelationship.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create entity relationship: %w", err)
	}

	return nil
}

// Update modifies an existing entity relationship
func (q *EntityRelationshipQueries) Update(ctx context.Context, entityRelationship *models.EntityRelationship) error {
	query := `
		UPDATE entity_relationships
		SET entity_a_id = $2, entity_b_id = $3, relationship_type_id = $4, updated_at = $5
		WHERE id = $1
	`

	result, err := q.db.ExecContext(ctx, query,
		entityRelationship.ID,
		entityRelationship.EntityAID,
		entityRelationship.EntityBID,
		entityRelationship.RelationshipTypeID,
		entityRelationship.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to update entity relationship: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("entity relationship not found")
	}

	return nil
}

// Delete removes an entity relationship by ID
func (q *EntityRelationshipQueries) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM entity_relationships WHERE id = $1`

	result, err := q.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete entity relationship: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("entity relationship not found")
	}

	return nil
}
