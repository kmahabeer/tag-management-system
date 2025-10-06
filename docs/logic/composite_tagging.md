# Composite Tagging

This document outlines how composite tags (e.g., "very big red car") are constructed, validated, stored, and queried within the Tagging Service. Composite tags are treated as first-class tags and are stored in the `tags` table with UUIDs like any other tag.

The composition logic is implemented in the application layer. The database is responsible for persistence and lookup.

## 1. Phrase Construction (Application Layer)

**Purpose:** Assemble a composite tag from multiple atomic tags in a user-defined order.

### Inputs

- A list of `tag_id`s representing atomic tags (e.g., `["uuid_very", "uuid_big", "uuid_red", "uuid_car"]`)

### Process

- Fetch tag names from the database for each `tag_id`
- Normalize and order them by intended composition sequence
- Concatenate names with spaces (e.g., `"very big red car"`)
- Optionally: construct `display_name` (e.g., `"Very Big Red Car"`)

### Output

- `composite_tag_name`: `"very big red car"`
- `display_name`: `"Very Big, Red Car"`
- `component_tag_ids`: ordered list of atomic tag UUIDs

## 2. Phrase Validation (Application Layer)

**Purpose:** Ensure that the proposed composition is semantically valid and consistent with existing business rules.

### Validation Steps

- Ensure at least two components
- Ensure tag names are not duplicated within the same composition
- (Optional) Enforce logical order, e.g., modifier → adjective → noun
- (Optional) Reject certain compositions (e.g., `"very car"`)
- Ensure that each tag ID is valid and exists in the `tags` table

### Output

- Boolean success/failure
- Reason message if invalid

## 3. Duplicate Check (Application Layer + DB Constraint)

**Purpose:** Prevent creation of the same composite tag more than once.

### Process

- Generate a normalized composite name (e.g., lowercase)
- Query the `tags` table for an exact match on `name`
- If found, return existing `tag_id`

### Example SQL

```sql
SELECT id FROM tags WHERE name = 'very big red car';
```

### Constraint (Optional)

```sql
ALTER TABLE tags ADD CONSTRAINT unique_tag_name UNIQUE(name);
```

## 4. Data Persistence (SQL)

**Purpose:** Save a new composite tag and its structure to the database.

### Steps

1. Insert composite into `tags`

```sql
INSERT INTO tags (id, name, display_name) VALUES (uuid, 'very big red car', 'Very Big Red Car');
```

1. Insert parts into `tag_compositions`

```sql
INSERT INTO tag_compositions (composite_tag_id, component_tag_id, position)
VALUES
  ('uuid_composite', 'uuid_very', 1),
  ('uuid_composite', 'uuid_big', 2),
  ('uuid_composite', 'uuid_red', 3),
  ('uuid_composite', 'uuid_car', 4);
```

## 5. Lookup / Query (SQL or View)

**Purpose:** Retrieve composite tags and their parts for display or search.

### Query all composite tags

```sql
SELECT id, name, display_name FROM tags
WHERE id IN (SELECT DISTINCT composite_tag_id FROM tag_compositions);
```

### Get parts of a composite tag

```sql
SELECT c.position, t.name
FROM tag_compositions c
JOIN tags t ON c.component_tag_id = t.id
WHERE c.composite_tag_id = 'uuid_composite'
ORDER BY c.position;
```

### Optional: Create a view

```sql
CREATE VIEW view_composite_tags AS
SELECT
  c.composite_tag_id,
  STRING_AGG(t.name, ' ' ORDER BY c.position) AS composition
FROM tag_compositions c
JOIN tags t ON c.component_tag_id = t.id
GROUP BY c.composite_tag_id;
```
