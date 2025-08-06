# Delete an existing entity

## Parameters

- `id` (UUID, required): The ID of the entity to delete

## Request

```curl
curl -X DELETE http://localhost:8000/api/v1/entities/<entity_id> \
  -H "Authorization: Bearer <YOUR_SECRET_KEY>"
```

## Returns

### Response Attributes

- `status`: Confirmation string (e.g., "deleted")

### Response Body

```json
{
	"status": "deleted"
}
```

### Error Codes

| HTTP status code | Code      | Description      |
| ---------------- | --------- | ---------------- |
| 404              | not_found | Entity not found |
