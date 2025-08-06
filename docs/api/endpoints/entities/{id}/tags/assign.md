# List tags assigned to an entity

## Parameters

- `id` (UUID, required): Entity ID

## Request

```curl
curl -X GET http://localhost:8000/api/v1/entities/<entity_id>/tags \
  -H "Authorization: Bearer <YOUR_SECRET_KEY>"
```

## Returns

### Response Attributes

- `tags`: Array of tag assignments with context and metadata

### Response Body

```json
{
	"tags": [
		{
			"tag": {
				"id": "UUID",
				"name": "car",
				"display_name": "Car"
			},
			"context": "content",
			"metadata": {
				"source": "Label Studio",
				"confidence": 0.98
			}
		}
	]
}
```

### Error Codes

| HTTP status code | Code      | Description      |
| ---------------- | --------- | ---------------- |
| 404              | not_found | Entity not found |
