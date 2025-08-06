# Delete a tag

## Parameters

- `id` (UUID, required): The ID of the tag to delete

## Request

```curl
curl -X DELETE http://localhost:8000/api/v1/tags/<tag_id> \
  -H "Authorization: Bearer <YOUR_SECRET_KEY>"
```

## Returns

### Response Attributes

- `status`: String confirmation (e.g., "deleted")

### Response Body

```json
{
	"status": "deleted"
}
```

### Error Codes

| HTTP status code | Code      | Description   |
| ---------------- | --------- | ------------- |
| 404              | not_found | Tag not found |
