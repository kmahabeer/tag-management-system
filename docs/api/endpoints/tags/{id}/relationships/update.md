# Update tag relationships

## Parameters

- `id` (UUID, required): Tag A ID (the dominant tag)
- `relationships`: Array of objects containing:
  - `tag_b_id` (UUID, required): The related tag
  - `relationship_type_id` (UUID, required): Type of relationship
  - `description` (string, optional): Human-readable explanation

## Request

```curl
curl -X PATCH http://localhost:8000/api/v1/tags/<tag_id>/relationships \
  -H "Authorization: Bearer <YOUR_SECRET_KEY>" \
  -H "Content-Type: application/json" \
  -d '{
    "relationships": [
      {
        "tag_b_id": "UUID",
        "relationship_type_id": "UUID",
        "description": "Parent-child relationship"
      }
    ]
  }'
```

## Returns

### Response Attributes

- `relationships`: Updated relationship records

### Response Body

```json
{
	"relationships": [
		{
			"tag_a_id": "UUID",
			"tag_b_id": "UUID",
			"relationship_type_id": "UUID",
			"description": "Parent-child relationship"
		}
	]
}
```

### Error Codes

| HTTP status code | Code         | Description                |
| ---------------- | ------------ | -------------------------- |
| 400              | invalid_data | Invalid relationship input |
| 404              | not_found    | Tag not found              |
