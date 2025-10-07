---
title: Entity Relationships
parent: Entities
nav_order: 2
---
# Entity Relationships

Defines how entities are connected — hierarchically, associatively, or as alternate versions.  
Relationships support **versioning**, **grouping**, and **derivation** semantics between digital artifacts.

## `entity_relationships` Table

Join table linking two entities through a defined relationship type.

| Column | Type | Description |
|--------|------|--------------|
| `id` | UUID | Primary key. |
| `entity_a_id` | UUID | Foreign key to [**Entity A**](../index.md#entities-table), representing the source or parent entity. |
| `entity_b_id` | UUID | Foreign key to [**Entity B**](../index.md#entities-table), representing the target or derived entity. |
| `relationship_type_id` | UUID | Foreign key to [`entity_relationship_types`](#entity_relationship_types-table). |

### Directionality Semantics

The relationship model is **unidirectional**:

- `entity_a_id` → the *source*, canonical, or primary entity.  
- `entity_b_id` → the *derived*, alternate, or related entity.  

This structure enables consistent lineage queries and version reconstruction.

Example Relationships:

| Entity A | Relationship | Entity B |
|-----------|---------------|----------|
| Film Still A | alternate version of | Film Still B |
| Sketch X | belongs to set | Study Group A |
| Render v1 | version of | Render v2 |

## `entity_relationship_types` Table

Defines allowed relationship types between entities.

| Column | Type | Description |
|--------|------|--------------|
| `id` | UUID | Primary key. |
| `name` | TEXT | Type name (e.g., `"version-of"`, `"belongs-to-set"`, `"alternate-of"`). |

## API Endpoint Summary

| Category | Endpoint | CRUD Coverage | Description |
|-----------|-----------|----------------|--------------|
| **Entity Relationships** | `/entities/{id}/relationships` | **GET**, **POST**, **PATCH** | Manage relationships originating from a specific entity. |
| | `/entities/{id}/relationships/{relationship_id}` | **GET**, **PATCH**, **DELETE** | Retrieve, modify, or remove a specific relationship. |
| **Entity Relationship Ratings** | `/entities/{id}/relationship_ratings` | **GET**, **PATCH** | Rate the strength or confidence of relationships from this entity. |
| | `/entities/{id}/relationship_ratings/{rating_id}` | **GET**, **PATCH**, **DELETE** | Retrieve or modify a specific relationship rating. |
| **System-Wide Relationships** | `/entity-relationships`, `/entity-relationships/{id}` | **GET**, **POST**, **PATCH**, **DELETE** | Manage entity-to-entity relationships across the entire system. |
| **System-Wide Relationship Ratings** | `/entity-relationship-ratings`, `/entity-relationship-ratings/{id}` | **GET**, **POST**, **PATCH**, **DELETE** | Manage ratings applied to entity relationships globally. |

### CRUD Parity

| Scope | CRUD | Description |
|--------|------|--------------|
| Per Entity | ✅ Full | Manage relationships and ratings linked to a single entity. |
| System-Wide | ✅ Full | Manage relationships and ratings across all entities. |

✅ = Full REST-compliant CRUD support.

### Notes

- `entity_a_id` ≠ `entity_b_id` — self-links are invalid.  
- (`entity_a_id`, `entity_b_id`, `relationship_type_id`) must be unique.  
- Entity lineage (versioning) is derived from this relationship model; no dedicated versioning table exists.
