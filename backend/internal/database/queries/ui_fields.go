package queries

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/kmahabeer/tag-management-system/backend/internal/models"
)

// UIFieldQueries implements UIFieldRepository
type UIFieldQueries struct {
	db *sql.DB
}

// NewUIFieldQueries creates a new UIFieldQueries instance
func NewUIFieldQueries(db *sql.DB) *UIFieldQueries {
	return &UIFieldQueries{db: db}
}

// GetByID retrieves a UI field by its ID
func (q *UIFieldQueries) GetByID(ctx context.Context, id uuid.UUID) (*models.UIField, error) {
	query := `
		SELECT id, ui_layout_id, ui_group_id, context_id, category_tag_id, sort_order, created_at, updated_at
		FROM ui_fields
		WHERE id = $1
	`

	var uiField models.UIField
	err := q.db.QueryRowContext(ctx, query, id).Scan(
		&uiField.ID,
		&uiField.UILayoutID,
		&uiField.UIGroupID,
		&uiField.ContextID,
		&uiField.CategoryTagID,
		&uiField.SortOrder,
		&uiField.CreatedAt,
		&uiField.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("UI field not found")
		}
		return nil, fmt.Errorf("failed to get UI field: %w", err)
	}

	return &uiField, nil
}

// List retrieves a list of UI fields with pagination
func (q *UIFieldQueries) List(ctx context.Context, limit, offset int) ([]*models.UIField, error) {
	query := `
		SELECT id, ui_layout_id, ui_group_id, context_id, category_tag_id, sort_order, created_at, updated_at
		FROM ui_fields
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := q.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list UI fields: %w", err)
	}
	defer rows.Close()

	var uiFields []*models.UIField
	for rows.Next() {
		var uiField models.UIField
		err := rows.Scan(
			&uiField.ID,
			&uiField.UILayoutID,
			&uiField.UIGroupID,
			&uiField.ContextID,
			&uiField.CategoryTagID,
			&uiField.SortOrder,
			&uiField.CreatedAt,
			&uiField.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan UI field: %w", err)
		}
		uiFields = append(uiFields, &uiField)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return uiFields, nil
}

// Create inserts a new UI field
func (q *UIFieldQueries) Create(ctx context.Context, uiField *models.UIField) error {
	query := `
		INSERT INTO ui_fields (id, ui_layout_id, ui_group_id, context_id, category_tag_id, sort_order, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := q.db.ExecContext(ctx, query,
		uiField.ID,
		uiField.UILayoutID,
		uiField.UIGroupID,
		uiField.ContextID,
		uiField.CategoryTagID,
		uiField.SortOrder,
		uiField.CreatedAt,
		uiField.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create UI field: %w", err)
	}

	return nil
}

// Update modifies an existing UI field
func (q *UIFieldQueries) Update(ctx context.Context, uiField *models.UIField) error {
	query := `
		UPDATE ui_fields
		SET ui_layout_id = $2, ui_group_id = $3, context_id = $4, category_tag_id = $5, sort_order = $6, updated_at = $7
		WHERE id = $1
	`

	result, err := q.db.ExecContext(ctx, query,
		uiField.ID,
		uiField.UILayoutID,
		uiField.UIGroupID,
		uiField.ContextID,
		uiField.CategoryTagID,
		uiField.SortOrder,
		uiField.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to update UI field: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("UI field not found")
	}

	return nil
}

// Delete removes a UI field by ID
func (q *UIFieldQueries) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM ui_fields WHERE id = $1`

	result, err := q.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete UI field: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("UI field not found")
	}

	return nil
}
