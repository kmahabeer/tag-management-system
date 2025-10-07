---
title: Tags Compositions
---
# Tag Compositions

Allows atomic tags to be combined into composite phrases while preserving their grammatical structure.

## `tag_compositions` Table

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

## Composition Grammar Rules (Conceptual)

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
