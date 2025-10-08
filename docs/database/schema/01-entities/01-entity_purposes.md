---
title: Entity Purposes
parent: Entities
nav_order: 1
---
# Entity Purposes

Defines the intended roles or functions of each entity by linking it to one or more *purpose tags*.

A **purpose** represents *why* an entity exists — for example, `"Concept Art"`, `"Reference Photo"`, or `"Storyboard Frame"`. Each entity can have multiple purposes, but only one may be marked as **primary**.

## `entity_purposes` Table

Many-to-many join table linking entities to purpose-defining tags. A "purpose" is an entry within the [`tags`](./tags.md) table. Only one "purpose" should be marked as *primary* for each entity.

| Column | Type | Description |
|--------|------|--------------|
| `id` | UUID | Primary key. |
| `entity_id` | UUID | Foreign key to the [`entities`](../index.md#entities-table) table. |
| `purpose_tag_id` | UUID | Foreign key to the [`tags`](../../tags/index.md#tags-table) table. Must reference a tag classified as a purpose. |
| `is_primary` | BOOLEAN | Indicates whether this purpose is the entity’s primary role. |

### Constraints

- Each entity can have **multiple** purpose rows, but only **one** where `is_primary = TRUE`.  
- `purpose_tag_id` must reference a tag explicitly typed as `"Purpose"`.  
- Enforceable via SQL `CHECK` or trigger logic.

Example:

| entity_id | purpose_tag_id | is_primary |
|------------|----------------|-------------|
| a123... | tag_concept_art | true |
| a123... | tag_reference_photo | false |

## API Endpoint Summary

| Category | Endpoint | CRUD Coverage | Description |
|-----------|-----------|----------------|--------------|
| **Entity Purposes** | `/entities/{id}/purposes` | **GET**, **PATCH** | Retrieve or update all purposes associated with an entity. |
| | `/entities/{id}/purposes/{purpose_id}` | **GET**, **PATCH**, **DELETE** | Retrieve, modify, or remove a specific purpose assignment. |
| **System-Wide Purposes** | `/entity-purposes`, `/entity-purposes/{id}` | **GET**, **POST**, **PATCH**, **DELETE** | Manage purpose assignments across all entities. |

### CRUD Parity

| Scope | CRUD | Description |
|--------|------|--------------|
| Per Entity | ✅ Full | Assign or modify an entity’s purposes. |
| Global | ✅ Full | Manage all purpose records across entities. |

✅ = Full REST-compliant CRUD support.
