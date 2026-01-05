package queries

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/kmahabeer/tag-management-system/backend/internal/models"
)

// EntityRelationshipRatingQueries implements EntityRelationshipRatingRepository
type EntityRelationshipRatingQueries struct {
	db *sql.DB
}

// NewEntityRelationshipRatingQueries creates a new EntityRelationshipRatingQueries instance
func NewEntityRelationshipRatingQueries(db *sql.DB) *EntityRelationshipRatingQueries {
	return &EntityRelationshipRatingQueries{db: db}
}

// GetByID retrieves an entity relationship rating by its ID
func (q *EntityRelationshipRatingQueries) GetByID(ctx context.Context, id uuid.UUID) (*models.EntityRelationshipRating, error) {
	query := `
		SELECT id, entity_a_id, entity_b_id, context_id, rating_id, created_at, updated_at
		FROM entity_relationship_ratings
		WHERE id = $1
	`

	var entityRelationshipRating models.EntityRelationshipRating
	err := q.db.QueryRowContext(ctx, query, id).Scan(
		&entityRelationshipRating.ID,
		&entityRelationshipRating.EntityAID,
		&entityRelationshipRating.EntityBID,
		&entityRelationshipRating.ContextID,
		&entityRelationshipRating.RatingID,
		&entityRelationshipRating.CreatedAt,
		&entityRelationshipRating.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("entity relationship rating not found")
		}
		return nil, fmt.Errorf("failed to get entity relationship rating: %w", err)
	}

	return &entityRelationshipRating, nil
}

// List retrieves a list of entity relationship ratings with pagination
func (q *EntityRelationshipRatingQueries) List(ctx context.Context, limit, offset int) ([]*models.EntityRelationshipRating, error) {
	query := `
		SELECT id, entity_a_id, entity_b_id, context_id, rating_id, created_at, updated_at
		FROM entity_relationship_ratings
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := q.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list entity relationship ratings: %w", err)
	}
	defer rows.Close()

	var entityRelationshipRatings []*models.EntityRelationshipRating
	for rows.Next() {
		var entityRelationshipRating models.EntityRelationshipRating
		err := rows.Scan(
			&entityRelationshipRating.ID,
			&entityRelationshipRating.EntityAID,
			&entityRelationshipRating.EntityBID,
			&entityRelationshipRating.ContextID,
			&entityRelationshipRating.RatingID,
			&entityRelationshipRating.CreatedAt,
			&entityRelationshipRating.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan entity relationship rating: %w", err)
		}
		entityRelationshipRatings = append(entityRelationshipRatings, &entityRelationshipRating)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return entityRelationshipRatings, nil
}

// Create inserts a new entity relationship rating
func (q *EntityRelationshipRatingQueries) Create(ctx context.Context, entityRelationshipRating *models.EntityRelationshipRating) error {
	query := `
		INSERT INTO entity_relationship_ratings (id, entity_a_id, entity_b_id, context_id, rating_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := q.db.ExecContext(ctx, query,
		entityRelationshipRating.ID,
		entityRelationshipRating.EntityAID,
		entityRelationshipRating.EntityBID,
		entityRelationshipRating.ContextID,
		entityRelationshipRating.RatingID,
		entityRelationshipRating.CreatedAt,
		entityRelationshipRating.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create entity relationship rating: %w", err)
	}

	return nil
}

// Update modifies an existing entity relationship rating
func (q *EntityRelationshipRatingQueries) Update(ctx context.Context, entityRelationshipRating *models.EntityRelationshipRating) error {
	query := `
		UPDATE entity_relationship_ratings
		SET entity_a_id = $2, entity_b_id = $3, context_id = $4, rating_id = $5, updated_at = $6
		WHERE id = $1
	`

	result, err := q.db.ExecContext(ctx, query,
		entityRelationshipRating.ID,
		entityRelationshipRating.EntityAID,
		entityRelationshipRating.EntityBID,
		entityRelationshipRating.ContextID,
		entityRelationshipRating.RatingID,
		entityRelationshipRating.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to update entity relationship rating: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("entity relationship rating not found")
	}

	return nil
}

// Delete removes an entity relationship rating by ID
func (q *EntityRelationshipRatingQueries) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM entity_relationship_ratings WHERE id = $1`

	result, err := q.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete entity relationship rating: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("entity relationship rating not found")
	}

	return nil
}
