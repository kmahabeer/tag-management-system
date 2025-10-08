---
title: Ratings
parent: Utilities
nav_order: 2
---
# Ratings

Ratings provide a unified way to represent **qualitative** or **quantitative** evaluations across tags, relationships, and entities. They allow the system to express clarity, confidence, quality, or other contextual judgments in a consistent and machine-readable format.

## `ratings` Table

| Column           | Type | Description                                                                        |
| ---------------- | ---- | ---------------------------------------------------------------------------------- |
| `id`             | UUID | Primary key                                                                        |
| `name`           | TEXT | Label used to identify the rating (e.g., '9/10', 'excellent').                     |
| `score`          | INT  | Numeric scale                                                                      |
| `description`    | TEXT |                                                                                    |
| `rating_type_id` |      | Foreign key to the [`rating_types`](ui_configurations.md#rating_types-table) table |

### Purpose

`rating_types` add **semantic meaning** to raw rating values, ensuring consistency across domains.

**Examples:**

- `"10/10"` can mean “I love it” under the *likeness* rating type.  
- `"9/10"` can mean “very clear” under the *clarity* rating type.  
- The same score may therefore mean different things in different contexts.

## `rating_types` Table

| Column          | Type    | Description                                         |
| --------------- | ------- | --------------------------------------------------- |
| `id`            | UUID    | Primary key                                         |
| `name`          | TEXT    | e.g., `"likeness"`, `"confidence"`, `"clarity"`     |
| `is_normalized` | BOOLEAN | Whether the rating should be on a 1–10 or 0–1 scale |

## `tag_relationship_ratings` Table

Stores **contextual ratings** of semantic relationships between tags, such as how strongly two tags are related.

| Column       | Type | Description                                                          |
| ------------ | ---- | -------------------------------------------------------------------- |
| `id`         | UUID | Primary key                                                          |
| `tag_a_id`   |      | Foreign key to [**Tag A**](../../tags/index.md#tags-table).                       |
| `tag_b_id`   |      | Foreign key to [**Tag B**](../../tags/index.md#tags-table).                       |
| `context_id` |      | Foreign key to the [`contexts`](../../entity_tagging/contexts.md#contexts-table) table. |
| `rating_id`  |      | Foreign key to the [`ratings`](#ratings-table) table.	   |

**Examples:**

- `"Dog"` → `"Animal"` relevance = 0.95  
- `"Vehicle"` → `"Car"` clarity = 8/10  

## `tag_context_ratings` Table

Stores **contextual ratings** directly on tags (e.g., evaluating a tag’s relevance or quality within a given context).

| Column | Type | Description |
|--------|------|--------------|
| `id` | UUID | Primary key. |
| `tag_id` | UUID | Foreign key to the [`tags`](../../tags/index.md#tags-table) table. |
| `context_id` | UUID | Foreign key to the [`contexts`](../../entity_tagging/contexts.md#contexts-table) table. |
| `rating_id` | UUID | Foreign key to the [`ratings`](#ratings-table) table. |
| `user_id` | UUID | Foreign key to the `users` table to support per-user contextual ratings. |

**Examples:**

- Tag `"Dog"` rated `9/10` in `"content"` context.  
- Tag `"Monochrome"` rated `7/10` in `"style"` context.

## `entity_relationship_ratings` Table

Stores **contextual ratings of relationships between entities**, similar to how `tag_relationship_ratings` measures semantic strength between tags.

| Column | Type | Description |
|--------|------|--------------|
| `id` | UUID | Primary key. |
| `entity_a_id` | UUID | Foreign key to the [`entities`](../../entities/index.md#entities-table) table (source entity). |
| `entity_b_id` | UUID | Foreign key to the [`entities`](../../entities/index.md#entities-table) table (target entity). |
| `context_id` | UUID | Foreign key to the [`contexts`](../../entity_tagging/contexts.md#contexts-table) table. |
| `rating_id` | UUID | Foreign key to the [`ratings`](#ratings-table) table. |

**Examples:**

- `"Processed Render"` → `"Final Render"` confidence = 0.92  
- `"RAW Image"` → `"Compressed Version"` clarity = 6/10  
- `"Dataset"` → `"Output Report"` relevance = 0.87  
- Rating confidence of a derived version (“Processed → Final Render” confidence = 0.92)
- Quality assessment between source and derivative (“RAW → Compressed” = low clarity)
- Workflow relevance (“Input Dataset → Output Report” rated for consistency)

## API Endpoint Summary

| Category | Endpoint | CRUD Coverage | Description |
|-----------|-----------|----------------|--------------|
| **Ratings** | `/ratings` | **GET**, **POST** | Retrieve or create rating values. |
| | `/ratings/{id}` | **GET**, **PATCH**, **DELETE** | Retrieve, update, or delete a specific rating. |
| **Rating Types** | `/rating-types` | **GET**, **POST** | Retrieve or create rating type definitions. |
| | `/rating-types/{id}` | **GET**, **PATCH**, **DELETE** | Retrieve, update, or delete a rating type. |

### CRUD Parity Overview

| Scope | CRUD | Description |
|--------|------|--------------|
| Ratings | ✅ Full | Manage individual rating values (scores and labels). |
| Rating Types | ✅ Full | Manage rating categories and normalization rules. |
| Relationship Ratings | ✅ Full | Manage contextual ratings between entities or tags. |

✅ = Full REST-compliant CRUD support.

## Summary

Ratings unify the representation of **evaluative data** across entities and tags. By separating *rating values* from *rating types*, the system can express subjective and objective assessments with clear semantics — making it possible to rank, compare, and query meaningfully across domains.
