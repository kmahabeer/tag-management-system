# Tags

The Tagging Service uses a unified tag table, enriched by types, aliases, and compositional relationships.

## `tags` Table

Stores the atomic and composite tags used throughout the system.

| Column      | Type  | Description                                                                              |
| ----------- | ----- | ---------------------------------------------------------------------------------------- |
| `id`        | UUID  | Primary key                                                                              |
| `name`      | TEXT  | Unique name for the tag                                                                  |
| `metadata` | JSONB | Arbitrary key-value data about the tag assignment (e.g., `{ "source": "Label Studio" }`) |

Composite tags (e.g., `"very big red car"`) are stored as standalone entries in this table, and linked to their component tags using the `tag_compositions` table described below.

## Core Tables

### `tag_aliases` Table

Alternative names or synonyms for a tag. Useful for search, filtering, and UI flexibility.

| Column   | Type | Description                     |
| -------- | ---- | ------------------------------- |
| `id`     | UUID | Primary key                     |
| `name`   | TEXT | Unique name for a tag's alias   |
| `tag_id` | UUID | Foreign key to the `tags` table |

### `tag_types` Table

Defines classifications for each tag such as "Person", "Vehicle", "Genre", etc.

| Column   | Type | Description                                  |
| -------- | ---- | -------------------------------------------- |
| `id`     | UUID | Primary key                                  |
| `name`   | TEXT | Name for a tag's type (e.g. "Color", "Role") |
| `tag_id` | UUID | Foreign key to the `tags` table              |

### `tag_relationship_types` Table

Defines the types of relationship between tags.

| Column | Type | Description                                                          |
| ------ | ---- | -------------------------------------------------------------------- |
| `id`   | UUID | Primary key                                                          |
| `name` | TEXT | Name for a tag relationship type (e.g. "Parent-Child", "Associated") |

## Join (aka Junction or linking) Tables

### `tag_tag_types` Table

Many-to-many join table linking `tags` to one or more `tag_types`.

| Column        | Type | Description                                          |
| ------------- | ---- | ---------------------------------------------------- |
| `id`          |      | Primary key; composite of `tag_id` and `tag_type_id` |
| `tag_id`      | UUID | Foreign key to `tags`                                |
| `tag_type_id` | UUID | Foreign key to `tag_types`                           |

### `tag_relationships` Table

Many-to-many join table linking `tags` in either **hierarchical** or **associative** ways. Relationships define how tags are connected or composed. The first tag, **Tag A**, represents the more "dominant" tag to **Tag B**; in other words, if the relationship between two tags is "Parent-Child", then **Tag A** is the Parent and **Tag B** is the child.

| Column                 | Type | Description                                                              |
| ---------------------- | ---- | ------------------------------------------------------------------------ |
| `id`                   |      | Primary key; composite of `tag_a_id`, `tag_b_id`, `relationship_type_id` |
| `tag_a_id`             | UUID | Foreign key to **Tag A**                                                 |
| `tag_b_id`             | UUID | Foreign key to **Tag B**                                                 |
| `relationship_type_id` | UUID | Foreign key to `tag_relationship_types` table                            |

Example: `"Director"` → `role-of` → `"Quentin Tarantino"`

**What it means semantically:**

> "**Tag A** is related to **Tag B** through a relationship of type `X`."

E.g.:

- **“Vehicle”** is the **parent** of **“Car”**
- **“Director”** is a **role** of **“Quentin Tarantino”**

### 🆕 `tag_compositions` Table

Used to define composite phrases (e.g., `"very big red car"`), where the composed tag is treated as a **single** tag in the UI and tagging logic, but internally linked to multiple atomic tags.

| Column             | Type | Description                                                        |
| ------------------ | ---- | ------------------------------------------------------------------ |
| `id`               |      | Primary key; Composite of `composite_tag_id`, `component_tag_id`   |
| `composite_tag_id` | UUID | Foreign key to a tag representing the composed phrase              |
| `component_tag_id` | UUID | Foreign key to a tag used as a part of the composite               |
| `position`         | INT  | Order of the component tag in the composite phrase (1-based index) |

For example, the phrase `"very big red car"` would be stored in `tags` as a single composite tag, and then linked to `"very"`, `"big"`, `"red"`, and `"car"` via this table in that order.

## `tag_types` Table vs `tag_relationships` table

### `tag_types` Table

This table classifies an individual tag into one or more **types** or **categories**.

**Example Use Cases:**

- "Car" → is of type → "Object"
- "Red" → is of type → "Color"
- "Quentin Tarantino" → is of type → "Person"
- "Director" → is of type → "Role"

**This means:**

> "**Tag X** is a member of the type **Y**."

### `tag_relationships` Table

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
| Can model composition?  | ❌ No                               | ✅ Yes (e.g., "Director" of "Film")             |
| Can model hierarchy?    | ❌ No                               | ✅ Yes (e.g., Parent-Child)                     |

### Use cases

| You want to…                                  | Use…                           |
| --------------------------------------------- | ------------------------------ |
| Group tags into categories like “Color”       | `tag_types`                    |
| Define how tags are **related to each other** | `tag_relationships`            |
| Represent roles, hierarchies, compositions    | `tag_relationships`            |
| Help the UI decide where tags go              | `tag_types` (with `ui_fields`) |
