# Create a new tag

Creates a new tag.

## Parameters

- `name` (string, required): Canonical name for the tag
- `display_name` (string, optional): Human-readable label used in UI
- `part_of_speech_id` (UUID, optional): FK to `parts_of_speech`
- `metadata` (object, optional): Arbitrary metadata (e.g., `{ "source": "Label Studio" }`)

## Request

```curl
curl -X POST http://localhost:8000/api/v1/tags \
  -H "Authorization: Bearer <YOUR_SECRET_KEY>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "fast",
    "display_name": "Fast",
    "part_of_speech_id": "UUID",
    "metadata": { "source": "admin" }
  }'
```

## Returns

### Response Attributes

- `id` (UUID): Unique identifier
- `name`, `display_name`, `metadata`, `part_of_speech_id`, etc.
- `created_at`, `updated_at`

### Response Body

```json
{
	"id": "UUID",
	"name": "fast",
	"display_name": "Fast",
	"metadata": { "source": "admin" },
	"part_of_speech_id": "UUID",
	"created_at": "2025-08-06T12:00:00Z",
	"updated_at": "2025-08-06T12:00:00Z"
}
```

### Error Codes

| HTTP status code | Code         | Description                     |
| ---------------- | ------------ | ------------------------------- |
| 400              | invalid_data | Missing or malformed parameters |
| 409              | conflict     | Tag with the same name exists   |
