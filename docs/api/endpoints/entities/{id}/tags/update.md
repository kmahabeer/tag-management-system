# Assign or unassign tags to an entity

## Parameters

- `id` (UUID, required): Entity ID
- `tags`: Array of tag assignments, each containing:

  - `tag_id` (UUID, required)
  - `context_id` (UUID, required)
  - `metadata` (object, optional)

## Request

```curl
curl -X PATCH http://localhost:8000/api/v1/entities/<entity_id>/tags \
  -H "Authorization: Bearer <YOUR_SECRET_KEY>" \
  -H "Content-Type: application/json" \
  -d '{
    "tags": [
      {
        "tag_id": "UUID",
        "context_id": "UUID",
        "metadata": { "confidence": 0.95 }
      }
    ]
  }'
```

## Returns

### Response Attributes

- `status`: Confirmation string (e.g., "updated")

### Response Body

```json
{
	"status": "updated"
}
```

### Error Codes

| HTTP status code | Code         | Description              |
| ---------------- | ------------ | ------------------------ |
| 400              | invalid_data | Invalid or missing input |
| 404              | not_found    | Entity not found         |
