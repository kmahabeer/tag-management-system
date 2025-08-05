# Entity Tagging

## `entity_tags` Join Table

Junction table which links `tags` to `entities`.

| Column      | Type  | Description                                                                                                                          |
| ----------- | ----- | ------------------------------------------------------------------------------------------------------------------------------------ |
| `id`        |       | Primary key; composite of `entity_id` and `tag_id`                                                                                   |
| `asset_id`  | UUID  | Foreign key to `assets`                                                                                                              |
| `tag_id`    | UUID  | Foreign key to `tags`                                                                                                                |
| `metadata` | JSONB | Arbitrary key-value data about the tag assignment (e.g., `{ "source": "Label Studio", "confidence": 0.92, "annotator": "user123" }`) |
