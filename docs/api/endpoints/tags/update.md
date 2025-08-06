# Update a tag

## Parameters

- `id` (UUID, required): Tag ID to update
- `display_name` (string, optional): Updated UI label
- `metadata` (object, optional): Additional metadata
- `part_of_speech_id` (UUID, optional): Update POS reference

## Request

```curl
curl -X PATCH http://localhost:8000/api/v1/tags/<tag_id> \
  -H "Authorization: Bearer <YOUR_SECRET_KEY>" \
  -H "Content-Type: application/json" \
  -d '{
    "display_name": "Very Fast"
  }'
```

## Returns

### Response Attributes

- All tag fields including updated values

### Response Body

```json
{
	"id": "UUID",
	"name": "fast",
	"display_name": "Very Fast",
	"metadata": {},
	"updated_at": "2025-08-06T12:05:00Z"
}
```

### Error Codes

| HTTP status code | Code         | Description       |
| ---------------- | ------------ | ----------------- |
| 404              | not_found    | Tag not found     |
| 400              | invalid_data | Malformed request |
