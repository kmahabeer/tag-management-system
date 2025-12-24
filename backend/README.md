# Tag Management System - Backend

This is the Go backend for the Tag Management System, providing RESTful APIs for managing tags, entities, relationships, and UI configurations.

## Building and Running

### Prerequisites

- Go 1.25.5 or later
- PostgreSQL database

### Setup

1. Ensure dependencies are installed: `make go-tidy`
2. Build the server: `make go-build`
3. Run the server: `make go-run`

### Development Workflow

- To clean build artifacts: `make go-clean`
- To rebuild from scratch: `make go-rebuild`
- To run after building: `make go-run`

The binary is output to `bin/server` and can be run directly: `./bin/server`

### Validation

API input structs include `Validate()` methods that enforce required fields and basic constraints as defined in the OpenAPI specification. These validations ensure data integrity before processing requests.

### Vector Database Compatibility

API schemas include optional `embedding` fields (`[]float64`) in `Tag`, `Entity`, `TagInput`, and `EntityInput` structs for compatibility with vector databases like pgVector. These fields allow storing and retrieving vector embeddings for similarity search and AI-powered features.

### Type Conversions

The `backend/internal/api/converters.go` file provides functions to convert between database models and API schemas:

- `TagToAPI(models.Tag) Tag`: Converts database Tag to API Tag
- `TagInputToDB(TagInput) models.Tag`: Converts API TagInput to database Tag
- `EntityToAPI(models.Entity) Entity`: Converts database Entity to API Entity
- `EntityInputToDB(EntityInput) models.Entity`: Converts API EntityInput to database Entity

These functions handle type conversions between nullable database fields (`sql.NullString`, `json.RawMessage`) and API fields (`*string`, `any`).

## Database Schema

The Tag Management System uses a PostgreSQL database schema designed to manage digital artifacts (entities), hierarchical and composite tags, contextual tagging, ratings, and UI configurations. The schema emphasizes semantic relationships, versioning, and flexible tagging to support workflows like content classification, annotation, and retrieval. Below is a detailed summary based on the SQL initialization script and documentation.

## Database Structure Overview

The schema consists of 18 core tables, grouped into categories: shared lookup tables, tag-related tables, entity-related tables, UI configurations, and utilities. It uses UUIDs for primary keys, JSONB for flexible metadata, and foreign key constraints for referential integrity. Triggers automatically update `updated_at` timestamps. Extensions include `uuid-ossp` for UUID generation.

### Shared Lookup Tables

These provide foundational data for classification and relationships.

- **`parts_of_speech`**: Defines grammatical roles (e.g., noun, adjective, adverb) for tags, enabling validation of composite tag structures. Supports linguistic rules in tag composition.
- **`contexts`**: Classifies tagging contexts as "subjective" (e.g., style, emotion) or "objective" (e.g., content, subject, metadata). Ensures tags are applied with semantic meaning (e.g., "Dog" as content vs. subject).
- **`tag_relationship_types`**: Defines relationship types between tags (e.g., "is a", "parent of").
- **`entity_relationship_types`**: Defines relationship types between entities (e.g., "version-of", "belongs-to-set").
- **`rating_types`**: Categories for ratings (e.g., "likeness", "confidence"), with a flag for normalization (0-1 or 1-10 scale).
- **`ratings`**: Specific rating values (e.g., "9/10", "excellent") linked to rating types, used for evaluating tags, entities, and relationships.

### Core Tag Tables

Tags are central, supporting aliases, relationships, compositions, and ratings.

- **`tags`**: Stores all tags (atomic and composite) with optional display names, metadata, and part-of-speech links. Unique names ensure no duplicates.
- **`tag_aliases`**: Alternative names for tags (e.g., "puppy" for "dog"), improving searchability. Each alias maps to one canonical tag.
- **`tag_relationships`**: Directional links between tags (e.g., "Vehicle" → "Car" as "parent of"). Enforces uniqueness on (tag_a_id, tag_b_id, relationship_type_id) and prevents self-references.
- **`tag_compositions`**: Breaks down composite tags (e.g., "very big red car") into ordered components (e.g., "very" at position 1, "car" at 4). Ensures uniqueness on (base_tag_id, component_tag_id, position). Supports recursive composition (e.g., "red car" as a base for "big red car").
- **`tag_context_ratings`**: Contextual ratings on tags (e.g., "Dog" rated 9/10 in "content" context), with optional user attribution.
- **`tag_relationship_ratings`**: Ratings on tag relationships (e.g., strength of "Dog" → "Animal" link).

### Entity Tables

Entities represent digital artifacts, with tagging, purposes, and relationships.

- **`entities`**: Core table for artifacts (e.g., images, documents) with names, locations, primary flags, and metadata. Primary entities are canonical versions.
- **`entity_purposes`**: Links entities to purpose tags (e.g., "Concept Art"), with one primary purpose per entity. Enforces uniqueness on primary purposes.
- **`entity_relationships`**: Directional links between entities (e.g., "Sketch v1" → "Final Render" as "version-of"). Prevents self-references and enforces uniqueness on (entity_a_id, entity_b_id, relationship_type_id).
- **`entity_relationship_ratings`**: Ratings on entity relationships (e.g., confidence in a version link).
- **`entity_tags`**: Many-to-many join for tagging entities with contexts and metadata. Allows the same tag-entity pair in multiple contexts (e.g., "Obama" as subject and content).

### UI Configuration Tables

These enable dynamic UI grouping based on entity purposes and contexts.

- **`ui_layouts`**: Defines layouts per purpose tag (e.g., for "Concept Art"), with fallbacks for null purposes.
- **`ui_groups`**: Reusable UI sections (e.g., "Subjects", "Environment").
- **`ui_fields`**: Maps contexts, category tags, and groups within layouts, with sort orders. Drives UI rendering by grouping entity tags hierarchically (context → group → tags).

### Key Relationships and Constraints

- **Foreign Keys**: Extensive use ensures referential integrity (e.g., tags reference parts_of_speech; entity_tags reference contexts).
- **Unique Indexes**: Prevent duplicates (e.g., unique tag names, unique compositions, unique relationships).
- **Check Constraints**: Enforce rules like no self-referencing entities/tags, primary purpose limits, and rating normalization.
- **Triggers**: Auto-update timestamps on changes.
- **Cascade Deletes**: Remove dependent records (e.g., deleting a tag cascades to aliases and compositions).

## Complex Relationships and Features

### Versioning

- Handled via `entity_relationships` with directional semantics: `entity_a_id` (primary/source) → `entity_b_id` (alternate/derived). Supports hierarchies like versions, sets, or alternates.
- No dedicated versioning table; derived from relationships. Ratings on relationships assess quality (e.g., clarity of a derived version).
- Constraints prevent cycles and ensure unidirectional flow.

### Ratings

- Unified across tags, entities, and relationships using `ratings` and `rating_types`. Supports subjective (e.g., likeness) and objective (e.g., confidence) evaluations.
- Contextual: Always tied to contexts (e.g., rating a tag in "style" context).
- Normalization: Enforces scales based on type (e.g., 0-1 for probabilities).
- Enables ranking, filtering, and querying (e.g., high-confidence tags).

### Composite Tagging

- Composite tags (e.g., "very big red car") are stored as full tags in `tags`, linked to atomic components via `tag_compositions` with positions.
- Recursive: Composites can be bases for further compositions (e.g., "red car" → "big red car").
- Validation: Application-layer logic enforces grammar (e.g., adverb → adjective → noun) using parts-of-speech. Prevents invalid structures (e.g., "very big" without a noun).
- Uniqueness: Enforced on ordered component lists to avoid duplicates.
- Querying: Views or joins reconstruct compositions for display/search.

## Support for Tag Management Functionality

- **Tagging Flexibility**: Contextual tagging via `entity_tags` allows nuanced assignments (e.g., same tag in multiple roles). Aliases and relationships enable taxonomy building (e.g., hierarchies for search).
- **Entity Management**: Purposes and relationships support workflows (e.g., grouping by purpose in UI). Versioning tracks evolution without redundancy.
- **UI Customization**: Configurations dynamically group tags by purpose and context, ensuring consistent displays (e.g., concept art layouts).
- **Semantic Integrity**: Constraints and validations maintain consistency (e.g., grammatical compositions, unique relationships).
- **Scalability**: JSONB metadata allows extensibility. UUIDs support distributed systems.
- **API Coverage**: Full CRUD on all tables via REST endpoints, with per-entity, per-tag, and global scopes.

This schema supports complex tagging needs like annotation pipelines, content libraries, and AI-assisted classification, with strong emphasis on semantics, relationships, and user-driven configurations. For implementation details, refer to the API routes and service layers.
