# Labeled Data Import Logic

This document describes the logic for importing labeled data from external tools like **Label Studio** into the **Tagging Service**. It covers normalization, validation, enrichment, and persistence of tag-related data, as well as the handling of source metadata.

- All imported tags should be canonicalized and linked to entities
- Source metadata should be stored in `*_metadata` tables
- Composite or enriched tags can be generated after ingest
- All logic should be handled in the application layer before data persistence

> This approach maintains data integrity while preserving provenance and enabling traceability from external labeling tools.

## 1. Ingest Raw Data

**Input:** JSON export or webhook payload from an external annotation tool (e.g., Label Studio)

### Example:

```json
{
  "data": {
    "image": "https://example.com/image.jpg"
  },
  "annotations": [
    {
      "result": [
        {
          "value": {
            "labels": ["Red", "Car"]
          },
          "from_name": "label",
          "to_name": "image",
          "type": "choices"
        }
      ]
    }
  ]
}
```

### Action:

- Parse out the entity reference (e.g., image URL or ID)
- Parse out the label values (e.g., "Red", "Car")
- Capture any available metadata:
	- Tool name (`label_studio`)
	- Task ID / Annotation ID
	- User who labeled it
	- Confidence (if available)
	- Timestamp

## 2. Normalize & Validate Tags

**Purpose:** Ensure that incoming tags map cleanly to your system’s canonical tag structure.

### Steps:

- Convert labels to lowercase
- Strip whitespace / punctuation
- Check for existing tag match by name or alias
	- If found → use existing tag `id`
	- If not found → create new tag in `tags`
- (Optional) Infer tag type if not provided
	- e.g., "red" → `Color`, "car" → `Object`

## 3. Enrich with Metadata

Attach relevant source metadata to each imported tag or entity.

### Metadata fields to consider:

|Field|Applies to|Example value|
|---|---|---|
|`source`|tag / entity|`label_studio`|
|`annotated_by`|tag|`user@example.com`|
|`annotation_id`|tag / entity|`789`|
|`confidence`|tag|`0.93`|
|`timestamp`|tag / entity|`2025-08-05T13:45:00Z`|

**Where to store it:**

- Store metadata fields within the `metadata` JSONB column within respective tables.

### Metadata Validation (Required)

Before persisting metadata to the database, it must be validated for consistency and correctness.

#### Validation Rules:

- Only **known keys** should be accepted (or warn if using open schema)
- Required keys (e.g., `source`, `timestamp`) must be present
- Data types should be enforced:
  - `confidence` must be a float between `0.0` and `1.0`
  - `timestamp` must be ISO-8601 parseable
  - `annotation_id` should be a string or UUID
- Reject invalid structures, malformed payloads, or unsupported nested objects

#### Example (Python pseudocode):

```python
def validate_metadata(meta):
    required_keys = ["source", "timestamp"]
    if not all(k in meta for k in required_keys):
        raise ValueError("Missing required metadata field")

    if "confidence" in meta and not (0 <= float(meta["confidence"]) <= 1):
        raise ValueError("Confidence must be between 0 and 1")

    if "timestamp" in meta:
        datetime.fromisoformat(meta["timestamp"])  # raises error if invalid
```

Validating metadata ensures predictable behavior in downstream queries and improves traceability across external systems.

## 4. Store Tags and Relationships

### 1. Insert entity

```sql
INSERT INTO entities (id, name, location, source) VALUES (...);
```

### 2. Insert or upsert tags

```sql
INSERT INTO tags (id, name, display_name) VALUES (...)
ON CONFLICT (name) DO UPDATE SET updated_at = CURRENT_TIMESTAMP;
```

### 3. Link tags to entity

```sql
INSERT INTO entity_tags (entity_id, tag_id) VALUES (...);
```

### 4. Optionally insert metadata

```sql
INSERT INTO tag_metadata (tag_id, key, value) VALUES ('uuid_red', 'source', 'label_studio');
```

## 5. Post-processing (Optional)

You may wish to:

- Trigger tag enrichment (e.g., auto-assign `tag_types`)
- Compose multi-part tags (e.g., "red" + "car" → `"red car"` composite tag)
- Flag low-confidence tags for review
- Apply purpose heuristics to auto-tag `entity_purposes`

## 6. Query for Imported Tags

```sql
SELECT t.name, tm.value AS source
FROM tags t
JOIN tag_metadata tm ON t.id = tm.tag_id
WHERE tm.key = 'source' AND tm.value = 'label_studio';
```
