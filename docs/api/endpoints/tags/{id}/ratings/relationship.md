# Rate tag relationship

## Parameters

- `id` (UUID, required): Tag A ID
- `relationship_ratings`: List of relationship ratings

  - `tag_b_id` (UUID, required): Related tag
  - `context_id` (UUID, required): Context of the rating
  - `rating_id` (UUID, required): Rating score

## Request

```curl
curl -X PATCH http://localhost:8000/api/v1/tags/<tag_id>/relationship_ratings \
  -H "Authorization: Bearer <YOUR_SECRET_KEY>" \
  -H "Content-Type: application/json" \
  -d '{
    "relationship_ratings": [
      {
        "tag_b_id": "UUID",
        "context_id": "UUID",
        "rating_id": "UUID"
      }
    ]
  }'
```

## Returns

### Response Attributes

- `tag_a_id`: UUID
- `ratings`: List of relationship rating records

### Response Body

```json
{
	"tag_a_id": "UUID",
	"ratings": [
		{
			"tag_b_id": "UUID",
			"context_id": "UUID",
			"rating_id": "UUID"
		}
	]
}
```

### Error Codes

| HTTP status code | Code         | Description                  |
| ---------------- | ------------ | ---------------------------- |
| 400              | invalid_data | Missing or invalid fields    |
| 404              | not_found    | Tag or related tag not found |
