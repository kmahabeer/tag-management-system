package queries

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/kmahabeer/tag-management-system/backend/internal/models"
)

// EntityRelationshipTypeQueries implements EntityRelationshipTypeRepository
type EntityRelationshipTypeQueries struct {
	db *sql.DB
}

// NewEntityRelationshipTypeQueries creates a new EntityRelationshipTypeQueries instance
func NewEntityRelationshipTypeQueries(db *sql.DB) *EntityRelationshipTypeQueries {
	return &EntityRelationshipTypeQueries{db: db}
}

// GetByID retrieves an entity relationship type by its ID
func (q *EntityRelationshipTypeQueries) GetByID(ctx context.Context, id uuid.UUID) (*models.EntityRelationshipType, error) {
	query := `
		SELECT id, name, created_at, updated_at
		FROM entity_relationship_types
		WHERE id = $1
	`

	var entityRelationshipType models.EntityRelationshipType
	err := q.db.QueryRowContext(ctx, query, id).Scan(
		&entityRelationshipType.ID,
		&entityRelationshipType.Name,
		&entityRelationshipType.CreatedAt,
		&entityRelationshipType.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("entity relationship type not found")
		}
		return nil, fmt.Errorf("failed to get entity relationship type: %w", err)
	}

	return &entityRelationshipType, nil
}

// List retrieves a list of entity relationship types with pagination
func (q *EntityRelationshipTypeQueries) List(ctx context.Context, limit, offset int) ([]*models.EntityRelationshipType, error) {
	query := `
		SELECT id, name, created_at, updated_at
		FROM entity_relationship_types
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := q.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list entity relationship types: %w", err)
	}
	defer rows.Close()

	var entityRelationshipTypes []*models.EntityRelationshipType
	for rows.Next() {
		var entityRelationshipType models.EntityRelationshipType
		err := rows.Scan(
			&entityRelationshipType.ID,
			&entityRelationshipType.Name,
			&entityRelationshipType.CreatedAt,
			&entityRelationshipType.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan entity relationship type: %w", err)
		}
		entityRelationshipTypes = append(entityRelationshipTypes, &entityRelationshipType)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return entityRelationshipTypes, nil
}

// Create inserts a new entity relationship type
func (q *EntityRelationshipTypeQueries) Create(ctx context.Context, entityRelationshipType *models.EntityRelationshipType) error {
	query := `
		INSERT INTO entity_relationship_types (id, name, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
	`

	_, err := q.db.ExecContext(ctx, query,
		entityRelationshipType.ID,
		entityRelationshipType.Name,
		entityRelationshipType.CreatedAt,
		entityRelationshipType.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create entity relationship type: %w", err)
	}

	return nil
}

// Update modifies an existing entity relationship type
func (q *EntityRelationshipTypeQueries) Update(ctx context.Context, entityRelationshipType *models.EntityRelationshipType) error {
	query := `
		UPDATE entity_relationship_types
		SET name = $2, updated_at = $3
		WHERE id = $1
	`

	result, err := q.db.ExecContext(ctx, query,
		entityRelationshipType.ID,
		entityRelationshipType.Name,
		entityRelationshipType.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to update entity relationship type: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("entity relationship type not found")
	}

	return nil
}

// Delete removes an entity relationship type by ID
func (q *EntityRelationshipTypeQueries) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM entity_relationship_types WHERE id = $1`

	result, err := q.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete entity relationship type: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("entity relationship type not found")
	}

	return nil
}
