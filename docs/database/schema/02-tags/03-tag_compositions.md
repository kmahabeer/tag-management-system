---
title: Tags Compositions
parent: Tags
nav_order: 3
---
# Tag Compositions

Allows atomic tags to be combined into structured composite phrases while preserving grammatical order and semantic meaning. Composite tags (e.g., `"very big red car"`) are stored as standalone entries in the [`tags`](../index.md#tags-table) table.

## `tag_compositions` Table

This table defines **how those composites are built** from smaller atomic components such as `"very"`, `"big"`, and `"car"`.

Used to define composite phrases (e.g., `"very big red car"`), where the composite tag is treated as a **single** tag but internally linked to multiple atomic tags.

Each composite tag is stored in the `tags` table and linked to one or more component tags using this table. The `position` field determines the order of components (1-based index).

| Column             | Type | Description                                                                                                  |
| ------------------ | ---- | ------------------------------------------------------------------------------------------------------------ |
| `id`               | UUID | Primary key                                                                                                  |
| `base_tag_id`      |      | Foreign key to a [tag](./tags.md#tags-table), representing the **base tag** (e.g., `"very big red car"`)     |
| `component_tag_id` |      | Foreign key to a [tag](./tags.md#tags-table), representing an atomic **component** (e.g., `"very"`, `"big"`) |
| `position`         | INT  | Order of the component tag in the composite phrase (1-based index)                                           |

### Compositional Structure

**Base tags** are treated as atomic when composing higher-level phrases.  
For instance, once `"red car"` is stored as a tag in the [`tags`](../index.md#tags-table), it may itself be used as a **base** to form `"big red car"`, which in turn may be used to form `"very big red car"`, and so on.

In this way, while a composite tag may contain multiple atomic components, **each composition step always joins two tags**:

```txt
(base_tag) + (component_tag) → (new_composite)
```

For example:

| Base | Component | Result |
|------|------------|---------|
| "car" | "red" | "red car" |
| "red car" | "big" | "big red car" |
| "big red car" | "very" | "very big red car" |

This recursive compositional model enables incremental phrase construction while preserving structural consistency.

> **Base tags** are treated as atomic when composing higher-level phrases. For instance, once "**red car**" is stored as a tag in the [`tags`](./tags.md#tags-table) table, it may itself be used as a base to form "big ***red car***", which in turn may be used to form "very ***big red car***", and so on.
>
> In this way, while a composite tag may consist of many atomic components, each composition step always joins two tags: the left-hand "base" (which may itself be composite) and a new component.

> [!todo] TODO
>
> - Enforce uniqueness on `base_tag_id` + ordered list of `component_tag_ids`.
> - Prevent adding similar components to already composited tags, e.g., prevent adding `big` to `big car`.
> - Prevent semantic composites from being atomic: e.g., `"very big"` must only exist as a composite, not a standalone atomic tag.

## Composition Grammar Rules (Conceptual)

To prevent illogical or ungrammatical composites, the system assumes a **part-of-speech (POS)**–aware grammatical model.  
This model can be optionally enforced via the [`parts_of_speech`](../../utilities/parts_of_speech.md) table.

### Basic POS Rules

| Pattern | Example | Valid |
|----------|----------|--------|
| adjective + noun | `"red car"` | ✅ |
| adverb + adjective + noun | `"very big car"` | ✅ |
| adverb + adverb | `"very extremely"` | ❌ |
| noun + noun | `"race car"` | ✅ (conditionally meaningful) |

### Future Enforcement Strategy

In future versions, this compositional model may include **POS compatibility validation** to automatically reject invalid constructions (e.g., `"very big"` without a noun).  
Composite tags violating POS rules could instead be represented as **aliases** for valid, atomic tags (e.g., `"very big"` → alias for `"huge"`).

If implemented, this would rely on two supporting features:

1. A `part_of_speech` or `tag_type` column in the `tags` table.  
2. A rule validation engine similar to:

```python
def is_valid_composition(base_pos: str, component_pos: str) -> bool:
    return (base_pos == "noun" and component_pos == "adjective") or \
           (base_pos == "adjective" and component_pos == "adverb") or \
           (base_pos == "noun" and component_pos == "noun")
```

## Constraints

To maintain data and semantic integrity:

- (`base_tag_id`, ordered `component_tag_ids`) must be unique.
- Prevent redundant additions (e.g., adding "`big`" to "`big car`").
- Composite tags must contain at least one `noun` (or `noun phrase`) as the semantic root.
- Prevent partial composites (e.g., "`very big`") from being stored as atomic tags.

## API Endpoint Summary

| Category                         | Endpoint                                      | CRUD Coverage                            | Description                                                             |
| -------------------------------- | --------------------------------------------- | ---------------------------------------- | ----------------------------------------------------------------------- |
| **Tag Compositions**             | `/tags/{id}/compositions`                     | **GET**, **POST**, **PATCH**, **DELETE** | Manage component tags that form composite phrases for a given base tag. |
|                                  | `/tags/{id}/compositions/{composition_id}`    | **GET**, **PATCH**, **DELETE**           | Inspect, modify, or remove a specific composition record.               |
| **System-Wide Tag Compositions** | `/tag-compositions`, `/tag-compositions/{id}` | **GET**, **POST**, **PATCH**, **DELETE** | Manage all tag composition records globally.                            |

### CRUD Parity Overview

| Scope       | CRUD   | Description                                  |
| ----------- | ------ | -------------------------------------------- |
| Per Tag     | ✅ Full | Manage component tags for a single base tag. |
| System-Wide | ✅ Full | Manage all composition records across tags.  |

✅ = Full REST-compliant CRUD support.

## Summary

Tag compositions provide a structured, linguistic foundation for building multi-word tags that remain machine-interpretable and semantically consistent. They bridge lexical hierarchy with relational data, enabling grammatically-aware search, tagging, and inference throughout the system.
