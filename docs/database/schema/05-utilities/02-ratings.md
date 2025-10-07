---
title: Ratings
parent: Utilities
nav_order: 2
---
# Ratings

## `ratings` Table

| Column           | Type | Description                                                                        |
| ---------------- | ---- | ---------------------------------------------------------------------------------- |
| `id`             | UUID | Primary key                                                                        |
| `name`           | TEXT | Label used to identify the rating (e.g., '9/10', 'excellent').                     |
| `score`          | INT  | Numeric scale                                                                      |
| `description`    | TEXT |                                                                                    |
| `rating_type_id` |      | Foreign key to the [`rating_types`](ui_configurations.md#rating_types-table) table |

Rating types add semantic meaning to user selected ratings. This way:

- "10/10" can mean "I love it" if it's a *likeness* rating.
- "9/10" can mean "very clear" if it's a *clarity* rating.
- You can use the same rating table for different semantics based on context.

## `rating_types` Table

| Column          | Type    | Description                                         |
| --------------- | ------- | --------------------------------------------------- |
| `id`            | UUID    | Primary key                                         |
| `name`          | TEXT    | e.g., `"likeness"`, `"confidence"`, `"clarity"`     |
| `is_normalized` | BOOLEAN | Whether the rating should be on a 1–10 or 0–1 scale |

## `tag_relationship_ratings` Table

| Column       | Type | Description                                                          |
| ------------ | ---- | -------------------------------------------------------------------- |
| `id`         | UUID | Primary key                                                          |
| `tag_a_id`   |      | Foreign key to [**Tag A**](tags.md#tags-table)                       |
| `tag_b_id`   |      | Foreign key to [**Tag B**](tags.md#tags-table)                       |
| `context_id` |      | Foreign key to the [`contexts`](ui_configurations.md#contexts) table |
| `rating_id`  |      | Foreign key to the [`ratings`](ui_configurations.md#ratings) table   |

## `tag_context_ratings` Table

| Column       | Type | Description                                                                |
| ------------ | ---- | -------------------------------------------------------------------------- |
| `id`         | UUID | Primary key                                                                |
| `tag_id`     | UUID | Foreign key to the [`tags`](tags.md#tags-table) table                      |
| `context_id` |      | Foreign key to the [`contexts`](ui_configurations.md#contexts) table       |
| `rating`     |      | Foreign key to the [`ratings`](ui_configurations.md#ratings) table         |
| `user_id`    |      | Foreign key to the `users` table to allow per-user ratings per tag context |

## `entity_relationship_ratings` Table

Stores contextual ratings of relationships between entities, similar to how `tag_relationship_ratings` captures semantic strength between tags.

| Column       | Type | Description                                                                 	|
| ------------- | ---- | ------------------------------------------------------------------------------ |
| `id`          | UUID | Primary key                                                                 	|
| `entity_a_id` | UUID | Foreign key to the [`entities`](../01-entities/index.md) table; source entity 	|
| `entity_b_id` | UUID | Foreign key to the [`entities`](../01-entities/index.md) table; target entity 	|
| `context_id`  | UUID | Foreign key to the [`contexts`](./contexts.md#contexts-table) table          	|
| `rating_id`   | UUID | Foreign key to the [`ratings`](./ratings.md#ratings-table) table             	|

Examples:

- Rating confidence of a derived version (“Processed → Final Render” confidence = 0.92)
- Quality assessment between source and derivative (“RAW → Compressed” = low clarity)
- Workflow relevance (“Input Dataset → Output Report” rated for consistency)
