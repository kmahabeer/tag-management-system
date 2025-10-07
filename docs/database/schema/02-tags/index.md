---
title: Tags
parent: Database Schema
has_children: true
nav_order: 2
---
# Tags

The Tag Management System uses a unified tag table, enriched by types, aliases, and compositional relationships.

## `tags` Table

Stores the atomic and composite tags used throughout the system.

| Column           		| Type  | Description                                                                              |
| --------------------- | ----- | ---------------------------------------------------------------------------------------- |
| `id`             		| UUID  | Primary key                                                                              |
| `name`           		| TEXT  | Unique name for the tag                                                                  |
| `display_name`   		| TEXT  | An optional human-friendly label for use in the UI, defaults to `name` if blank.         |
| `metadata`       		| JSONB | Arbitrary key-value data about the tag assignment (e.g., `{ "source": "Label Studio" }`) |
| `part_of_speech_id`	|       | Foreign key to [`parts_of_speech`](./utilities/parts_of_speech.md) table                 |

`Metadata` in this context informs about the tag itself (its origin, category, creation information).

Composite tags (e.g., `"very big red car"`) are stored as standalone entries in this table, and linked to their component tags using the [`tag_compositions`](./tags.md#tag_compositions-table) table described below. Composite phrases that are semantically equivalent (e.g., "very big car" and "huge car") can be linked through the [`tag_aliases`](./tags.md#tag_aliases-table) table, allowing flexible user input while maintaining canonical tags for storage and querying.

## Sub-Resource Summary

| Subsystem                | Collection Endpoint               | Individual Endpoint                           | CRUD Coverage                            | Description                                                                      |
| ------------------------ | --------------------------------- | --------------------------------------------- | ---------------------------------------- | -------------------------------------------------------------------------------- |
| **Aliases**              | `/tags/{id}/aliases`              | `/tags/{id}/aliases/{alias_id}`               | **GET**, **POST**, **PATCH**, **DELETE** | Manage alternate names or synonyms for a canonical tag.                          |
| **Relationships**        | `/tags/{id}/relationships`        | `/tags/{id}/relationships/{relationship_id}`  | **GET**, **POST**, **PATCH**, **DELETE** | Define directional semantic links between tags (e.g., “Vehicle → Car”).          |
| **Compositions**         | `/tags/{id}/compositions`         | `/tags/{id}/compositions/{composition_id}`    | **GET**, **POST**, **PATCH**, **DELETE** | Manage atomic components that form composite phrases (e.g., “very big red car”). |
| **Tag Ratings**          | `/tags/{id}/ratings`              | `/tags/{id}/ratings/{rating_id}`              | **GET**, **PATCH**, **DELETE**           | Apply contextual ratings to a tag (e.g., *Dog* rated 9/10 for “likeness”).       |
| **Relationship Ratings** | `/tags/{id}/relationship_ratings` | `/tags/{id}/relationship_ratings/{rating_id}` | **GET**, **PATCH**, **DELETE**           | Rate the connection between two tags (e.g., “Dog → Animal” relevance = 0.95).    |

### Supporting System-Wide Resources

These global collections provide reusable data that all tag sub-resources reference:

| Global Resource                     | Endpoints                                       | CRUD Coverage                            | Description                                                         |
| ----------------------------------- | ----------------------------------------------- | ---------------------------------------- | ------------------------------------------------------------------- |
| **Tag Relationships (system-wide)** | `/tag-relationships`, `/tag-relationships/{id}` | **GET**, **POST**, **PATCH**, **DELETE** | Manage relationships globally, independent of a specific tag scope. |
| **Tag Aliases (system-wide)**       | `/tag-aliases`, `/tag-aliases/{id}`             | **GET**, **POST**, **PATCH**, **DELETE** | Manage aliases across all tags.                                     |
| **Tag Compositions (system-wide)**  | `/tag-compositions`, `/tag-compositions/{id}`   | **GET**, **POST**, **PATCH**, **DELETE** | Manage all composition records globally.                            |
| **Ratings**                         | `/ratings`, `/ratings/{id}`                     | **GET**, **POST**, **PATCH**, **DELETE** | Define specific rating values (e.g., “Excellent”, 9/10).            |
| **Rating Types**                    | `/rating-types`, `/rating-types/{id}`           | **GET**, **POST**, **PATCH**, **DELETE** | Define categories of ratings (e.g., “likeness”, “clarity”).         |
| **Contexts**                        | `/contexts`, `/contexts/{id}`                   | **GET**, **POST**, **PATCH**, **DELETE** | Define the contextual dimensions for applying tags and ratings.     |

### Directionality and Relationship Rules

| Relationship                       | Description                                                          | Enforced Rule                                 |
| ---------------------------------- | -------------------------------------------------------------------- | --------------------------------------------- |
| `tag_a_id` → `tag_b_id`            | Defines the direction of dominance or hierarchy between tags.        | Tag A = *source*, Tag B = *target*.           |
| `base_tag_id` → `component_tag_id` | Defines composite structure for multi-word phrases.                  | Base = composite, Component = atomic element. |
| `context_id`                       | Links ratings and tags to their contextual category (e.g., “style”). | Must reference an active context.             |
| `rating_id`                        | References a specific score or label in `/ratings`.                  | Must belong to a valid `rating_type_id`.      |

### CRUD Parity Overview

| Subsystem            | Collection | Individual | System-Wide |
| -------------------- | ---------- | ---------- | ----------- |
| Aliases              | ✅          | ✅          | ✅           |
| Relationships        | ✅          | ✅          | ✅           |
| Compositions         | ✅          | ✅          | ✅           |
| Ratings              | ✅          | ✅          | ✅           |
| Relationship Ratings | ✅          | ✅          | ✅           |

✅ = full REST-compliant CRUD support (GET, POST, PATCH, DELETE).

### Notes

- Every sub-resource (e.g., alias, relationship, composition, rating) has a unique UUID ID field. This ensures full traceability and allows for addressable, auditable records.
- System-wide tables (`tag_aliases`, `tag_relationships`, `tag_compositions`) act as join tables that power both collection-level and individual-level endpoints.
- `contexts` and `rating_types` form the semantic layer for applying structured meaning to tags and ratings.
