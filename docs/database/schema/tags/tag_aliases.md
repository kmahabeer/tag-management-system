---
title: Tags Aliases
---
# Tag Aliases

Represents alternative names or synonyms that point to a canonical tag, aiding search and tagging flexibility.

## `tag_aliases` Table

Alternative names or synonyms for a tag. Useful for search, filtering, and UI flexibility.

| Column   | Type | Description                                           |
| -------- | ---- | ----------------------------------------------------- |
| `id`     | UUID | Primary key                                           |
| `name`   | TEXT | Unique name for a tag's alias                         |
| `tag_id` |      | Foreign key to the [`tags`](tags.md#tags-table) table |
