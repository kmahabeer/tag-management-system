package queries

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/kmahabeer/tag-management-system/backend/internal/models"
)

// UIGroupQueries implements UIGroupRepository
type UIGroupQueries struct {
	db *sql.DB
}

// NewUIGroupQueries creates a new UIGroupQueries instance
func NewUIGroupQueries(db *sql.DB) *UIGroupQueries {
	return &UIGroupQueries{db: db}
}

// GetByID retrieves a UI group by its ID
func (q *UIGroupQueries) GetByID(ctx context.Context, id uuid.UUID) (*models.UIGroup, error) {
	query := `
		SELECT id, name, created_at, updated_at
		FROM ui_groups
		WHERE id = $1
	`

	var uiGroup models.UIGroup
	err := q.db.QueryRowContext(ctx, query, id).Scan(
		&uiGroup.ID,
		&uiGroup.Name,
		&uiGroup.CreatedAt,
		&uiGroup.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("UI group not found")
		}
		return nil, fmt.Errorf("failed to get UI group: %w", err)
	}

	return &uiGroup, nil
}

// List retrieves a list of UI groups with pagination
func (q *UIGroupQueries) List(ctx context.Context, limit, offset int) ([]*models.UIGroup, error) {
	query := `
		SELECT id, name, created_at, updated_at
		FROM ui_groups
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := q.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list UI groups: %w", err)
	}
	defer rows.Close()

	var uiGroups []*models.UIGroup
	for rows.Next() {
		var uiGroup models.UIGroup
		err := rows.Scan(
			&uiGroup.ID,
			&uiGroup.Name,
			&uiGroup.CreatedAt,
			&uiGroup.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan UI group: %w", err)
		}
		uiGroups = append(uiGroups, &uiGroup)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return uiGroups, nil
}

// Create inserts a new UI group
func (q *UIGroupQueries) Create(ctx context.Context, uiGroup *models.UIGroup) error {
	query := `
		INSERT INTO ui_groups (id, name, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
	`

	_, err := q.db.ExecContext(ctx, query,
		uiGroup.ID,
		uiGroup.Name,
		uiGroup.CreatedAt,
		uiGroup.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create UI group: %w", err)
	}

	return nil
}

// Update modifies an existing UI group
func (q *UIGroupQueries) Update(ctx context.Context, uiGroup *models.UIGroup) error {
	query := `
		UPDATE ui_groups
		SET name = $2, updated_at = $3
		WHERE id = $1
	`

	result, err := q.db.ExecContext(ctx, query,
		uiGroup.ID,
		uiGroup.Name,
		uiGroup.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to update UI group: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("UI group not found")
	}

	return nil
}

// Delete removes a UI group by ID
func (q *UIGroupQueries) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM ui_groups WHERE id = $1`

	result, err := q.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete UI group: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("UI group not found")
	}

	return nil
}
