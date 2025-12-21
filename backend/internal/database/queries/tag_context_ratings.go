package queries

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/kmahabeer/tag-management-system/backend/internal/models"
)

// TagContextRatingQueries implements TagContextRatingRepository
type TagContextRatingQueries struct {
	db *sql.DB
}

// NewTagContextRatingQueries creates a new TagContextRatingQueries instance
func NewTagContextRatingQueries(db *sql.DB) *TagContextRatingQueries {
	return &TagContextRatingQueries{db: db}
}

// GetByID retrieves a tag context rating by its ID
func (q *TagContextRatingQueries) GetByID(ctx context.Context, id uuid.UUID) (*models.TagContextRating, error) {
	query := `
		SELECT id, tag_id, context_id, rating_id, user_id, created_at, updated_at
		FROM tag_context_ratings
		WHERE id = $1
	`

	var tagContextRating models.TagContextRating
	var userID sql.NullString
	err := q.db.QueryRowContext(ctx, query, id).Scan(
		&tagContextRating.ID,
		&tagContextRating.TagID,
		&tagContextRating.ContextID,
		&tagContextRating.RatingID,
		&userID,
		&tagContextRating.CreatedAt,
		&tagContextRating.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("tag context rating not found")
		}
		return nil, fmt.Errorf("failed to get tag context rating: %w", err)
	}

	if userID.Valid {
		parsedID, err := uuid.Parse(userID.String)
		if err != nil {
			return nil, fmt.Errorf("invalid user_id: %w", err)
		}
		tagContextRating.UserID = &parsedID
	} else {
		tagContextRating.UserID = nil
	}

	return &tagContextRating, nil
}

// List retrieves a list of tag context ratings with pagination
func (q *TagContextRatingQueries) List(ctx context.Context, limit, offset int) ([]*models.TagContextRating, error) {
	query := `
		SELECT id, tag_id, context_id, rating_id, user_id, created_at, updated_at
		FROM tag_context_ratings
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := q.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list tag context ratings: %w", err)
	}
	defer rows.Close()

	var tagContextRatings []*models.TagContextRating
	for rows.Next() {
		var tagContextRating models.TagContextRating
		var userID sql.NullString
		err := rows.Scan(
			&tagContextRating.ID,
			&tagContextRating.TagID,
			&tagContextRating.ContextID,
			&tagContextRating.RatingID,
			&userID,
			&tagContextRating.CreatedAt,
			&tagContextRating.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan tag context rating: %w", err)
		}

		if userID.Valid {
			parsedID, err := uuid.Parse(userID.String)
			if err != nil {
				return nil, fmt.Errorf("invalid user_id: %w", err)
			}
			tagContextRating.UserID = &parsedID
		} else {
			tagContextRating.UserID = nil
		}
		tagContextRatings = append(tagContextRatings, &tagContextRating)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return tagContextRatings, nil
}

// Create inserts a new tag context rating
func (q *TagContextRatingQueries) Create(ctx context.Context, tagContextRating *models.TagContextRating) error {
	query := `
		INSERT INTO tag_context_ratings (id, tag_id, context_id, rating_id, user_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	var userID interface{}
	if tagContextRating.UserID != nil {
		userID = tagContextRating.UserID.String()
	} else {
		userID = nil
	}

	_, err := q.db.ExecContext(ctx, query,
		tagContextRating.ID,
		tagContextRating.TagID,
		tagContextRating.ContextID,
		tagContextRating.RatingID,
		userID,
		tagContextRating.CreatedAt,
		tagContextRating.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create tag context rating: %w", err)
	}

	return nil
}

// Update modifies an existing tag context rating
func (q *TagContextRatingQueries) Update(ctx context.Context, tagContextRating *models.TagContextRating) error {
	query := `
		UPDATE tag_context_ratings
		SET tag_id = $2, context_id = $3, rating_id = $4, user_id = $5, updated_at = $6
		WHERE id = $1
	`

	var userID interface{}
	if tagContextRating.UserID != nil {
		userID = tagContextRating.UserID.String()
	} else {
		userID = nil
	}

	result, err := q.db.ExecContext(ctx, query,
		tagContextRating.ID,
		tagContextRating.TagID,
		tagContextRating.ContextID,
		tagContextRating.RatingID,
		userID,
		tagContextRating.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to update tag context rating: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("tag context rating not found")
	}

	return nil
}

// Delete removes a tag context rating by ID
func (q *TagContextRatingQueries) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM tag_context_ratings WHERE id = $1`

	result, err := q.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete tag context rating: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("tag context rating not found")
	}

	return nil
}
