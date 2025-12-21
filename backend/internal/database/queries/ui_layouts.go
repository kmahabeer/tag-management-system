package queries

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/kmahabeer/tag-management-system/backend/internal/models"
)

// UILayoutQueries implements UILayoutRepository
type UILayoutQueries struct {
	db *sql.DB
}

// NewUILayoutQueries creates a new UILayoutQueries instance
func NewUILayoutQueries(db *sql.DB) *UILayoutQueries {
	return &UILayoutQueries{db: db}
}

// GetByID retrieves a UI layout by its ID
func (q *UILayoutQueries) GetByID(ctx context.Context, id uuid.UUID) (*models.UILayout, error) {
	query := `
		SELECT id, purpose_tag_id, name, created_at, updated_at
		FROM ui_layouts
		WHERE id = $1
	`

	var uiLayout models.UILayout
	var purposeTagID sql.NullString
	var name sql.NullString
	err := q.db.QueryRowContext(ctx, query, id).Scan(
		&uiLayout.ID,
		&purposeTagID,
		&name,
		&uiLayout.CreatedAt,
		&uiLayout.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("UI layout not found")
		}
		return nil, fmt.Errorf("failed to get UI layout: %w", err)
	}

	if purposeTagID.Valid {
		parsedID, err := uuid.Parse(purposeTagID.String)
		if err != nil {
			return nil, fmt.Errorf("invalid purpose_tag_id: %w", err)
		}
		uiLayout.PurposeTagID = &parsedID
	} else {
		uiLayout.PurposeTagID = nil
	}

	uiLayout.Name = name

	return &uiLayout, nil
}

// List retrieves a list of UI layouts with pagination
func (q *UILayoutQueries) List(ctx context.Context, limit, offset int) ([]*models.UILayout, error) {
	query := `
		SELECT id, purpose_tag_id, name, created_at, updated_at
		FROM ui_layouts
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := q.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list UI layouts: %w", err)
	}
	defer rows.Close()

	var uiLayouts []*models.UILayout
	for rows.Next() {
		var uiLayout models.UILayout
		var purposeTagID sql.NullString
		var name sql.NullString
		err := rows.Scan(
			&uiLayout.ID,
			&purposeTagID,
			&name,
			&uiLayout.CreatedAt,
			&uiLayout.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan UI layout: %w", err)
		}

		if purposeTagID.Valid {
			parsedID, err := uuid.Parse(purposeTagID.String)
			if err != nil {
				return nil, fmt.Errorf("invalid purpose_tag_id: %w", err)
			}
			uiLayout.PurposeTagID = &parsedID
		} else {
			uiLayout.PurposeTagID = nil
		}

		uiLayout.Name = name
		uiLayouts = append(uiLayouts, &uiLayout)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return uiLayouts, nil
}

// Create inserts a new UI layout
func (q *UILayoutQueries) Create(ctx context.Context, uiLayout *models.UILayout) error {
	query := `
		INSERT INTO ui_layouts (id, purpose_tag_id, name, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`

	var purposeTagID interface{}
	if uiLayout.PurposeTagID != nil {
		purposeTagID = uiLayout.PurposeTagID.String()
	} else {
		purposeTagID = nil
	}

	_, err := q.db.ExecContext(ctx, query,
		uiLayout.ID,
		purposeTagID,
		uiLayout.Name,
		uiLayout.CreatedAt,
		uiLayout.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create UI layout: %w", err)
	}

	return nil
}

// Update modifies an existing UI layout
func (q *UILayoutQueries) Update(ctx context.Context, uiLayout *models.UILayout) error {
	query := `
		UPDATE ui_layouts
		SET purpose_tag_id = $2, name = $3, updated_at = $4
		WHERE id = $1
	`

	var purposeTagID interface{}
	if uiLayout.PurposeTagID != nil {
		purposeTagID = uiLayout.PurposeTagID.String()
	} else {
		purposeTagID = nil
	}

	result, err := q.db.ExecContext(ctx, query,
		uiLayout.ID,
		purposeTagID,
		uiLayout.Name,
		uiLayout.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to update UI layout: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("UI layout not found")
	}

	return nil
}

// Delete removes a UI layout by ID
func (q *UILayoutQueries) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM ui_layouts WHERE id = $1`

	result, err := q.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete UI layout: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("UI layout not found")
	}

	return nil
}
