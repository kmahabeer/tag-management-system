# Retrieve an existing entity

## Parameters

- `id` (UUID, required): ID of the entity to retrieve

## Request

```curl
curl -X GET http://localhost:8000/api/v1/entities/<entity_id> \
  -H "Authorization: Bearer <YOUR_SECRET_KEY>"
```

## Returns

### Response Attributes

- `id`, `name`, `location`, `is_primary`, `metadata`, `created_at`, `updated_at`

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

| HTTP status code | Code      | Description      |
| ---------------- | --------- | ---------------- |
| 404              | not_found | Entity not found |
