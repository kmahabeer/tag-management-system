# Rate tag in context

## Parameters

- `id` (UUID, required): Tag ID
- `contextual_ratings`: List of ratings to apply per context

  - `context_id` (UUID, required): Context of the rating
  - `rating_id` (UUID, required): Rating value
  - `user_id` (UUID, optional): Who submitted the rating (if needed)

## Request

```curl
curl -X PATCH http://localhost:8000/api/v1/tags/<tag_id>/ratings \
  -H "Authorization: Bearer <YOUR_SECRET_KEY>" \
  -H "Content-Type: application/json" \
  -d '{
    "contextual_ratings": [
      {
        "context_id": "UUID",
        "rating_id": "UUID",
        "user_id": "UUID"
      }
    ]
  }'
```

## Returns

### Response Attributes

- `tag_id`: UUID
- `ratings`: List of context ratings

### Response Body

```json
{
	"tag_id": "UUID",
	"ratings": [
		{
			"context_id": "UUID",
			"rating_id": "UUID",
			"user_id": "UUID"
		}
	]
}
```

### Error Codes

| HTTP status code | Code         | Description               |
| ---------------- | ------------ | ------------------------- |
| 400              | invalid_data | Missing or invalid fields |
| 404              | not_found    | Tag or context not found  |
