package database

import (
	"context"
	"database/sql"

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

// EntityPurposeRepository defines the interface for entity purpose database operations
type EntityPurposeRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.EntityPurpose, error)
	List(ctx context.Context, limit, offset int) ([]*models.EntityPurpose, error)
	Create(ctx context.Context, entityPurpose *models.EntityPurpose) error
	Update(ctx context.Context, entityPurpose *models.EntityPurpose) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// EntityRelationshipRepository defines the interface for entity relationship database operations
type EntityRelationshipRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.EntityRelationship, error)
	List(ctx context.Context, limit, offset int) ([]*models.EntityRelationship, error)
	Create(ctx context.Context, entityRelationship *models.EntityRelationship) error
	Update(ctx context.Context, entityRelationship *models.EntityRelationship) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// EntityRelationshipTypeRepository defines the interface for entity relationship type database operations
type EntityRelationshipTypeRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.EntityRelationshipType, error)
	List(ctx context.Context, limit, offset int) ([]*models.EntityRelationshipType, error)
	Create(ctx context.Context, entityRelationshipType *models.EntityRelationshipType) error
	Update(ctx context.Context, entityRelationshipType *models.EntityRelationshipType) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// TagAliasRepository defines the interface for tag alias database operations
type TagAliasRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.TagAlias, error)
	List(ctx context.Context, limit, offset int) ([]*models.TagAlias, error)
	Create(ctx context.Context, tagAlias *models.TagAlias) error
	Update(ctx context.Context, tagAlias *models.TagAlias) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// TagRelationshipRepository defines the interface for tag relationship database operations
type TagRelationshipRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.TagRelationship, error)
	List(ctx context.Context, limit, offset int) ([]*models.TagRelationship, error)
	Create(ctx context.Context, tagRelationship *models.TagRelationship) error
	Update(ctx context.Context, tagRelationship *models.TagRelationship) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// TagRelationshipTypeRepository defines the interface for tag relationship type database operations
type TagRelationshipTypeRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.TagRelationshipType, error)
	List(ctx context.Context, limit, offset int) ([]*models.TagRelationshipType, error)
	Create(ctx context.Context, tagRelationshipType *models.TagRelationshipType) error
	Update(ctx context.Context, tagRelationshipType *models.TagRelationshipType) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// TagCompositionRepository defines the interface for tag composition database operations
type TagCompositionRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.TagComposition, error)
	List(ctx context.Context, limit, offset int) ([]*models.TagComposition, error)
	Create(ctx context.Context, tagComposition *models.TagComposition) error
	Update(ctx context.Context, tagComposition *models.TagComposition) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// EntityTagRepository defines the interface for entity tag database operations
type EntityTagRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.EntityTag, error)
	List(ctx context.Context, limit, offset int) ([]*models.EntityTag, error)
	Create(ctx context.Context, entityTag *models.EntityTag) error
	Update(ctx context.Context, entityTag *models.EntityTag) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// ContextRepository defines the interface for context database operations
type ContextRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.Context, error)
	List(ctx context.Context, limit, offset int) ([]*models.Context, error)
	Create(ctx context.Context, context *models.Context) error
	Update(ctx context.Context, context *models.Context) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// UILayoutRepository defines the interface for UI layout database operations
type UILayoutRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.UILayout, error)
	List(ctx context.Context, limit, offset int) ([]*models.UILayout, error)
	Create(ctx context.Context, uiLayout *models.UILayout) error
	Update(ctx context.Context, uiLayout *models.UILayout) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// UIGroupRepository defines the interface for UI group database operations
type UIGroupRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.UIGroup, error)
	List(ctx context.Context, limit, offset int) ([]*models.UIGroup, error)
	Create(ctx context.Context, uiGroup *models.UIGroup) error
	Update(ctx context.Context, uiGroup *models.UIGroup) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// UIFieldRepository defines the interface for UI field database operations
type UIFieldRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.UIField, error)
	List(ctx context.Context, limit, offset int) ([]*models.UIField, error)
	Create(ctx context.Context, uiField *models.UIField) error
	Update(ctx context.Context, uiField *models.UIField) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// PartOfSpeechRepository defines the interface for part of speech database operations
type PartOfSpeechRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.PartOfSpeech, error)
	List(ctx context.Context, limit, offset int) ([]*models.PartOfSpeech, error)
	Create(ctx context.Context, partOfSpeech *models.PartOfSpeech) error
	Update(ctx context.Context, partOfSpeech *models.PartOfSpeech) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// RatingRepository defines the interface for rating database operations
type RatingRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.Rating, error)
	List(ctx context.Context, limit, offset int) ([]*models.Rating, error)
	Create(ctx context.Context, rating *models.Rating) error
	Update(ctx context.Context, rating *models.Rating) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// RatingTypeRepository defines the interface for rating type database operations
type RatingTypeRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.RatingType, error)
	List(ctx context.Context, limit, offset int) ([]*models.RatingType, error)
	Create(ctx context.Context, ratingType *models.RatingType) error
	Update(ctx context.Context, ratingType *models.RatingType) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// TagRelationshipRatingRepository defines the interface for tag relationship rating database operations
type TagRelationshipRatingRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.TagRelationshipRating, error)
	List(ctx context.Context, limit, offset int) ([]*models.TagRelationshipRating, error)
	Create(ctx context.Context, tagRelationshipRating *models.TagRelationshipRating) error
	Update(ctx context.Context, tagRelationshipRating *models.TagRelationshipRating) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// TagContextRatingRepository defines the interface for tag context rating database operations
type TagContextRatingRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.TagContextRating, error)
	List(ctx context.Context, limit, offset int) ([]*models.TagContextRating, error)
	Create(ctx context.Context, tagContextRating *models.TagContextRating) error
	Update(ctx context.Context, tagContextRating *models.TagContextRating) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// EntityRelationshipRatingRepository defines the interface for entity relationship rating database operations
type EntityRelationshipRatingRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.EntityRelationshipRating, error)
	List(ctx context.Context, limit, offset int) ([]*models.EntityRelationshipRating, error)
	Create(ctx context.Context, entityRelationshipRating *models.EntityRelationshipRating) error
	Update(ctx context.Context, entityRelationshipRating *models.EntityRelationshipRating) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// TransactionManager defines the interface for transaction management
type TransactionManager interface {
	BeginTx(ctx context.Context) (*sql.Tx, error)
	Commit(tx *sql.Tx) error
	Rollback(tx *sql.Tx) error
}
