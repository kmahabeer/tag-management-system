---
title: Tags Relationships
---
# Tag Relationships

Defines how tags are connected hierarchically or associatively, enabling semantic linkages across the tag graph.

## `tag_relationships` Table

Many-to-many join table linking [`tags`](tags.md#tags-table) in either **hierarchical** or **associative** ways. Relationships define how tags are connected or composed. The first tag, **Tag A**, represents the more "dominant" tag to **Tag B**; in other words, if the relationship between two tags is "Parent-Child", then **Tag A** is the Parent and **Tag B** is the child.

| Column                 | Type | Description                                                                                                      |
| ---------------------- | ---- | ---------------------------------------------------------------------------------------------------------------- |
| `id`                   | UUID | Primary key                                                                                                      |
| `tag_a_id`             |      | Foreign key to [**Tag A**](./tags.md#tags-table)                                                                 |
| `tag_b_id`             |      | Foreign key to [**Tag B**](./tags.md#tags-table)                                                                 |
| `relationship_type_id` |      | Foreign key to [`tag_relationship_types`](./tags.md#tag_relationship_types-table) table                          |
| `description`          | TEXT | Description explaining the relationship (e.g., `"{Tag A} is a {Tag B}"`, `"{Tag A} is associated with {Tag B}"`) |

The description allows systems to render a human-readable explanation of the relationship. Example: `"Director"` → `role-of` → `"Quentin Tarantino"`

## What it means semantically

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
