---
title: Tags
parent: Database Schema
has_children: true
nav_order: 2
---
# Tags

The Tag Management System uses a unified `tags` table enriched by aliases, relationships, compositions, and contextual ratings.

## `tags` Table

| Column | Type | Description |
|--------|------|--------------|
| `id` | UUID | Primary key. |
| `name` | TEXT | Unique internal name for the tag. |
| `display_name` | TEXT | Optional human-readable label for UI display (defaults to `name`). |
| `metadata` | JSONB | Arbitrary key-value metadata about the tag. |
| `part_of_speech_id` | UUID | Foreign key to [`parts_of_speech`](../utilities/parts_of_speech.md). |

Composite tags (e.g., “very big red car”) are stored as standalone rows in this table, and linked to atomic tags via `tag_compositions`.  

## Tag Sub-Resources

| Subsystem | Collection Endpoint | Individual Endpoint | CRUD Coverage | Description |
|------------|--------------------|---------------------|----------------|--------------|
| **Aliases** | `/tags/{id}/aliases` | `/tags/{id}/aliases/{alias_id}` | **GET**, **POST**, **PATCH**, **DELETE** | Manage synonyms or alternate names for a canonical tag. |
| **Relationships** | `/tags/{id}/relationships` | `/tags/{id}/relationships/{relationship_id}` | **GET**, **POST**, **PATCH**, **DELETE** | Manage hierarchical or associative links between tags (e.g., “Vehicle → Car”). |
| **Compositions** | `/tags/{id}/compositions` | `/tags/{id}/compositions/{composition_id}` | **GET**, **POST**, **PATCH**, **DELETE** | Manage atomic components forming a composite phrase. |
| **Tag Ratings** | `/tags/{id}/ratings` | `/tags/{id}/ratings/{rating_id}` | **GET**, **PATCH**, **DELETE** | Apply contextual ratings to a tag (e.g., *Dog* rated 9/10 for “likeness”). |
| **Relationship Ratings** | `/tags/{id}/relationship_ratings` | `/tags/{id}/relationship_ratings/{rating_id}` | **GET**, **PATCH**, **DELETE** | Rate the strength or clarity of a tag relationship. |

## Global Tag Resources

| Resource | Endpoints | CRUD Coverage | Description |
|-----------|------------|----------------|--------------|
| **Tag Aliases** | `/tag-aliases`, `/tag-aliases/{id}` | **GET**, **POST**, **PATCH**, **DELETE** | Manage aliases across all tags. |
| **Tag Relationships** | `/tag-relationships`, `/tag-relationships/{id}` | **GET**, **POST**, **PATCH**, **DELETE** | Manage relationships across all tags. |
| **Tag Compositions** | `/tag-compositions`, `/tag-compositions/{id}` | **GET**, **POST**, **PATCH**, **DELETE** | Manage composite relationships globally. |
| **Contexts** | `/contexts`, `/contexts/{id}` | **GET**, **POST**, **PATCH**, **DELETE** | Define semantic dimensions for tagging and ratings. |
| **Ratings & Rating Types** | `/ratings`, `/rating-types` (and `/{id}` variants) | **GET**, **POST**, **PATCH**, **DELETE** | Define rating scales and categories. |

### Directionality Rules

| Relationship | Description | Enforced Rule |
|---------------|-------------|---------------|
| `tag_a_id → tag_b_id` | Directional dominance or hierarchy. | Tag A = source, Tag B = target. |
| `base_tag_id → component_tag_id` | Composite structure between tags. | Base = composite, Component = atomic. |
| `context_id` | Contextual category under which a tag or rating applies. | Must reference an active context. |

### CRUD Parity Overview

| Subsystem | Collection | Individual | System-Wide |
|------------|------------|-------------|-------------|
| Aliases | ✅ | ✅ | ✅ |
| Relationships | ✅ | ✅ | ✅ |
| Compositions | ✅ | ✅ | ✅ |
| Ratings | ✅ | ✅ | ✅ |
| Relationship Ratings | ✅ | ✅ | ✅ |

✅ = Full REST-compliant CRUD coverage.
