# Tags

The Tagging Service uses a unified tag table, enriched by types, aliases, and compositional relationships.

## `tags` Table

Stores the atomic and composite tags used throughout the system.

| Column         | Type  | Description                                                                              |
| -------------- | ----- | ---------------------------------------------------------------------------------------- |
| `id`           | UUID  | Primary key                                                                              |
| `name`         | TEXT  | Unique name for the tag                                                                  |
| `display_name` | TEXT  |                                                                                          |
| `metadata`     | JSONB | Arbitrary key-value data about the tag assignment (e.g., `{ "source": "Label Studio" }`) |

`Metadata` in this context informs about the tag itself (its origin, category, creation information).

Composite tags (e.g., `"very big red car"`) are stored as standalone entries in this table, and linked to their component tags using the [`tag_compositions`](./tags.md#tag_compositions-table) table described below.

## Core Tables

### `tag_aliases` Table

Alternative names or synonyms for a tag. Useful for search, filtering, and UI flexibility.

| Column   | Type | Description                                           |
| -------- | ---- | ----------------------------------------------------- |
| `id`     | UUID | Primary key                                           |
| `name`   | TEXT | Unique name for a tag's alias                         |
| `tag_id` |      | Foreign key to the [`tags`](tags.md#tags-table) table |

### `tag_relationship_types` Table

Defines the types of relationship between tags.

| Column | Type | Description                                                                                                   |
| ------ | ---- | ------------------------------------------------------------------------------------------------------------- |
| `id`   | UUID | Primary key                                                                                                   |
| `name` | TEXT | Name for a tag relationship type (e.g. "is a", "is associated with", "is a parent to", "is a predecessor to") |

### `tag_context_ratings` Table

| Column    | Type | Description                                                                |
| --------- | ---- | -------------------------------------------------------------------------- |
| `id`      | UUID | Primary key                                                                |
| `tag_id`  | UUID | Foreign key to the [`tags`](tags.md#tags-table) table                      |
| `context` |      | Foreign key to the [`contexts`](./utilities.md#contexts) table             |
| `rating`  |      | Foreign key to the [`ratings`](./utilities.md#ratings) table               |
| `user_id` |      | Foreign key to the `users` table to allow per-user ratings per tag context |

## Join (aka Junction or linking) Tables

### `tag_relationships` Table

Many-to-many join table linking [`tags`](tags.md#tags-table) in either **hierarchical** or **associative** ways. Relationships define how tags are connected or composed. The first tag, **Tag A**, represents the more "dominant" tag to **Tag B**; in other words, if the relationship between two tags is "Parent-Child", then **Tag A** is the Parent and **Tag B** is the child.

| Column                 | Type | Description                                                                                                      |
| ---------------------- | ---- | ---------------------------------------------------------------------------------------------------------------- |
| `id`                   | UUID | Primary key                                                                                                      |
| `tag_a_id`             |      | Foreign key to [**Tag A**](./tags.md#tags-table)                                                                 |
| `tag_b_id`             |      | Foreign key to [**Tag B**](./tags.md#tags-table)                                                                 |
| `relationship_type_id` |      | Foreign key to [`tag_relationship_types`](./tags.md#tag_relationship_types-table) table                          |
| `description`          | TEXT | Description explaining the relationship (e.g., `"{Tag A} is a {Tag B}"`, `"{Tag A} is associated with {Tag B}"`) |

The description allows systems to render a human-readable explanation of the relationship. Example: `"Director"` → `role-of` → `"Quentin Tarantino"`

**What it means semantically:**

> "**Tag A** is related to **Tag B** through a relationship of type `X`."

E.g.:

- **“Vehicle”** is the **parent** of **“Car”**
- **“Director”** is a **role** of **“Quentin Tarantino”**

### `tag_relationship_ratings` Table

| Column      | Type | Description                                                    |
| ----------- | ---- | -------------------------------------------------------------- |
| `id`        | UUID | Primary key                                                    |
| `tag_a_id`  |      | Foreign key to [**Tag A**](./tags.md#tags-table)               |
| `tag_b_id`  |      | Foreign key to [**Tag B**](./tags.md#tags-table)               |
| `context`   |      | Foreign key to the [`contexts`](./utilities.md#contexts) table |
| `rating_id` |      | Foreign key to the [`ratings`](./utilities.md#ratings) table   |

### `tag_compositions` Table

Used to define composite phrases (e.g., `"very big red car"`), where the composite tag is treated as a **single** tag but internally linked to multiple atomic tags.

Each composite tag is stored in the `tags` table and linked to one or more component tags using this table. The `position` field determines the order of components (1-based index).

| Column             | Type | Description                                                                                                  |
| ------------------ | ---- | ------------------------------------------------------------------------------------------------------------ |
| `id`               | UUID | Primary key                                                                                                  |
| `base_tag_id`      |      | Foreign key to a [tag](./tags.md#tags-table), representing the **base tag** (e.g., `"very big red car"`)     |
| `component_tag_id` |      | Foreign key to a [tag](./tags.md#tags-table), representing an atomic **component** (e.g., `"very"`, `"big"`) |
| `position`         | INT  | Order of the component tag in the composite phrase (1-based index)                                           |

> **Base tags** are treated as atomic when composing higher-level phrases. For instance, once "**red car**" is stored as a tag in the [`tags`](./tags.md#tags-table) table, it may itself be used as a base to form "big **_red car_**", which in turn may be used to form "very **_big red car_**", and so on.
>
> In this way, while a composite tag may consist of many atomic components, each composition step always joins two tags: the left-hand "base" (which may itself be composite) and a new component.

> [!warning] TODO
>
> - Enforce uniqueness on base_tag_id + ordered list of component_tag_ids.
> - Prevent semantic composites from being atomic: e.g., `"very big"` must only exist as a composite, not a standalone atomic tag.
