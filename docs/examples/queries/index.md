# Structured Inspector JSON

For **asset 1** (Fargo still):

```json
{
  "asset_id": 1,
  "filename": "fargo_still_01.jpg",
  "available_purposes": ["Film Still", "Movie reference"],
  "selected_purpose": "Film Still",
  "movie": {
    "title": "Pulp Fiction",
    "director": ["Quentin Tarantino", "Joel Coen", "Ethan Coen"],
    "writer": ["Quentin Tarantino", "Joel Coen", "Ethan Coen"],
    "actors": [
      "William H. Macy",
      "Frances McDormand",
      "Steve Buscemi",
      "Samuel L. Jackson"
    ],
    "genre": ["Dark Comedy", "Crime", "Drama"]
  },
  "content_tags": {
    "subject": [
      "William H. Macy",
      "Frances McDormand",
      "Steve Buscemi",
      "Samuel L. Jackson"
    ],
    "characters": ["Vincent Vega", "Jules Winnfield"],
    "costumes": ["Black Suit"],
    "props": ["Gun"]
  },
  "other_tags": []
}
```

For **asset 3** (Hateful Eight still):

```json
{
  "asset_id": 3,
  "filename": "hateful_eight_table.jpg",
  "available_purposes": ["Film Still"],
  "selected_purpose": "Film Still",
  "movie": {
    "title": "The Hateful Eight",
    "director": ["Quentin Tarantino"],
    "writer": ["Quentin Tarantino"],
    "actors": ["Kurt Russell", "Samuel L. Jackson"],
    "genre": ["Western"]
  },
  "content_tags": {
    "subject": ["Kurt Russell", "Samuel L. Jackson"],
    "props": ["Table"],
    "setting": ["Winter"]
  },
  "other_tags": []
}
```

For **asset 4** (Car photo):

```json
{
  "asset_id": 4,
  "filename": "datsun_240z.jpg",
  "available_purposes": ["Figure Study", "Movie reference"],
  "selected_purpose": "Figure Study",
  "content_tags": {
    "subject": ["Datsun 240z"],
    "vehicle": ["Car", "Vehicle"]
  },
  "other_tags": []
}
```

For **asset 5** (Figure study):

```json
{
  "asset_id": 5,
  "filename": "figure_study_001.jpg",
  "available_purposes": ["Figure Study"],
  "selected_purpose": "Figure Study",
  "content_tags": {
    "subject": ["Model"],
    "pose": [],
    "lighting": [],
    "setting": []
  },
  "other_tags": []
}
```

## SQL Query Examples

1. **Fetch asset’s tags for a given purpose**

```sql
SELECT t.name, tt.name AS type
FROM asset_purposes ap
JOIN assets a        ON a.id = ap.asset_id
JOIN asset_tags at   ON at.asset_id = a.id
JOIN tags t          ON t.id = at.tag_id
JOIN tag_tag_types ttt ON t.id = ttt.tag_id
JOIN tag_types tt    ON tt.id = ttt.tag_type_id
WHERE a.id = 1 AND ap.is_primary = TRUE
ORDER BY tt.name, t.name;
```

1. **Assemble movie section (child tags of the movie tag)**

```sql
WITH movie_tag AS (
  SELECT tag_id FROM asset_tags WHERE asset_id = 1 AND tag_id IN (SELECT id FROM tags WHERE name = 'Pulp Fiction')
)
SELECT rel.relationship, child.name
FROM movie_tag m
JOIN tag_relationships rel ON rel.parent_tag_id = m.tag_id
JOIN tags child           ON child.id = rel.child_tag_id;
```

1. **List available purposes for an asset**

```sql
SELECT t.name, ap.is_primary
FROM asset_purposes ap
JOIN tags t ON t.id = ap.purpose_tag_id
WHERE ap.asset_id = 1;
```

1. **Fetch uncategorized tags (other_tags)**

```sql
-- Assuming you have logic to collect categorized tag_ids
WITH categorized AS (
  SELECT tag_id FROM asset_tags WHERE asset_id = 1 AND tag_id IN (
    -- list of tag_ids used in 'movie' or 'content_tags'
  )
)
SELECT t.name
FROM asset_tags at
JOIN tags t ON t.id = at.tag_id
WHERE at.asset_id = 1
  AND at.tag_id NOT IN (SELECT tag_id FROM categorized);
```
