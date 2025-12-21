package queries

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/kmahabeer/tag-management-system/backend/internal/models"
)

// EntityPurposeQueries implements EntityPurposeRepository
type EntityPurposeQueries struct {
	db *sql.DB
}

// NewEntityPurposeQueries creates a new EntityPurposeQueries instance
func NewEntityPurposeQueries(db *sql.DB) *EntityPurposeQueries {
	return &EntityPurposeQueries{db: db}
}

// GetByID retrieves an entity purpose by its ID
func (q *EntityPurposeQueries) GetByID(ctx context.Context, id uuid.UUID) (*models.EntityPurpose, error) {
	query := `
		SELECT id, entity_id, purpose_tag_id, is_primary, created_at, updated_at
		FROM entity_purposes
		WHERE id = $1
	`

	var entityPurpose models.EntityPurpose
	err := q.db.QueryRowContext(ctx, query, id).Scan(
		&entityPurpose.ID,
		&entityPurpose.EntityID,
		&entityPurpose.PurposeTagID,
		&entityPurpose.IsPrimary,
		&entityPurpose.CreatedAt,
		&entityPurpose.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("entity purpose not found")
		}
		return nil, fmt.Errorf("failed to get entity purpose: %w", err)
	}

	return &entityPurpose, nil
}

// List retrieves a list of entity purposes with pagination
func (q *EntityPurposeQueries) List(ctx context.Context, limit, offset int) ([]*models.EntityPurpose, error) {
	query := `
		SELECT id, entity_id, purpose_tag_id, is_primary, created_at, updated_at
		FROM entity_purposes
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := q.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list entity purposes: %w", err)
	}
	defer rows.Close()

	var entityPurposes []*models.EntityPurpose
	for rows.Next() {
		var entityPurpose models.EntityPurpose
		err := rows.Scan(
			&entityPurpose.ID,
			&entityPurpose.EntityID,
			&entityPurpose.PurposeTagID,
			&entityPurpose.IsPrimary,
			&entityPurpose.CreatedAt,
			&entityPurpose.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan entity purpose: %w", err)
		}
		entityPurposes = append(entityPurposes, &entityPurpose)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return entityPurposes, nil
}

// Create inserts a new entity purpose
func (q *EntityPurposeQueries) Create(ctx context.Context, entityPurpose *models.EntityPurpose) error {
	query := `
		INSERT INTO entity_purposes (id, entity_id, purpose_tag_id, is_primary, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := q.db.ExecContext(ctx, query,
		entityPurpose.ID,
		entityPurpose.EntityID,
		entityPurpose.PurposeTagID,
		entityPurpose.IsPrimary,
		entityPurpose.CreatedAt,
		entityPurpose.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create entity purpose: %w", err)
	}

	return nil
}

// Update modifies an existing entity purpose
func (q *EntityPurposeQueries) Update(ctx context.Context, entityPurpose *models.EntityPurpose) error {
	query := `
		UPDATE entity_purposes
		SET entity_id = $2, purpose_tag_id = $3, is_primary = $4, updated_at = $5
		WHERE id = $1
	`

	result, err := q.db.ExecContext(ctx, query,
		entityPurpose.ID,
		entityPurpose.EntityID,
		entityPurpose.PurposeTagID,
		entityPurpose.IsPrimary,
		entityPurpose.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to update entity purpose: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("entity purpose not found")
	}

	return nil
}

// Delete removes an entity purpose by ID
func (q *EntityPurposeQueries) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM entity_purposes WHERE id = $1`

	result, err := q.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete entity purpose: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("entity purpose not found")
	}

	return nil
}
