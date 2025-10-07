---
title: Tags
parent: Database Schema
has_children: true
nav_order: 2
---
# Tags

The Tag Management System uses a unified tag table, enriched by types, aliases, and compositional relationships.

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
