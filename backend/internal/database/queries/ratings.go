package queries

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/kmahabeer/tag-management-system/backend/internal/models"
)

// RatingQueries implements RatingRepository
type RatingQueries struct {
	db *sql.DB
}

// NewRatingQueries creates a new RatingQueries instance
func NewRatingQueries(db *sql.DB) *RatingQueries {
	return &RatingQueries{db: db}
}

// GetByID retrieves a rating by its ID
func (q *RatingQueries) GetByID(ctx context.Context, id uuid.UUID) (*models.Rating, error) {
	query := `
		SELECT id, name, score, description, rating_type_id, created_at, updated_at
		FROM ratings
		WHERE id = $1
	`

	var rating models.Rating
	err := q.db.QueryRowContext(ctx, query, id).Scan(
		&rating.ID,
		&rating.Name,
		&rating.Score,
		&rating.Description,
		&rating.RatingTypeID,
		&rating.CreatedAt,
		&rating.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("rating not found")
		}
		return nil, fmt.Errorf("failed to get rating: %w", err)
	}

	return &rating, nil
}

// List retrieves a list of ratings with pagination
func (q *RatingQueries) List(ctx context.Context, limit, offset int) ([]*models.Rating, error) {
	query := `
		SELECT id, name, score, description, rating_type_id, created_at, updated_at
		FROM ratings
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := q.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list ratings: %w", err)
	}
	defer rows.Close()

	var ratings []*models.Rating
	for rows.Next() {
		var rating models.Rating
		err := rows.Scan(
			&rating.ID,
			&rating.Name,
			&rating.Score,
			&rating.Description,
			&rating.RatingTypeID,
			&rating.CreatedAt,
			&rating.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan rating: %w", err)
		}
		ratings = append(ratings, &rating)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return ratings, nil
}

// Create inserts a new rating
func (q *RatingQueries) Create(ctx context.Context, rating *models.Rating) error {
	query := `
		INSERT INTO ratings (id, name, score, description, rating_type_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := q.db.ExecContext(ctx, query,
		rating.ID,
		rating.Name,
		rating.Score,
		rating.Description,
		rating.RatingTypeID,
		rating.CreatedAt,
		rating.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create rating: %w", err)
	}

	return nil
}

// Update modifies an existing rating
func (q *RatingQueries) Update(ctx context.Context, rating *models.Rating) error {
	query := `
		UPDATE ratings
		SET name = $2, score = $3, description = $4, rating_type_id = $5, updated_at = $6
		WHERE id = $1
	`

	result, err := q.db.ExecContext(ctx, query,
		rating.ID,
		rating.Name,
		rating.Score,
		rating.Description,
		rating.RatingTypeID,
		rating.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to update rating: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("rating not found")
	}

	return nil
}

// Delete removes a rating by ID
func (q *RatingQueries) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM ratings WHERE id = $1`

	result, err := q.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete rating: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("rating not found")
	}

	return nil
}
