# List all entities

## Parameters

- `limit` (int, optional): Number of results per page
- `offset` (int, optional): Pagination offset

## Request

```curl
curl -X GET http://localhost:8000/api/v1/entities \
  -H "Authorization: Bearer <YOUR_SECRET_KEY>"
```

## Returns

### Response Attributes

- `results`: Array of entity objects
- `total`: Total number of entities

### Response Body

```json
{
	"results": [
		{
			"id": "UUID",
			"name": "Sketch A",
			"location": "images/sketch-a.jpg",
			"is_primary": true,
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
| 200              | ok   | Success     |
