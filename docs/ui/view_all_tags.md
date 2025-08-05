# SQL View: View all tags

Here is a unified SQL View that combines both:

- Atomic tags (from the tags table), and
- Composite/modified tags (from the `tag_modifiers` table, joined to `tags` twice)

This is ideal for exposing a consistent tag list to your UI layer, whether for dropdowns, chips, or autocomplete fields.

```sql
CREATE OR REPLACE VIEW view_all_tags AS
SELECT
    t.id AS id,
    t.name AS name,
    NULL AS modifier_tag_id,
    NULL AS base_tag_id,
    'atomic' AS type
FROM
    tags t

UNION ALL

SELECT
    gen_random_uuid() AS id, -- optional: if you need a fake UUID for composite tags
    mod_tag.name || ' ' || base_tag.name AS name,
    tm.modifier_tag_id,
    tm.base_tag_id,
    'composite' AS type
FROM
    tag_modifiers tm
JOIN tags mod_tag ON tm.modifier_tag_id = mod_tag.id
JOIN tags base_tag ON tm.base_tag_id = base_tag.id;

```

## Sample results

| id (uuid) | name       | modifier_tag_id | base_tag_id | type      |
| --------- | ---------- | --------------- | ----------- | --------- |
| `uuid1`   | Black      | `NULL`          | `NULL`      | atomic    |
| `uuid2`   | Suit       | `NULL`          | `NULL`      | atomic    |
| `uuid3`   | Black Suit | `uuid1`         | `uuid2`     | composite |

> **Note**: If you want real UUIDs for composite tags, you can materialize the view into a table and assign persistent IDs. Otherwise, use `gen_random_uuid()` for ephemeral use.

## How to Use in UI/API

You can use this view as the source for tag suggestions:

```sql
SELECT id, name, type, modifier_tag_id, base_tag_id
FROM view_all_tags
ORDER BY name;
```

And in your API response:

```json
[
  { "id": "uuid1", "name": "Black", "type": "atomic" },
  { "id": "uuid2", "name": "Suit", "type": "atomic" },
  { "id": "uuid3", "name": "Black Suit", "type": "composite", "modifier_tag_id": "uuid1", "base_tag_id": "uuid2" }
]
```
