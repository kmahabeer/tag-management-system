# View tag relationships

## Parameters

- `id` (UUID, required): Tag ID

## Request

```curl
curl -X GET http://localhost:8000/api/v1/tags/<tag_id>/relationships \
  -H "Authorization: Bearer <YOUR_SECRET_KEY>"
```

## Returns

### Response Attributes

- `relationships`: List of tag relationship records

	- `tag_a_id`, `tag_b_id`, `relationship_type_id`, `description`

### Response Body

```json
{
	"relationships": [
		{
			"tag_a_id": "UUID",
			"tag_b_id": "UUID",
			"relationship_type_id": "UUID",
			"description": "Tag A is a parent of Tag B"
		}
	]
}
```

### Error Codes

| HTTP status code | Code      | Description   |
| ---------------- | --------- | ------------- |
| 404              | not_found | Tag not found |
