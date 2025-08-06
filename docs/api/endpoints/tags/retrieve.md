# Retrieve a tag

Retrieves a tag.

## Parameters

- `id` (UUID, required): ID of the tag to retrieve

## Request

```curl
curl -X GET http://localhost:8000/api/v1/tags/<tag_id> \
  -H "Authorization: Bearer <YOUR_SECRET_KEY>"
```

## Returns

### Response Attributes

- `id` (UUID), `name`, `display_name`, `metadata`, `part_of_speech_id`, etc.
- `created_at`, `updated_at`

### Response Body

```json
{
	"id": "UUID",
	"name": "fast",
	"display_name": "Fast",
	"metadata": {},
	"created_at": "2025-08-06T12:00:00Z",
	"updated_at": "2025-08-06T12:00:00Z"
}
```

### Error Codes

| HTTP status code | Code      | Description   |
| ---------------- | --------- | ------------- |
| 404              | not_found | Tag not found |
