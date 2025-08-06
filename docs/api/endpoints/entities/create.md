# Create a new entity

## Parameters

- `name` (string, required): Name of the entity
- `location` (string, optional): File path, URL, or URI
- `is_primary` (boolean, optional): Whether this is the primary version
- `metadata` (object, optional): Key-value data (e.g., `{ "source": "Label Studio" }`)

## Request

```curl
curl -X POST http://localhost:8000/api/v1/entities \
  -H "Authorization: Bearer <YOUR_SECRET_KEY>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Sketch A",
    "location": "images/sketch-a.jpg",
    "is_primary": true,
    "metadata": { "source": "Label Studio" }
  }'
```

## Returns

### Response Attributes

- `id`, `name`, `location`, `is_primary`, `metadata`
- `created_at`, `updated_at`

### Response Body

```json
{
	"id": "UUID",
	"name": "Sketch A",
	"location": "images/sketch-a.jpg",
	"is_primary": true,
	"metadata": { "source": "Label Studio" },
	"created_at": "2025-08-06T12:00:00Z",
	"updated_at": "2025-08-06T12:00:00Z"
}
```

### Error Codes

| HTTP status code | Code         | Description              |
| ---------------- | ------------ | ------------------------ |
| 400              | invalid_data | Missing or invalid input |
