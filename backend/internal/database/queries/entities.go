package queries

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/kmahabeer/tag-management-system/backend/internal/models"
)

// EntityQueries implements EntityRepository
type EntityQueries struct {
	db *sql.DB
}

// NewEntityQueries creates a new EntityQueries instance
func NewEntityQueries(db *sql.DB) *EntityQueries {
	return &EntityQueries{db: db}
}

// GetByID retrieves an entity by its ID
func (q *EntityQueries) GetByID(ctx context.Context, id uuid.UUID) (*models.Entity, error) {
	query := `
		SELECT id, name, location, is_primary, metadata, created_at, updated_at
		FROM entities
		WHERE id = $1
	`

	var entity models.Entity
	err := q.db.QueryRowContext(ctx, query, id).Scan(
		&entity.ID,
		&entity.Name,
		&entity.Location,
		&entity.IsPrimary,
		&entity.Metadata,
		&entity.CreatedAt,
		&entity.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("entity not found")
		}
		return nil, fmt.Errorf("failed to get entity: %w", err)
	}

	return &entity, nil
}

// List retrieves a list of entities with pagination
func (q *EntityQueries) List(ctx context.Context, limit, offset int) ([]*models.Entity, error) {
	query := `
		SELECT id, name, location, is_primary, metadata, created_at, updated_at
		FROM entities
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := q.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list entities: %w", err)
	}
	defer rows.Close()

	var entities []*models.Entity
	for rows.Next() {
		var entity models.Entity
		err := rows.Scan(
			&entity.ID,
			&entity.Name,
			&entity.Location,
			&entity.IsPrimary,
			&entity.Metadata,
			&entity.CreatedAt,
			&entity.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan entity: %w", err)
		}
		entities = append(entities, &entity)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return entities, nil
}

// Create inserts a new entity
func (q *EntityQueries) Create(ctx context.Context, entity *models.Entity) error {
	query := `
		INSERT INTO entities (id, name, location, is_primary, metadata, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := q.db.ExecContext(ctx, query,
		entity.ID,
		entity.Name,
		entity.Location,
		entity.IsPrimary,
		entity.Metadata,
		entity.CreatedAt,
		entity.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create entity: %w", err)
	}

	return nil
}

// Update modifies an existing entity
func (q *EntityQueries) Update(ctx context.Context, entity *models.Entity) error {
	query := `
		UPDATE entities
		SET name = $2, location = $3, is_primary = $4, metadata = $5, updated_at = $6
		WHERE id = $1
	`

	result, err := q.db.ExecContext(ctx, query,
		entity.ID,
		entity.Name,
		entity.Location,
		entity.IsPrimary,
		entity.Metadata,
		entity.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to update entity: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("entity not found")
	}

	return nil
}

// Delete removes an entity by ID
func (q *EntityQueries) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM entities WHERE id = $1`

	result, err := q.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete entity: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("entity not found")
	}

	return nil
}
