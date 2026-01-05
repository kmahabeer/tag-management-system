package queries

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/kmahabeer/tag-management-system/backend/internal/models"
)

// TagRelationshipRatingQueries implements TagRelationshipRatingRepository
type TagRelationshipRatingQueries struct {
	db *sql.DB
}

// NewTagRelationshipRatingQueries creates a new TagRelationshipRatingQueries instance
func NewTagRelationshipRatingQueries(db *sql.DB) *TagRelationshipRatingQueries {
	return &TagRelationshipRatingQueries{db: db}
}

// GetByID retrieves a tag relationship rating by its ID
func (q *TagRelationshipRatingQueries) GetByID(ctx context.Context, id uuid.UUID) (*models.TagRelationshipRating, error) {
	query := `
		SELECT id, tag_a_id, tag_b_id, context_id, rating_id, created_at, updated_at
		FROM tag_relationship_ratings
		WHERE id = $1
	`

	var tagRelationshipRating models.TagRelationshipRating
	err := q.db.QueryRowContext(ctx, query, id).Scan(
		&tagRelationshipRating.ID,
		&tagRelationshipRating.TagAID,
		&tagRelationshipRating.TagBID,
		&tagRelationshipRating.ContextID,
		&tagRelationshipRating.RatingID,
		&tagRelationshipRating.CreatedAt,
		&tagRelationshipRating.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("tag relationship rating not found")
		}
		return nil, fmt.Errorf("failed to get tag relationship rating: %w", err)
	}

	return &tagRelationshipRating, nil
}

// List retrieves a list of tag relationship ratings with pagination
func (q *TagRelationshipRatingQueries) List(ctx context.Context, limit, offset int) ([]*models.TagRelationshipRating, error) {
	query := `
		SELECT id, tag_a_id, tag_b_id, context_id, rating_id, created_at, updated_at
		FROM tag_relationship_ratings
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := q.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list tag relationship ratings: %w", err)
	}
	defer rows.Close()

	var tagRelationshipRatings []*models.TagRelationshipRating
	for rows.Next() {
		var tagRelationshipRating models.TagRelationshipRating
		err := rows.Scan(
			&tagRelationshipRating.ID,
			&tagRelationshipRating.TagAID,
			&tagRelationshipRating.TagBID,
			&tagRelationshipRating.ContextID,
			&tagRelationshipRating.RatingID,
			&tagRelationshipRating.CreatedAt,
			&tagRelationshipRating.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan tag relationship rating: %w", err)
		}
		tagRelationshipRatings = append(tagRelationshipRatings, &tagRelationshipRating)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return tagRelationshipRatings, nil
}

// Create inserts a new tag relationship rating
func (q *TagRelationshipRatingQueries) Create(ctx context.Context, tagRelationshipRating *models.TagRelationshipRating) error {
	query := `
		INSERT INTO tag_relationship_ratings (id, tag_a_id, tag_b_id, context_id, rating_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := q.db.ExecContext(ctx, query,
		tagRelationshipRating.ID,
		tagRelationshipRating.TagAID,
		tagRelationshipRating.TagBID,
		tagRelationshipRating.ContextID,
		tagRelationshipRating.RatingID,
		tagRelationshipRating.CreatedAt,
		tagRelationshipRating.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create tag relationship rating: %w", err)
	}

	return nil
}

// Update modifies an existing tag relationship rating
func (q *TagRelationshipRatingQueries) Update(ctx context.Context, tagRelationshipRating *models.TagRelationshipRating) error {
	query := `
		UPDATE tag_relationship_ratings
		SET tag_a_id = $2, tag_b_id = $3, context_id = $4, rating_id = $5, updated_at = $6
		WHERE id = $1
	`

	result, err := q.db.ExecContext(ctx, query,
		tagRelationshipRating.ID,
		tagRelationshipRating.TagAID,
		tagRelationshipRating.TagBID,
		tagRelationshipRating.ContextID,
		tagRelationshipRating.RatingID,
		tagRelationshipRating.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to update tag relationship rating: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("tag relationship rating not found")
	}

	return nil
}

// Delete removes a tag relationship rating by ID
func (q *TagRelationshipRatingQueries) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM tag_relationship_ratings WHERE id = $1`

	result, err := q.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete tag relationship rating: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("tag relationship rating not found")
	}

	return nil
}
