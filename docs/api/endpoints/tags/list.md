# List all tags

Lists all tags.

## Parameters

- `limit` (int, optional): Number of results to return per page
- `offset` (int, optional): Offset for pagination

## Request

```curl
curl -X GET http://localhost:8000/api/v1/tags \
  -H "Authorization: Bearer <YOUR_SECRET_KEY>"
```

## Returns

### Response Attributes

- `results`: Array of tag objects
- `total`: Total count of matching tags
- Pagination metadata if supported (e.g., `next`, `previous`)

### Response Body

```json
{
	"results": [
		{
			"id": "UUID",
			"name": "fast",
			"display_name": "Fast",
			"metadata": {},
			"created_at": "2025-08-06T12:00:00Z",
			"updated_at": "2025-08-06T12:00:00Z"
		}
	],
	"total": 1
}
```

### Error Codes

| HTTP status code | Code | Description |
| ---------------- | ---- | ----------- |
| 200              | ok   |             |
