package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/kmahabeer/tag-management-system/backend/internal/models"
)

// TagRepository defines the interface for tag database operations
type TagRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.Tag, error)
	List(ctx context.Context, limit, offset int) ([]*models.Tag, error)
	Create(ctx context.Context, tag *models.Tag) error
	Update(ctx context.Context, tag *models.Tag) error
	Delete(ctx context.Context, id uuid.UUID) error
}
