package queries

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/kmahabeer/tag-management-system/backend/internal/models"
)

// Fixtures provides database fixtures for testing
type Fixtures struct {
	entityQueries                   *EntityQueries
	tagQueries                      *TagQueries
	entityPurposeQueries            *EntityPurposeQueries
	entityRelationshipQueries       *EntityRelationshipQueries
	entityRelationshipTypeQueries   *EntityRelationshipTypeQueries
	tagAliasQueries                 *TagAliasQueries
	tagRelationshipQueries          *TagRelationshipQueries
	tagRelationshipTypeQueries      *TagRelationshipTypeQueries
	tagCompositionQueries           *TagCompositionQueries
	entityTagQueries                *EntityTagQueries
	contextQueries                  *ContextQueries
	uiLayoutQueries                 *UILayoutQueries
	uiGroupQueries                  *UIGroupQueries
	uiFieldQueries                  *UIFieldQueries
	partOfSpeechQueries             *PartOfSpeechQueries
	ratingQueries                   *RatingQueries
	ratingTypeQueries               *RatingTypeQueries
	tagRelationshipRatingQueries    *TagRelationshipRatingQueries
	tagContextRatingQueries         *TagContextRatingQueries
	entityRelationshipRatingQueries *EntityRelationshipRatingQueries
}

// NewFixtures creates a new Fixtures instance
func NewFixtures(db *sql.DB) *Fixtures {
	return &Fixtures{
		entityQueries:                   NewEntityQueries(db),
		tagQueries:                      NewTagQueries(db),
		entityPurposeQueries:            NewEntityPurposeQueries(db),
		entityRelationshipQueries:       NewEntityRelationshipQueries(db),
		entityRelationshipTypeQueries:   NewEntityRelationshipTypeQueries(db),
		tagAliasQueries:                 NewTagAliasQueries(db),
		tagRelationshipQueries:          NewTagRelationshipQueries(db),
		tagRelationshipTypeQueries:      NewTagRelationshipTypeQueries(db),
		tagCompositionQueries:           NewTagCompositionQueries(db),
		entityTagQueries:                NewEntityTagQueries(db),
		contextQueries:                  NewContextQueries(db),
		uiLayoutQueries:                 NewUILayoutQueries(db),
		uiGroupQueries:                  NewUIGroupQueries(db),
		uiFieldQueries:                  NewUIFieldQueries(db),
		partOfSpeechQueries:             NewPartOfSpeechQueries(db),
		ratingQueries:                   NewRatingQueries(db),
		ratingTypeQueries:               NewRatingTypeQueries(db),
		tagRelationshipRatingQueries:    NewTagRelationshipRatingQueries(db),
		tagContextRatingQueries:         NewTagContextRatingQueries(db),
		entityRelationshipRatingQueries: NewEntityRelationshipRatingQueries(db),
	}
}

// CreateTestEntity creates a test entity
func (f *Fixtures) CreateTestEntity(ctx context.Context, name, location string, isPrimary bool) (*models.Entity, error) {
	entity := &models.Entity{
		ID:        uuid.New(),
		Name:      name,
		Location:  sql.NullString{String: location, Valid: true},
		IsPrimary: isPrimary,
		Metadata:  []byte(`{"test": true}`),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := f.entityQueries.Create(ctx, entity)
	return entity, err
}

// CreateTestTag creates a test tag
func (f *Fixtures) CreateTestTag(ctx context.Context, name, displayName string, partOfSpeechID uuid.UUID) (*models.Tag, error) {
	tag := &models.Tag{
		ID:             uuid.New(),
		Name:           name,
		DisplayName:    sql.NullString{String: displayName, Valid: true},
		Metadata:       []byte(`{"test": true}`),
		PartOfSpeechID: partOfSpeechID,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	err := f.tagQueries.Create(ctx, tag)
	return tag, err
}

// CreateTestEntityPurpose creates a test entity purpose
func (f *Fixtures) CreateTestEntityPurpose(ctx context.Context, entityID, purposeTagID uuid.UUID, isPrimary bool) (*models.EntityPurpose, error) {
	entityPurpose := &models.EntityPurpose{
		ID:           uuid.New(),
		EntityID:     entityID,
		PurposeTagID: purposeTagID,
		IsPrimary:    isPrimary,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	err := f.entityPurposeQueries.Create(ctx, entityPurpose)
	return entityPurpose, err
}

// CreateTestEntityRelationship creates a test entity relationship
func (f *Fixtures) CreateTestEntityRelationship(ctx context.Context, entityAID, entityBID, relationshipTypeID uuid.UUID) (*models.EntityRelationship, error) {
	entityRelationship := &models.EntityRelationship{
		ID:                 uuid.New(),
		EntityAID:          entityAID,
		EntityBID:          entityBID,
		RelationshipTypeID: relationshipTypeID,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}
	err := f.entityRelationshipQueries.Create(ctx, entityRelationship)
	return entityRelationship, err
}

// CreateTestEntityRelationshipType creates a test entity relationship type
func (f *Fixtures) CreateTestEntityRelationshipType(ctx context.Context, name string) (*models.EntityRelationshipType, error) {
	entityRelationshipType := &models.EntityRelationshipType{
		ID:        uuid.New(),
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := f.entityRelationshipTypeQueries.Create(ctx, entityRelationshipType)
	return entityRelationshipType, err
}

// CreateTestTagAlias creates a test tag alias
func (f *Fixtures) CreateTestTagAlias(ctx context.Context, name string, tagID uuid.UUID) (*models.TagAlias, error) {
	tagAlias := &models.TagAlias{
		ID:        uuid.New(),
		Name:      name,
		TagID:     tagID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := f.tagAliasQueries.Create(ctx, tagAlias)
	return tagAlias, err
}

// CreateTestTagRelationship creates a test tag relationship
func (f *Fixtures) CreateTestTagRelationship(ctx context.Context, tagAID, tagBID, relationshipTypeID uuid.UUID, description string) (*models.TagRelationship, error) {
	tagRelationship := &models.TagRelationship{
		ID:                 uuid.New(),
		TagAID:             tagAID,
		TagBID:             tagBID,
		RelationshipTypeID: relationshipTypeID,
		Description:        sql.NullString{String: description, Valid: true},
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}
	err := f.tagRelationshipQueries.Create(ctx, tagRelationship)
	return tagRelationship, err
}

// CreateTestTagRelationshipType creates a test tag relationship type
func (f *Fixtures) CreateTestTagRelationshipType(ctx context.Context, name string) (*models.TagRelationshipType, error) {
	tagRelationshipType := &models.TagRelationshipType{
		ID:        uuid.New(),
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := f.tagRelationshipTypeQueries.Create(ctx, tagRelationshipType)
	return tagRelationshipType, err
}

// CreateTestTagComposition creates a test tag composition
func (f *Fixtures) CreateTestTagComposition(ctx context.Context, baseTagID, componentTagID uuid.UUID, position int) (*models.TagComposition, error) {
	tagComposition := &models.TagComposition{
		ID:             uuid.New(),
		BaseTagID:      baseTagID,
		ComponentTagID: componentTagID,
		Position:       position,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	err := f.tagCompositionQueries.Create(ctx, tagComposition)
	return tagComposition, err
}

// CreateTestEntityTag creates a test entity tag
func (f *Fixtures) CreateTestEntityTag(ctx context.Context, entityID, tagID, contextID uuid.UUID) (*models.EntityTag, error) {
	entityTag := &models.EntityTag{
		ID:        uuid.New(),
		EntityID:  entityID,
		TagID:     tagID,
		ContextID: contextID,
		Metadata:  []byte(`{"test": true}`),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := f.entityTagQueries.Create(ctx, entityTag)
	return entityTag, err
}

// CreateTestContext creates a test context
func (f *Fixtures) CreateTestContext(ctx context.Context, name, classificationType string, description string, isActive bool) (*models.Context, error) {
	context := &models.Context{
		ID:                 uuid.New(),
		Name:               name,
		ClassificationType: classificationType,
		Description:        sql.NullString{String: description, Valid: true},
		IsActive:           isActive,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}
	err := f.contextQueries.Create(ctx, context)
	return context, err
}

// CreateTestUILayout creates a test UI layout
func (f *Fixtures) CreateTestUILayout(ctx context.Context, purposeTagID *uuid.UUID, name string) (*models.UILayout, error) {
	uiLayout := &models.UILayout{
		ID:           uuid.New(),
		PurposeTagID: purposeTagID,
		Name:         sql.NullString{String: name, Valid: true},
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	err := f.uiLayoutQueries.Create(ctx, uiLayout)
	return uiLayout, err
}

// CreateTestUIGroup creates a test UI group
func (f *Fixtures) CreateTestUIGroup(ctx context.Context, name string) (*models.UIGroup, error) {
	uiGroup := &models.UIGroup{
		ID:        uuid.New(),
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := f.uiGroupQueries.Create(ctx, uiGroup)
	return uiGroup, err
}

// CreateTestUIField creates a test UI field
func (f *Fixtures) CreateTestUIField(ctx context.Context, uiLayoutID, uiGroupID, contextID, categoryTagID uuid.UUID, sortOrder int) (*models.UIField, error) {
	uiField := &models.UIField{
		ID:            uuid.New(),
		UILayoutID:    uiLayoutID,
		UIGroupID:     uiGroupID,
		ContextID:     contextID,
		CategoryTagID: categoryTagID,
		SortOrder:     sortOrder,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	err := f.uiFieldQueries.Create(ctx, uiField)
	return uiField, err
}

// CreateTestPartOfSpeech creates a test part of speech
func (f *Fixtures) CreateTestPartOfSpeech(ctx context.Context, name, description string, isActive bool) (*models.PartOfSpeech, error) {
	partOfSpeech := &models.PartOfSpeech{
		ID:          uuid.New(),
		Name:        name,
		Description: sql.NullString{String: description, Valid: true},
		IsActive:    isActive,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err := f.partOfSpeechQueries.Create(ctx, partOfSpeech)
	return partOfSpeech, err
}

// CreateTestRating creates a test rating
func (f *Fixtures) CreateTestRating(ctx context.Context, name string, score int, description string, ratingTypeID uuid.UUID) (*models.Rating, error) {
	rating := &models.Rating{
		ID:           uuid.New(),
		Name:         name,
		Score:        score,
		Description:  &description,
		RatingTypeID: ratingTypeID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	err := f.ratingQueries.Create(ctx, rating)
	return rating, err
}

// CreateTestRatingType creates a test rating type
func (f *Fixtures) CreateTestRatingType(ctx context.Context, name string, isNormalized bool) (*models.RatingType, error) {
	ratingType := &models.RatingType{
		ID:           uuid.New(),
		Name:         name,
		IsNormalized: isNormalized,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	err := f.ratingTypeQueries.Create(ctx, ratingType)
	return ratingType, err
}

// CreateTestTagRelationshipRating creates a test tag relationship rating
func (f *Fixtures) CreateTestTagRelationshipRating(ctx context.Context, tagAID, tagBID, contextID, ratingID uuid.UUID) (*models.TagRelationshipRating, error) {
	tagRelationshipRating := &models.TagRelationshipRating{
		ID:        uuid.New(),
		TagAID:    tagAID,
		TagBID:    tagBID,
		ContextID: contextID,
		RatingID:  ratingID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := f.tagRelationshipRatingQueries.Create(ctx, tagRelationshipRating)
	return tagRelationshipRating, err
}

// CreateTestTagContextRating creates a test tag context rating
func (f *Fixtures) CreateTestTagContextRating(ctx context.Context, tagID, contextID, ratingID uuid.UUID, userID *uuid.UUID) (*models.TagContextRating, error) {
	tagContextRating := &models.TagContextRating{
		ID:        uuid.New(),
		TagID:     tagID,
		ContextID: contextID,
		RatingID:  ratingID,
		UserID:    userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := f.tagContextRatingQueries.Create(ctx, tagContextRating)
	return tagContextRating, err
}

// CreateTestEntityRelationshipRating creates a test entity relationship rating
func (f *Fixtures) CreateTestEntityRelationshipRating(ctx context.Context, entityAID, entityBID, contextID, ratingID uuid.UUID) (*models.EntityRelationshipRating, error) {
	entityRelationshipRating := &models.EntityRelationshipRating{
		ID:        uuid.New(),
		EntityAID: entityAID,
		EntityBID: entityBID,
		ContextID: contextID,
		RatingID:  ratingID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := f.entityRelationshipRatingQueries.Create(ctx, entityRelationshipRating)
	return entityRelationshipRating, err
}
