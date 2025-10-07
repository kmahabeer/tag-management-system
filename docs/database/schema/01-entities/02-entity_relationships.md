---
title: Entities Relationships
parent: Entities
nav_order: 2
---
# Entity Relationships

Enables [versioning](./utilities/versioning.md) and structural grouping between related entities such as alternates, studies, or variants.

## `entity_relationships` Table

Many-to-one join table linking many versions of an entity to the primary version of the entity. Defines versioning, grouping, or derivational relationships between entities.

| Column                 | Type | Description                                                                                 |
| ---------------------- | ---- | ------------------------------------------------------------------------------------------- |
| `id`                   | UUID | Primary key                                                                                 |
| `entity_a_id`          |      | Foreign key to [**Entity A**](./entities.md#entities)                                       |
| `entity_b_id`          |      | Foreign key to [**Entity B**](./entities.md#entities)                                       |
| `relationship_type_id` |      | Foreign key to [`entity_relationship_types`](./entities.md#entity_relationship_types) table |

> **Directionality Semantics**
>
> In the `entity_relationships` table, `entity_a_id` always refers to the source or parent entity, while `entity_b_id` refers to the derived, related, or child entity.
>
> For example:
>
> - "Sketch A" → belongs to group → "Figure Study Set"
> - "Film Still B" → alternate version of → "Film Still A"
>
> This convention enables consistent traversal of relationships and grouping logic.

### `entity_relationship_types` Table

Contains predefined relationship types for modeling associations between entities (e.g., versioning, grouping).

| Column | Type | Description                                                          |
| ------ | ---- | -------------------------------------------------------------------- |
| `id`   | UUID | Primary key                                                          |
| `name` | TEXT | Name for a tag relationship type (e.g. "Parent-Child", "Associated") |
