package api

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

// ErrorResponse represents an error response
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

// PaginatedResponse represents a paginated response
type PaginatedResponse struct {
	Results []any `json:"results"`
	Total   int   `json:"total"`
}

// Tag represents a tag
type Tag struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	DisplayName    *string   `json:"display_name,omitempty"`
	Metadata       any       `json:"metadata,omitempty"`
	PartOfSpeechID uuid.UUID `json:"part_of_speech_id"`
	Embedding      []float64 `json:"embedding,omitempty"` // pgVector compatible
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// TagInput represents input for creating or updating a tag
type TagInput struct {
	Name           string    `json:"name"`
	DisplayName    *string   `json:"display_name,omitempty"`
	Metadata       any       `json:"metadata,omitempty"`
	PartOfSpeechID uuid.UUID `json:"part_of_speech_id"`
	Embedding      []float64 `json:"embedding,omitempty"` // pgVector compatible
}

// Validate checks if the TagInput is valid
func (t *TagInput) Validate() error {
	if strings.TrimSpace(t.Name) == "" {
		return errors.New("name is required and cannot be empty")
	}
	if t.PartOfSpeechID == uuid.Nil {
		return errors.New("part_of_speech_id is required")
	}
	return nil
}

// TagAlias represents an alias for a tag
type TagAlias struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	TagID uuid.UUID `json:"tag_id"`
}

// AliasUpdate represents bulk update for aliases
type AliasUpdate struct {
	Aliases []string `json:"aliases"`
}

// TagRelationship represents a relationship between two tags
type TagRelationship struct {
	ID                 uuid.UUID `json:"id"`
	TagAID             uuid.UUID `json:"tag_a_id"`
	TagBID             uuid.UUID `json:"tag_b_id"`
	RelationshipTypeID uuid.UUID `json:"relationship_type_id"`
	Description        *string   `json:"description,omitempty"`
}

// TagRelationshipInput represents input for creating a tag relationship
type TagRelationshipInput struct {
	TagBID             uuid.UUID `json:"tag_b_id"`
	RelationshipTypeID uuid.UUID `json:"relationship_type_id"`
	Description        *string   `json:"description,omitempty"`
}

// Validate checks if the TagRelationshipInput is valid
func (t *TagRelationshipInput) Validate() error {
	if t.TagBID == uuid.Nil {
		return errors.New("tag_b_id is required")
	}
	if t.RelationshipTypeID == uuid.Nil {
		return errors.New("relationship_type_id is required")
	}
	return nil
}

// RelationshipUpdate represents bulk update for relationships
type RelationshipUpdate struct {
	Relationships []TagRelationshipInput `json:"relationships"`
}

// TagComponent represents a component in a composite tag
type TagComponent struct {
	ID             uuid.UUID `json:"id"`
	BaseTagID      uuid.UUID `json:"base_tag_id"`
	ComponentTagID uuid.UUID `json:"component_tag_id"`
	Position       int       `json:"position"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// TagComponentInput represents input for adding a component
type TagComponentInput struct {
	ComponentTagID uuid.UUID `json:"component_tag_id"`
	Position       int       `json:"position"`
}

// Validate checks if the TagComponentInput is valid
func (t *TagComponentInput) Validate() error {
	if t.ComponentTagID == uuid.Nil {
		return errors.New("component_tag_id is required")
	}
	if t.Position < 1 {
		return errors.New("position must be at least 1")
	}
	return nil
}

// CompositionUpdate represents bulk update for compositions
type CompositionUpdate struct {
	Components []TagComponentInput `json:"components"`
}

// Entity represents an entity
type Entity struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Location  *string   `json:"location,omitempty"`
	IsPrimary bool      `json:"is_primary"`
	Metadata  any       `json:"metadata,omitempty"`
	Embedding []float64 `json:"embedding,omitempty"` // pgVector compatible
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// EntityInput represents input for creating or updating an entity
type EntityInput struct {
	Name      string    `json:"name"`
	Location  *string   `json:"location,omitempty"`
	IsPrimary bool      `json:"is_primary"`
	Metadata  any       `json:"metadata,omitempty"`
	Embedding []float64 `json:"embedding,omitempty"` // pgVector compatible
}

// Validate checks if the EntityInput is valid
func (e *EntityInput) Validate() error {
	if strings.TrimSpace(e.Name) == "" {
		return errors.New("name is required and cannot be empty")
	}
	return nil
}

// EntityTagAssignment represents a tag assignment to an entity
type EntityTagAssignment struct {
	Tag       Tag       `json:"tag"`
	ContextID uuid.UUID `json:"context_id"`
	Metadata  any       `json:"metadata,omitempty"`
}

// EntityTagAssignmentInput represents input for assigning a tag
type EntityTagAssignmentInput struct {
	TagID     uuid.UUID `json:"tag_id"`
	ContextID uuid.UUID `json:"context_id"`
	Metadata  any       `json:"metadata,omitempty"`
}

// EntityTagUpdate represents bulk update for entity tags
type EntityTagUpdate struct {
	Tags []EntityTagAssignmentInput `json:"tags"`
}

// EntityPurposeInput represents input for assigning a purpose
type EntityPurposeInput struct {
	PurposeTagID uuid.UUID `json:"purpose_tag_id"`
	IsPrimary    bool      `json:"is_primary"`
}

// EntityPurposeUpdate represents bulk update for purposes
type EntityPurposeUpdate struct {
	Purposes []EntityPurposeInput `json:"purposes"`
}

// EntityRelationship represents a relationship between entities
type EntityRelationship struct {
	EntityAID          uuid.UUID `json:"entity_a_id"`
	EntityBID          uuid.UUID `json:"entity_b_id"`
	RelationshipTypeID uuid.UUID `json:"relationship_type_id"`
}

// EntityRelationshipInput represents input for creating an entity relationship
type EntityRelationshipInput struct {
	EntityBID          uuid.UUID `json:"entity_b_id"`
	RelationshipTypeID uuid.UUID `json:"relationship_type_id"`
}

// EntityRelationshipUpdate represents bulk update for entity relationships
type EntityRelationshipUpdate struct {
	Relationships []EntityRelationshipInput `json:"relationships"`
}

// Context represents a semantic context
type Context struct {
	ID                 uuid.UUID `json:"id"`
	Name               string    `json:"name"`
	ClassificationType string    `json:"classification_type"`
	Description        *string   `json:"description,omitempty"`
	IsActive           bool      `json:"is_active"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

// ContextInput represents input for creating or updating a context
type ContextInput struct {
	Name               string  `json:"name"`
	ClassificationType string  `json:"classification_type"`
	Description        *string `json:"description,omitempty"`
	IsActive           bool    `json:"is_active"`
}

// Validate checks if the ContextInput is valid
func (c *ContextInput) Validate() error {
	if strings.TrimSpace(c.Name) == "" {
		return errors.New("name is required and cannot be empty")
	}
	if c.ClassificationType != "objective" && c.ClassificationType != "subjective" {
		return errors.New("classification_type must be either 'objective' or 'subjective'")
	}
	return nil
}

// PartOfSpeech represents a grammatical classification
type PartOfSpeech struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// PartOfSpeechInput represents input for creating or updating a part of speech
type PartOfSpeechInput struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	IsActive    bool    `json:"is_active"`
}

// Validate checks if the PartOfSpeechInput is valid
func (p *PartOfSpeechInput) Validate() error {
	if strings.TrimSpace(p.Name) == "" {
		return errors.New("name is required and cannot be empty")
	}
	return nil
}

// RatingType represents a category for ratings
type RatingType struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	IsNormalized bool      `json:"is_normalized"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// RatingTypeInput represents input for creating or updating a rating type
type RatingTypeInput struct {
	Name         string `json:"name"`
	IsNormalized bool   `json:"is_normalized"`
}

// Validate checks if the RatingTypeInput is valid
func (r *RatingTypeInput) Validate() error {
	if strings.TrimSpace(r.Name) == "" {
		return errors.New("name is required and cannot be empty")
	}
	return nil
}

// Rating represents a specific rating value
type Rating struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Score        int       `json:"score"`
	Description  *string   `json:"description,omitempty"`
	RatingTypeID uuid.UUID `json:"rating_type_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// RatingInput represents input for creating or updating a rating
type RatingInput struct {
	Name         string    `json:"name"`
	Score        int       `json:"score"`
	Description  *string   `json:"description,omitempty"`
	RatingTypeID uuid.UUID `json:"rating_type_id"`
}

// Validate checks if the RatingInput is valid
func (r *RatingInput) Validate() error {
	if strings.TrimSpace(r.Name) == "" {
		return errors.New("name is required and cannot be empty")
	}
	if r.RatingTypeID == uuid.Nil {
		return errors.New("rating_type_id is required")
	}
	return nil
}

// TagContextualRatingInput represents a rating applied to a tag
type TagContextualRatingInput struct {
	ID        uuid.UUID  `json:"id"`
	ContextID uuid.UUID  `json:"context_id"`
	RatingID  uuid.UUID  `json:"rating_id"`
	UserID    *uuid.UUID `json:"user_id,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// EntityContextualRatingInput represents a rating applied to an entity
type EntityContextualRatingInput struct {
	ID        uuid.UUID  `json:"id"`
	ContextID uuid.UUID  `json:"context_id"`
	RatingID  uuid.UUID  `json:"rating_id"`
	UserID    *uuid.UUID `json:"user_id,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// EntityContextualRatingUpdate represents update for entity ratings
type EntityContextualRatingUpdate struct {
	Rating  float64 `json:"rating"`
	Context string  `json:"context"`
}

// ContextualRatingUpdate represents bulk update for tag ratings
type ContextualRatingUpdate struct {
	ContextualRatings []TagContextualRatingInput `json:"contextual_ratings"`
}

// TagRelationshipRatingInput represents a rating on a tag relationship
type TagRelationshipRatingInput struct {
	ID        uuid.UUID `json:"id"`
	TagBID    uuid.UUID `json:"tag_b_id"`
	ContextID uuid.UUID `json:"context_id"`
	RatingID  uuid.UUID `json:"rating_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// EntityRelationshipRatingInput represents a rating on an entity relationship
type EntityRelationshipRatingInput struct {
	ID        uuid.UUID `json:"id"`
	EntityBID uuid.UUID `json:"entity_b_id"`
	ContextID uuid.UUID `json:"context_id"`
	RatingID  uuid.UUID `json:"rating_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// EntityRelationshipRatingUpdate represents bulk update for entity relationship ratings
type EntityRelationshipRatingUpdate struct {
	RelationshipRatings []EntityRelationshipRatingInput `json:"relationship_ratings"`
}

// RelationshipRatingUpdate represents bulk update for tag relationship ratings
type RelationshipRatingUpdate struct {
	RelationshipRatings []TagRelationshipRatingInput `json:"relationship_ratings"`
}

// UiLayout represents a UI layout profile
type UiLayout struct {
	ID           uuid.UUID  `json:"id"`
	PurposeTagID *uuid.UUID `json:"purpose_tag_id,omitempty"`
	Name         string     `json:"name"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

// UiLayoutInput represents input for creating or updating a UI layout
type UiLayoutInput struct {
	PurposeTagID *uuid.UUID `json:"purpose_tag_id,omitempty"`
	Name         string     `json:"name"`
}

// Validate checks if the UiLayoutInput is valid
func (u *UiLayoutInput) Validate() error {
	if strings.TrimSpace(u.Name) == "" {
		return errors.New("name is required and cannot be empty")
	}
	return nil
}

// UiGroup represents a UI group
type UiGroup struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UiGroupInput represents input for creating or updating a UI group
type UiGroupInput struct {
	Name string `json:"name"`
}

// Validate checks if the UiGroupInput is valid
func (u *UiGroupInput) Validate() error {
	if strings.TrimSpace(u.Name) == "" {
		return errors.New("name is required and cannot be empty")
	}
	return nil
}

// UiField represents a UI field
type UiField struct {
	ID            uuid.UUID  `json:"id"`
	UiLayoutID    uuid.UUID  `json:"ui_layout_id"`
	UiGroupID     uuid.UUID  `json:"ui_group_id"`
	ContextID     *uuid.UUID `json:"context_id,omitempty"`
	CategoryTagID *uuid.UUID `json:"category_tag_id,omitempty"`
	SortOrder     int        `json:"sort_order"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

// UiFieldInput represents input for creating or updating a UI field
type UiFieldInput struct {
	UiLayoutID    uuid.UUID  `json:"ui_layout_id"`
	UiGroupID     uuid.UUID  `json:"ui_group_id"`
	ContextID     *uuid.UUID `json:"context_id,omitempty"`
	CategoryTagID *uuid.UUID `json:"category_tag_id,omitempty"`
	SortOrder     int        `json:"sort_order"`
}

// Validate checks if the UiFieldInput is valid
func (u *UiFieldInput) Validate() error {
	if u.UiLayoutID == uuid.Nil {
		return errors.New("ui_layout_id is required")
	}
	if u.UiGroupID == uuid.Nil {
		return errors.New("ui_group_id is required")
	}
	return nil
}

// Response wrappers

// TagsGet200Response represents the response for GET /tags
type TagsGet200Response struct {
	Results []Tag `json:"results"`
	Total   int   `json:"total"`
}

// EntitiesGet200Response represents the response for GET /entities
type EntitiesGet200Response struct {
	Results []Entity `json:"results"`
	Total   int      `json:"total"`
}

// TagsIdAliasesGet200Response represents the response for GET /tags/{id}/aliases
type TagsIdAliasesGet200Response struct {
	Aliases []TagAlias `json:"aliases"`
}

// TagsIdRelationshipsGet200Response represents the response for GET /tags/{id}/relationships
type TagsIdRelationshipsGet200Response struct {
	Relationships []TagRelationship `json:"relationships"`
}

// TagsIdCompositionsGet200Response represents the response for GET /tags/{id}/compositions
type TagsIdCompositionsGet200Response struct {
	BaseTagID  uuid.UUID      `json:"base_tag_id"`
	Components []TagComponent `json:"components"`
}

// TagCompositionsGet200Response represents the response for GET /tag-compositions
type TagCompositionsGet200Response struct {
	Compositions []TagComponent `json:"compositions"`
}

// TagsIdRatingsGet200Response represents the response for GET /tags/{id}/ratings
type TagsIdRatingsGet200Response struct {
	TagID   uuid.UUID                  `json:"tag_id"`
	Ratings []TagContextualRatingInput `json:"ratings"`
}

// TagsIdRelationshipRatingsGet200Response represents the response for GET /tags/{id}/relationship_ratings
type TagsIdRelationshipRatingsGet200Response struct {
	TagAID              uuid.UUID                    `json:"tag_a_id"`
	RelationshipRatings []TagRelationshipRatingInput `json:"relationship_ratings"`
}

// TagsIdRelationshipRatingsPatch200Response represents the response for PATCH /tags/{id}/relationship_ratings
type TagsIdRelationshipRatingsPatch200Response struct {
	TagAID  uuid.UUID                    `json:"tag_a_id"`
	Ratings []TagRelationshipRatingInput `json:"ratings"`
}

// EntitiesIdTagsGet200Response represents the response for GET /entities/{id}/tags
type EntitiesIdTagsGet200Response struct {
	Tags []EntityTagAssignment `json:"tags"`
}

// EntitiesIdTagsPatch200Response represents the response for PATCH /entities/{id}/tags
type EntitiesIdTagsPatch200Response struct {
	EntityID uuid.UUID             `json:"entity_id"`
	Tags     []EntityTagAssignment `json:"tags"`
}

// EntitiesIdPurposesGet200Response represents the response for GET /entities/{id}/purposes
type EntitiesIdPurposesGet200Response struct {
	Purposes []EntityPurposeInput `json:"purposes"`
}

// EntitiesIdPurposesPatch200Response represents the response for PATCH /entities/{id}/purposes
type EntitiesIdPurposesPatch200Response struct {
	Purposes []EntityPurposeInput `json:"purposes"`
	EntityID uuid.UUID            `json:"entity_id"`
}

// EntitiesIdVersionsGet200Response represents the response for GET /entities/{id}/versions
type EntitiesIdVersionsGet200Response struct {
	Versions []Entity `json:"versions"`
}

// EntitiesIdVersionsPatch200Response represents the response for PATCH /entities/{id}/versions
type EntitiesIdVersionsPatch200Response struct {
	EntityAID uuid.UUID            `json:"entity_a_id"`
	Versions  []EntityRelationship `json:"versions"`
}

// EntitiesIdRelationshipsGet200Response represents the response for GET /entities/{id}/relationships
type EntitiesIdRelationshipsGet200Response struct {
	Relationships []EntityRelationship `json:"relationships"`
}

// EntitiesIdRelationshipsPatch200Response represents the response for PATCH /entities/{id}/relationships
type EntitiesIdRelationshipsPatch200Response struct {
	EntityAID     uuid.UUID            `json:"entity_a_id"`
	Relationships []EntityRelationship `json:"relationships"`
}

// EntitiesIdRatingsGet200Response represents the response for GET /entities/{id}/ratings
type EntitiesIdRatingsGet200Response struct {
	Ratings []EntityContextualRatingInput `json:"ratings"`
}

// EntitiesIdRatingsPatch200Response represents the response for PATCH /entities/{id}/ratings
type EntitiesIdRatingsPatch200Response struct {
	EntityID uuid.UUID                     `json:"entity_id"`
	Ratings  []EntityContextualRatingInput `json:"ratings"`
}

// EntitiesIdRelationshipRatingsGet200Response represents the response for GET /entities/{id}/relationship_ratings
type EntitiesIdRelationshipRatingsGet200Response struct {
	RelationshipRatings []EntityRelationshipRatingInput `json:"relationship_ratings"`
}

// EntitiesIdRelationshipRatingsPatch200Response represents the response for PATCH /entities/{id}/relationship_ratings
type EntitiesIdRelationshipRatingsPatch200Response struct {
	EntityAID           uuid.UUID                       `json:"entity_a_id"`
	RelationshipRatings []EntityRelationshipRatingInput `json:"relationship_ratings"`
}

// EntityVersionsGet200Response represents the response for GET /entity-versions
type EntityVersionsGet200Response struct {
	Versions []EntityRelationship `json:"versions"`
}

// EntityRelationshipsGet200Response represents the response for GET /entity-relationships
type EntityRelationshipsGet200Response struct {
	Relationships []EntityRelationship `json:"relationships"`
}

// EntityRelationshipRatingsGet200Response represents the response for GET /entity-relationship-ratings
type EntityRelationshipRatingsGet200Response struct {
	RelationshipRatings []EntityRelationshipRatingInput `json:"relationship_ratings"`
}

// EntityRatingsGet200Response represents the response for GET /entity-ratings
type EntityRatingsGet200Response struct {
	Ratings []EntityContextualRatingInput `json:"ratings"`
}

// EntityPurposesGet200Response represents the response for GET /entity-purposes
type EntityPurposesGet200Response struct {
	Purposes []EntityPurposeInput `json:"purposes"`
}

// ContextsGet200Response represents the response for GET /contexts
type ContextsGet200Response struct {
	Contexts []Context `json:"contexts"`
}

// PartsOfSpeechGet200Response represents the response for GET /parts-of-speech
type PartsOfSpeechGet200Response struct {
	PartsOfSpeech []PartOfSpeech `json:"parts_of_speech"`
}

// RatingsGet200Response represents the response for GET /ratings
type RatingsGet200Response struct {
	Ratings []Rating `json:"ratings"`
}

// RatingTypesGet200Response represents the response for GET /rating-types
type RatingTypesGet200Response struct {
	RatingTypes []RatingType `json:"rating_types"`
}

// UiLayoutsGet200Response represents the response for GET /ui/layouts
type UiLayoutsGet200Response struct {
	UiLayouts []UiLayout `json:"ui_layouts"`
}

// UiGroupsGet200Response represents the response for GET /ui/groups
type UiGroupsGet200Response struct {
	UiGroups []UiGroup `json:"ui_groups"`
}

// UiFieldsGet200Response represents the response for GET /ui/fields
type UiFieldsGet200Response struct {
	UiFields []UiField `json:"ui_fields"`
}

// TagsIdDelete200Response represents the response for DELETE operations
type TagsIdDelete200Response struct {
	Status string `json:"status"`
}

// Additional input structs for patch requests

// TagsIdAliasesPostRequest represents input for POST /tags/{id}/aliases
type TagsIdAliasesPostRequest struct {
	Name string `json:"name"`
}

// Validate checks if the TagsIdAliasesPostRequest is valid
func (t *TagsIdAliasesPostRequest) Validate() error {
	if strings.TrimSpace(t.Name) == "" {
		return errors.New("name is required and cannot be empty")
	}
	return nil
}

// TagAliasesPostRequest represents input for POST /tag-aliases
type TagAliasesPostRequest struct {
	Name  string    `json:"name"`
	TagID uuid.UUID `json:"tag_id"`
}

// Validate checks if the TagAliasesPostRequest is valid
func (t *TagAliasesPostRequest) Validate() error {
	if strings.TrimSpace(t.Name) == "" {
		return errors.New("name is required and cannot be empty")
	}
	if t.TagID == uuid.Nil {
		return errors.New("tag_id is required")
	}
	return nil
}

// TagsIdAliasesAliasIdPatchRequest represents input for PATCH /tags/{id}/aliases/{alias_id}
type TagsIdAliasesAliasIdPatchRequest struct {
	Name  string    `json:"name"`
	TagID uuid.UUID `json:"tag_id"`
}

// Validate checks if the TagsIdAliasesAliasIdPatchRequest is valid
func (t *TagsIdAliasesAliasIdPatchRequest) Validate() error {
	// Both optional, but if provided, validate
	if t.Name != "" && strings.TrimSpace(t.Name) == "" {
		return errors.New("name cannot be empty if provided")
	}
	return nil
}

// TagRelationshipsIdPatchRequest represents input for PATCH /tag-relationships/{id}
type TagRelationshipsIdPatchRequest struct {
	RelationshipTypeID uuid.UUID `json:"relationship_type_id"`
	Description        *string   `json:"description,omitempty"`
}

// Validate checks if the TagRelationshipsIdPatchRequest is valid
func (t *TagRelationshipsIdPatchRequest) Validate() error {
	if t.RelationshipTypeID == uuid.Nil {
		return errors.New("relationship_type_id is required")
	}
	return nil
}

// TagsIdRelationshipsRelationshipIdPatchRequest represents input for PATCH /tags/{id}/relationships/{relationship_id}
type TagsIdRelationshipsRelationshipIdPatchRequest struct {
	RelationshipTypeID uuid.UUID `json:"relationship_type_id"`
	Description        *string   `json:"description,omitempty"`
}

// Validate checks if the TagsIdRelationshipsRelationshipIdPatchRequest is valid
func (t *TagsIdRelationshipsRelationshipIdPatchRequest) Validate() error {
	if t.RelationshipTypeID == uuid.Nil {
		return errors.New("relationship_type_id is required")
	}
	return nil
}

// TagCompositionsPostRequest represents input for POST /tag-compositions
type TagCompositionsPostRequest struct {
	BaseTagID      uuid.UUID `json:"base_tag_id"`
	ComponentTagID uuid.UUID `json:"component_tag_id"`
	Position       int       `json:"position"`
}

// Validate checks if the TagCompositionsPostRequest is valid
func (t *TagCompositionsPostRequest) Validate() error {
	if t.BaseTagID == uuid.Nil {
		return errors.New("base_tag_id is required")
	}
	if t.ComponentTagID == uuid.Nil {
		return errors.New("component_tag_id is required")
	}
	if t.Position < 1 {
		return errors.New("position must be at least 1")
	}
	return nil
}

// TagsIdCompositionsCompositionIdPatchRequest represents input for PATCH /tags/{id}/compositions/{composition_id}
type TagsIdCompositionsCompositionIdPatchRequest struct {
	Position int `json:"position"`
}

// Validate checks if the TagsIdCompositionsCompositionIdPatchRequest is valid
func (t *TagsIdCompositionsCompositionIdPatchRequest) Validate() error {
	if t.Position < 1 {
		return errors.New("position must be at least 1")
	}
	return nil
}

// TagCompositionsIdPatchRequest represents input for PATCH /tag-compositions/{id}
type TagCompositionsIdPatchRequest struct {
	Position int `json:"position"`
}

// Validate checks if the TagCompositionsIdPatchRequest is valid
func (t *TagCompositionsIdPatchRequest) Validate() error {
	if t.Position < 1 {
		return errors.New("position must be at least 1")
	}
	return nil
}

// EntitiesIdPurposesPurposeIdPatchRequest represents input for PATCH /entities/{id}/purposes/{purpose_id}
type EntitiesIdPurposesPurposeIdPatchRequest struct {
	PurposeTagID uuid.UUID `json:"purpose_tag_id"`
	IsPrimary    bool      `json:"is_primary"`
}

// Validate checks if the EntitiesIdPurposesPurposeIdPatchRequest is valid
func (e *EntitiesIdPurposesPurposeIdPatchRequest) Validate() error {
	if e.PurposeTagID == uuid.Nil {
		return errors.New("purpose_tag_id is required")
	}
	return nil
}

// EntitiesIdRelationshipsRelationshipIdPatchRequest represents input for PATCH /entities/{id}/relationships/{relationship_id}
type EntitiesIdRelationshipsRelationshipIdPatchRequest struct {
	RelationshipTypeID uuid.UUID `json:"relationship_type_id"`
}

// Validate checks if the EntitiesIdRelationshipsRelationshipIdPatchRequest is valid
func (e *EntitiesIdRelationshipsRelationshipIdPatchRequest) Validate() error {
	if e.RelationshipTypeID == uuid.Nil {
		return errors.New("relationship_type_id is required")
	}
	return nil
}

// EntitiesIdVersionsVersionIdPatchRequest represents input for PATCH /entities/{id}/versions/{version_id}
type EntitiesIdVersionsVersionIdPatchRequest struct {
	RelationshipTypeID uuid.UUID `json:"relationship_type_id"`
}

// Validate checks if the EntitiesIdVersionsVersionIdPatchRequest is valid
func (e *EntitiesIdVersionsVersionIdPatchRequest) Validate() error {
	if e.RelationshipTypeID == uuid.Nil {
		return errors.New("relationship_type_id is required")
	}
	return nil
}

// EntitiesIdRatingsRatingIdPatchRequest represents input for PATCH /entities/{id}/ratings/{rating_id}
type EntitiesIdRatingsRatingIdPatchRequest struct {
	ContextID uuid.UUID `json:"context_id"`
	RatingID  uuid.UUID `json:"rating_id"`
}

// Validate checks if the EntitiesIdRatingsRatingIdPatchRequest is valid
func (e *EntitiesIdRatingsRatingIdPatchRequest) Validate() error {
	if e.ContextID == uuid.Nil {
		return errors.New("context_id is required")
	}
	if e.RatingID == uuid.Nil {
		return errors.New("rating_id is required")
	}
	return nil
}

// TagsIdRelationshipRatingsRatingIdPatchRequest represents input for PATCH /tags/{id}/relationship_ratings/{rating_id}
type TagsIdRelationshipRatingsRatingIdPatchRequest struct {
	ContextID uuid.UUID `json:"context_id"`
	RatingID  uuid.UUID `json:"rating_id"`
}

// Validate checks if the TagsIdRelationshipRatingsRatingIdPatchRequest is valid
func (t *TagsIdRelationshipRatingsRatingIdPatchRequest) Validate() error {
	if t.ContextID == uuid.Nil {
		return errors.New("context_id is required")
	}
	if t.RatingID == uuid.Nil {
		return errors.New("rating_id is required")
	}
	return nil
}

// LabelStudio-compatible structs for data import

// LabelStudioData represents the data field in LabelStudio export
type LabelStudioData struct {
	Image string `json:"image"`
}

// LabelStudioValue represents the value field in annotations
type LabelStudioValue struct {
	Labels []string `json:"labels"`
}

// LabelStudioResult represents a single result in annotations
type LabelStudioResult struct {
	Value    LabelStudioValue `json:"value"`
	FromName string           `json:"from_name"`
	ToName   string           `json:"to_name"`
	Type     string           `json:"type"`
}

// LabelStudioAnnotation represents an annotation
type LabelStudioAnnotation struct {
	Result []LabelStudioResult `json:"result"`
}

// LabelStudioExport represents the full LabelStudio export JSON
type LabelStudioExport struct {
	Data        LabelStudioData         `json:"data"`
	Annotations []LabelStudioAnnotation `json:"annotations"`
}
