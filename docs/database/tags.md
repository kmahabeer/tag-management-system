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

Composite tags (e.g., `"very big red car"`) are stored as standalone entries in this table, and linked to their component tags using the [`tag_compositions`](./tags.md#tag_compositions-table) table described below.

## Core Tables

### `tag_aliases` Table

Alternative names or synonyms for a tag. Useful for search, filtering, and UI flexibility.

| Column   | Type | Description                                           |
| -------- | ---- | ----------------------------------------------------- |
| `id`     | UUID | Primary key                                           |
| `name`   | TEXT | Unique name for a tag's alias                         |
| `tag_id` |      | Foreign key to the [`tags`](tags.md#tags-table) table |

### `tag_types` Table

Defines classifications for each tag such as "Person", "Vehicle", "Genre", etc.

| Column   | Type | Description                                           |
| -------- | ---- | ----------------------------------------------------- |
| `id`     | UUID | Primary key                                           |
| `name`   | TEXT | Name for a tag's type (e.g. "Color", "Role")          |
| `tag_id` |      | Foreign key to the [`tags`](tags.md#tags-table) table |

### `tag_relationship_types` Table

Defines the types of relationship between tags.

| Column | Type | Description                                                          |
| ------ | ---- | -------------------------------------------------------------------- |
| `id`   | UUID | Primary key                                                          |
| `name` | TEXT | Name for a tag relationship type (e.g. "Parent-Child", "Associated") |

### `tag_context_ratings` Table

| Column    | Type | Description                                                                |
| --------- | ---- | -------------------------------------------------------------------------- |
| `id`      | UUID | Primary key                                                                |
| `tag_id`  | UUID | Foreign key to the [`tags`](tags.md#tags-table) table                      |
| `context` |      | Foreign key to the [`contexts`](./utilities.md#contexts) table             |
| `rating`  |      | Foreign key to the [`ratings`](./utilities.md#ratings) table               |
| `user_id` |      | Foreign key to the `users` table to allow per-user ratings per tag context |

## Join (aka Junction or linking) Tables

### `tag_tag_types` Table

Many-to-many join table linking [`tags`](tags.md#tags-table) to one or more [`tag_types`](tags.md#tag_types-table).

| Column        | Type | Description                                                     |
| ------------- | ---- | --------------------------------------------------------------- |
| `id`          | UUID | Primary key                                                     |
| `tag_id`      | UUID | Foreign key to the [`tags`](tags.md#tags-table) table           |
| `tag_type_id` | UUID | Foreign key to the [`tag_types`](tags.md#tag_types-table) table |

### `tag_relationships` Table

Many-to-many join table linking [`tags`](tags.md#tags-table) in either **hierarchical** or **associative** ways. Relationships define how tags are connected or composed. The first tag, **Tag A**, represents the more "dominant" tag to **Tag B**; in other words, if the relationship between two tags is "Parent-Child", then **Tag A** is the Parent and **Tag B** is the child.

| Column                 | Type | Description                                                                             |
| ---------------------- | ---- | --------------------------------------------------------------------------------------- |
| `id`                   | UUID | Primary key                                                                             |
| `tag_a_id`             |      | Foreign key to [**Tag A**](./tags.md#tags-table)                                        |
| `tag_b_id`             |      | Foreign key to [**Tag B**](./tags.md#tags-table)                                        |
| `relationship_type_id` |      | Foreign key to [`tag_relationship_types`](./tags.md#tag_relationship_types-table) table |

Example: `"Director"` → `role-of` → `"Quentin Tarantino"`

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

| Column             | Type | Description                                                                                                   |
| ------------------ | ---- | ------------------------------------------------------------------------------------------------------------- |
| `id`               | UUID | Primary key                                                                                                   |
| `composite_tag_id` |      | Foreign key to a [tag](./tags.md#tags-table), representing the **composite tag** (e.g., `"very big red car"`) |
| `component_tag_id` |      | Foreign key to a [tag](./tags.md#tags-table), representing an atomic **component** (e.g., `"very"`, `"big"`)  |
| `position`         | INT  | Order of the component tag in the composite phrase (1-based index)                                            |

For example, the phrase `"very big red car"` would be stored in [`tags`](./tags.md#tags-table) as a single composite tag, and then linked to `"very"`, `"big"`, `"red"`, and `"car"` via this table in that order.

Composite tags are first-class tags stored in the tags table, and therefore inherit support for:

- Tagging entities via `entity_tags`
- Contextual ratings via `tag_context_ratings`
- Aliases, metadata, and UI configuration

> [!warning] TODO
>
> - Enforce uniqueness on composite_tag_id + ordered list of component_tag_ids.
> - Prevent semantic composites from being atomic: e.g., `"very big"` must only exist as a composite, not a standalone atomic tag.

## `tag_types` Table vs `tag_relationships` Table

### Comparison: `tag_types` Table

This table classifies an individual tag into one or more **types** or **categories**.

**Example Use Cases:**

- "Car" → is of type → "Object"
- "Red" → is of type → "Color"
- "Quentin Tarantino" → is of type → "Person"
- "Director" → is of type → "Role"

**This means:**

> "**Tag X** is a member of the type **Y**."

### Comparison: `tag_relationships` Table

This table defines **relationships between two tags**, which can be either **hierarchical** or **associative**.

**Example Use Cases:**

- "Car" → is a → "Vehicle"
- "Director" → is a role of → "Quentin Tarantino"
- "New York" → is part of → "United States"

**What it means semantically:**

> "**Tag A** is related to **Tag B** through a relationship of type `X`."

**E.g.:**

- **“Vehicle”** is the **parent** of **“Car”**
- **“Director”** is a **role** of **“Quentin Tarantino”**

### Comparison

| Feature                 | `tag_types`                        | `tag_relationships`                            |
| ----------------------- | ---------------------------------- | ---------------------------------------------- |
| Type                    | Classification                     | Relationship / linkage                         |
| Scope                   | One tag → one or more categories   | Two tags → joined by a relationship type       |
| Data model              | Tag ↔ Tag Type (via join table)    | Tag A ↔ Tag B (via a third `relationship` tag) |
| Example                 | "Red" is a `Color`                 | "Red" modifies → "Car"                         |
| Used for                | UI grouping, filtering, validation | Inference, logic chains, hierarchy navigation  |
| Number of tags involved | 1                                  | 2 (Tag A and Tag B)                            |
| Relationship semantics  | "Belongs to category"              | "Is connected to another tag via…"             |
| Can model composition?  | ❌ No                              | ✅ Yes (e.g., "Director" of "Film")            |
| Can model hierarchy?    | ❌ No                              | ✅ Yes (e.g., Parent-Child)                    |

### Use cases

| You want to…                                  | Use…                           |
| --------------------------------------------- | ------------------------------ |
| Group tags into categories like “Color”       | `tag_types`                    |
| Define how tags are **related to each other** | `tag_relationships`            |
| Represent roles, hierarchies, compositions    | `tag_relationships`            |
| Help the UI decide where tags go              | `tag_types` (with `ui_fields`) |
