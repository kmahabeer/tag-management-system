package queries

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/kmahabeer/tag-management-system/backend/internal/models"
)

// ContextQueries implements ContextRepository
type ContextQueries struct {
	db *sql.DB
}

// NewContextQueries creates a new ContextQueries instance
func NewContextQueries(db *sql.DB) *ContextQueries {
	return &ContextQueries{db: db}
}

// GetByID retrieves a context by its ID
func (q *ContextQueries) GetByID(ctx context.Context, id uuid.UUID) (*models.Context, error) {
	query := `
		SELECT id, name, classification_type, description, is_active, created_at, updated_at
		FROM contexts
		WHERE id = $1
	`

	var context models.Context
	err := q.db.QueryRowContext(ctx, query, id).Scan(
		&context.ID,
		&context.Name,
		&context.ClassificationType,
		&context.Description,
		&context.IsActive,
		&context.CreatedAt,
		&context.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("context not found")
		}
		return nil, fmt.Errorf("failed to get context: %w", err)
	}

	return &context, nil
}

// List retrieves a list of contexts with pagination
func (q *ContextQueries) List(ctx context.Context, limit, offset int) ([]*models.Context, error) {
	query := `
		SELECT id, name, classification_type, description, is_active, created_at, updated_at
		FROM contexts
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := q.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list contexts: %w", err)
	}
	defer rows.Close()

	var contexts []*models.Context
	for rows.Next() {
		var context models.Context
		err := rows.Scan(
			&context.ID,
			&context.Name,
			&context.ClassificationType,
			&context.Description,
			&context.IsActive,
			&context.CreatedAt,
			&context.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan context: %w", err)
		}
		contexts = append(contexts, &context)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return contexts, nil
}

// Create inserts a new context
func (q *ContextQueries) Create(ctx context.Context, context *models.Context) error {
	query := `
		INSERT INTO contexts (id, name, classification_type, description, is_active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := q.db.ExecContext(ctx, query,
		context.ID,
		context.Name,
		context.ClassificationType,
		context.Description,
		context.IsActive,
		context.CreatedAt,
		context.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create context: %w", err)
	}

	return nil
}

// Update modifies an existing context
func (q *ContextQueries) Update(ctx context.Context, context *models.Context) error {
	query := `
		UPDATE contexts
		SET name = $2, classification_type = $3, description = $4, is_active = $5, updated_at = $6
		WHERE id = $1
	`

	result, err := q.db.ExecContext(ctx, query,
		context.ID,
		context.Name,
		context.ClassificationType,
		context.Description,
		context.IsActive,
		context.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to update context: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("context not found")
	}

	return nil
}

// Delete removes a context by ID
func (q *ContextQueries) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM contexts WHERE id = $1`

	result, err := q.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete context: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("context not found")
	}

	return nil
}
