---
title: Tag Relationships
parent: Tags
nav_order: 2
---
# Tag Relationships

Defines how tags are semantically connected — hierarchically, associatively, or contextually — enabling graph traversal and reasoning across the tagging taxonomy.

## `tag_relationships` Table

Many-to-many join table linking [`tags`](tags.md#tags-table) in either **hierarchical** or **associative** ways. Relationships define how tags are connected or composed. The first tag, **Tag A**, represents the more "dominant" tag to **Tag B**; in other words, if the relationship between two tags is "Parent-Child", then **Tag A** is the Parent and **Tag B** is the child.

| Column | Type | Description |
|--------|------|--------------|
| `id` | UUID | Primary key. |
| `tag_a_id` | UUID | Foreign key to [**Tag A**](../index.md#tags-table), representing the source or dominant tag. |
| `tag_b_id` | UUID | Foreign key to [**Tag B**](../index.md#tags-table), representing the related or subordinate tag. |
| `relationship_type_id` | UUID | Foreign key to [`tag_relationship_types`](#tag_relationship_types-table). |
| `description` | TEXT | Optional human-readable description of the relationship (e.g., `"Vehicle is a parent of Car"`). |

### Directionality Semantics

- `tag_a_id` → represents the **dominant** or **source** tag (e.g., parent or general concept).  
- `tag_b_id` → represents the **subordinate** or **target** tag (e.g., child or specific instance).  

Example Relationships:

| Tag A | Relationship | Tag B |
|--------|---------------|-------|
| Vehicle | parent of | Car |
| Director | role of | Quentin Tarantino |
| Dog | species of | Animal |

## `tag_relationship_types` Table

Defines allowed relationship types for modeling tag hierarchies and associations.

| Column | Type | Description |
|--------|------|--------------|
| `id` | UUID | Primary key. |
| `name` | TEXT | Name of the relationship type (e.g., `"is a"`, `"is part of"`, `"is associated with"`). |

## API Endpoint Summary

| Category | Endpoint | CRUD Coverage | Description |
|-----------|-----------|----------------|--------------|
| **Tag Relationships** | `/tags/{id}/relationships` | **GET**, **POST**, **PATCH**, **DELETE** | Manage relationships originating from a specific tag. |
| | `/tags/{id}/relationships/{relationship_id}` | **GET**, **PATCH**, **DELETE** | Inspect, modify, or remove a single relationship. |
| **Relationship Ratings** | `/tags/{id}/relationship_ratings` | **GET**, **PATCH** | Rate the strength or relevance of a tag-to-tag relationship. |
| | `/tags/{id}/relationship_ratings/{rating_id}` | **GET**, **PATCH**, **DELETE** | Retrieve or modify a specific relationship rating. |
| **System-Wide Relationships** | `/tag-relationships`, `/tag-relationships/{id}` | **GET**, **POST**, **PATCH**, **DELETE** | Manage all tag relationships globally. |
| **System-Wide Relationship Ratings** | `/tag-relationship-ratings`, `/tag-relationship-ratings/{id}` | **GET**, **POST**, **PATCH**, **DELETE** | Manage all tag relationship ratings across the system. |

### CRUD Parity Overview

| Scope | CRUD | Description |
|--------|------|--------------|
| Per Tag | ✅ Full | Manage relationships and ratings linked to a single tag. |
| System-Wide | ✅ Full | Manage relationships and ratings globally. |

✅ = Full REST-compliant CRUD support.

### Constraints

- `tag_a_id` ≠ `tag_b_id`.  
- (`tag_a_id`, `tag_b_id`, `relationship_type_id`) must be unique.  
- Directionality must align with the meaning of the relationship type (e.g., “is parent of” must not be reversed).  
