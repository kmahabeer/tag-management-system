# Tags

The Tagging Service uses a unified tag table, enriched by types, aliases, and compositional relationships.

## `tags` Table

Stores the atomic and composite tags used throughout the system.

| Column           | Type  | Description                                                                              |
| ---------------- | ----- | ---------------------------------------------------------------------------------------- |
| `id`             | UUID  | Primary key                                                                              |
| `name`           | TEXT  | Unique name for the tag                                                                  |
| `display_name`   | TEXT  | An optional human-friendly label for use in the UI, defaults to `name` if blank.         |
| `metadata`       | JSONB | Arbitrary key-value data about the tag assignment (e.g., `{ "source": "Label Studio" }`) |
| `part_of_speech` |       | Foreign key to [`parts_of_speech`](./utilities/parts_of_speech.md) table                 |

`Metadata` in this context informs about the tag itself (its origin, category, creation information).

Composite tags (e.g., `"very big red car"`) are stored as standalone entries in this table, and linked to their component tags using the [`tag_compositions`](./tags.md#tag_compositions-table) table described below. Composite phrases that are semantically equivalent (e.g., "very big car" and "huge car") can be linked through the [`tag_aliases`](./tags.md#tag_aliases-table) table, allowing flexible user input while maintaining canonical tags for storage and querying.

## Tag Aliases

Represents alternative names or synonyms that point to a canonical tag, aiding search and tagging flexibility.

### `tag_aliases` Table

Alternative names or synonyms for a tag. Useful for search, filtering, and UI flexibility.

| Column   | Type | Description                                           |
| -------- | ---- | ----------------------------------------------------- |
| `id`     | UUID | Primary key                                           |
| `name`   | TEXT | Unique name for a tag's alias                         |
| `tag_id` |      | Foreign key to the [`tags`](tags.md#tags-table) table |

## Tag Relationships

Defines how tags are connected hierarchically or associatively, enabling semantic linkages across the tag graph.

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

### What it means semantically

> "**Tag A** is related to **Tag B** through a relationship of type `X`."

E.g.:

- **“Vehicle”** is the **parent** of **“Car”**
- **“Director”** is a **role** of **“Quentin Tarantino”**

> [!TODO] TODO
> For `tag_relationships`, it says that Tag A is the dominant (e.g., parent) to Tag B, but the system doesn’t enforce that based on the `relationship_type`. Add constraints or semantic checks to enforce the relationships.

### `tag_relationship_types` Table

Defines the types of relationship between tags.

| Column | Type | Description                                                                                                   |
| ------ | ---- | ------------------------------------------------------------------------------------------------------------- |
| `id`   | UUID | Primary key                                                                                                   |
| `name` | TEXT | Name for a tag relationship type (e.g. "is a", "is associated with", "is a parent to", "is a predecessor to") |

## Tag Compositions

Allows atomic tags to be combined into composite phrases while preserving their grammatical structure.

### `tag_compositions` Table

Used to define composite phrases (e.g., `"very big red car"`), where the composite tag is treated as a **single** tag but internally linked to multiple atomic tags.

Each composite tag is stored in the `tags` table and linked to one or more component tags using this table. The `position` field determines the order of components (1-based index).

| Column             | Type | Description                                                                                                  |
| ------------------ | ---- | ------------------------------------------------------------------------------------------------------------ |
| `id`               | UUID | Primary key                                                                                                  |
| `base_tag_id`      |      | Foreign key to a [tag](./tags.md#tags-table), representing the **base tag** (e.g., `"very big red car"`)     |
| `component_tag_id` |      | Foreign key to a [tag](./tags.md#tags-table), representing an atomic **component** (e.g., `"very"`, `"big"`) |
| `position`         | INT  | Order of the component tag in the composite phrase (1-based index)                                           |

> **Base tags** are treated as atomic when composing higher-level phrases. For instance, once "**red car**" is stored as a tag in the [`tags`](./tags.md#tags-table) table, it may itself be used as a base to form "big ***red car***", which in turn may be used to form "very ***big red car***", and so on.
>
> In this way, while a composite tag may consist of many atomic components, each composition step always joins two tags: the left-hand "base" (which may itself be composite) and a new component.

> [!todo] TODO
>
> - Enforce uniqueness on `base_tag_id` + ordered list of `component_tag_ids`.
> - Prevent adding similar components to already composited tags, e.g., prevent adding `big` to `big car`.
> - Prevent semantic composites from being atomic: e.g., `"very big"` must only exist as a composite, not a standalone atomic tag.

### Composition Grammar Rules (Conceptual)

To prevent the creation of semantically invalid or grammatically incorrect composite tags, the system assumes a basic grammatical model:

- **Atomic tags** are labeled with part-of-speech (POS) roles (e.g., `adjective`, `noun`, `adverb`).
- **Composite tags** are created by combining a `base_tag` (e.g., a noun phrase) with a `component_tag` (e.g., an adjective or modifier), based on rules such as:
  - `adjective` + `noun` → ✅ `"red car"`
  - `adverb` + `adjective` + `noun` → ✅ `"very big car"`
  - `adverb` + `adverb` → ❌ invalid
  - `noun` + `noun` → ✅ only if meaningful (e.g., `"race car"`)

In the future, the system may optionally enforce POS compatibility using POS tags or lexical constraints, to avoid illogical composites like `"very big"` (an adverb + adjective phrase with no noun).

Composite tags that violate these patterns should:

- Be blocked at creation time,
- Or only be allowed as aliases pointing to a canonical tag (e.g., `"very big"` → alias for `"huge"`).

If you want to implement this later, you'd need:

1. A `part_of_speech` or `tag_type` column in tags, with values like `noun`, `adjective`, `adverb`, `modifier`, etc.
2. A rule engine or validation function that enforces:

```python
def is_valid_composition(base_pos: str, component_pos: str) -> bool:
    return (base_pos == "noun" and component_pos == "adjective") or \
           (base_pos == "adjective" and component_pos == "adverb") or \
           …
```
