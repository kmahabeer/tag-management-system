---
title: Entities Purposes
---
# Entity Purposes

Defines the roles or intended use cases of each entity by linking it to one or more "purpose" tags.

## `entity_purposes` Table

Many-to-many join table which defines one or more **purposes** for an entity. A "purpose" is an entry within the [`tags`](./tags.md) table. Only one "purpose" should be marked as *primary* for each entity.

| Column           | Type    | Description                                                                        |
| ---------------- | ------- | ---------------------------------------------------------------------------------- |
| `id`             | UUID    | Primary key                                                                        |
| `entity_id`      |         | Foreign key to the [`entities`](./entities.md) table                               |
| `purpose_tag_id` |         | Foreign key to the [`tags`](./tags.md) table (must be a tag of the type "purpose") |
| `is_primary`     | BOOLEAN | Indicates if this is the primary purpose                                           |

Constraint: `purpose_tag_id` must refer to a tag of type `"Purpose"`.

> [!warning] TODO
>
> - Enforce via a SQL CHECK using a function or trigger if needed.
> Example: An image file might be both a "Figure Study" and a "Film Still", but only one is primary.
