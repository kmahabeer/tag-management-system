package queries

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/kmahabeer/tag-management-system/backend/internal/models"
)

// RatingTypeQueries implements RatingTypeRepository
type RatingTypeQueries struct {
	db *sql.DB
}

// NewRatingTypeQueries creates a new RatingTypeQueries instance
func NewRatingTypeQueries(db *sql.DB) *RatingTypeQueries {
	return &RatingTypeQueries{db: db}
}

// GetByID retrieves a rating type by its ID
func (q *RatingTypeQueries) GetByID(ctx context.Context, id uuid.UUID) (*models.RatingType, error) {
	query := `
		SELECT id, name, is_normalized, created_at, updated_at
		FROM rating_types
		WHERE id = $1
	`

	var ratingType models.RatingType
	err := q.db.QueryRowContext(ctx, query, id).Scan(
		&ratingType.ID,
		&ratingType.Name,
		&ratingType.IsNormalized,
		&ratingType.CreatedAt,
		&ratingType.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("rating type not found")
		}
		return nil, fmt.Errorf("failed to get rating type: %w", err)
	}

	return &ratingType, nil
}

// List retrieves a list of rating types with pagination
func (q *RatingTypeQueries) List(ctx context.Context, limit, offset int) ([]*models.RatingType, error) {
	query := `
		SELECT id, name, is_normalized, created_at, updated_at
		FROM rating_types
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := q.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list rating types: %w", err)
	}
	defer rows.Close()

	var ratingTypes []*models.RatingType
	for rows.Next() {
		var ratingType models.RatingType
		err := rows.Scan(
			&ratingType.ID,
			&ratingType.Name,
			&ratingType.IsNormalized,
			&ratingType.CreatedAt,
			&ratingType.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan rating type: %w", err)
		}
		ratingTypes = append(ratingTypes, &ratingType)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return ratingTypes, nil
}

// Create inserts a new rating type
func (q *RatingTypeQueries) Create(ctx context.Context, ratingType *models.RatingType) error {
	query := `
		INSERT INTO rating_types (id, name, is_normalized, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := q.db.ExecContext(ctx, query,
		ratingType.ID,
		ratingType.Name,
		ratingType.IsNormalized,
		ratingType.CreatedAt,
		ratingType.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create rating type: %w", err)
	}

	return nil
}

// Update modifies an existing rating type
func (q *RatingTypeQueries) Update(ctx context.Context, ratingType *models.RatingType) error {
	query := `
		UPDATE rating_types
		SET name = $2, is_normalized = $3, updated_at = $4
		WHERE id = $1
	`

	result, err := q.db.ExecContext(ctx, query,
		ratingType.ID,
		ratingType.Name,
		ratingType.IsNormalized,
		ratingType.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to update rating type: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("rating type not found")
	}

	return nil
}

// Delete removes a rating type by ID
func (q *RatingTypeQueries) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM rating_types WHERE id = $1`

	result, err := q.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete rating type: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("rating type not found")
	}

	return nil
}
