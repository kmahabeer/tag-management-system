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

// EntityRepository defines the interface for entity database operations
type EntityRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.Entity, error)
	List(ctx context.Context, limit, offset int) ([]*models.Entity, error)
	Create(ctx context.Context, entity *models.Entity) error
	Update(ctx context.Context, entity *models.Entity) error
	Delete(ctx context.Context, id uuid.UUID) error
}
