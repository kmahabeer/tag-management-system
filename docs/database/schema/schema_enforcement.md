# Schema Enforcement Rules

This document defines integrity constraints, validations, and behavioral rules that enforce semantic correctness across the Tagging Service database.

## Tag Composition Constraints

- A composite tag (e.g., "very big red car") must be composed of valid atomic tags.
- Composite tags should not themselves be used as component tags of other composites unless the base is well-formed.
- A composite tag must include exactly one noun (or noun phrase) as the semantic root.
- Disallow semantically invalid compositions (e.g., adverb + adverb).
- If needed, add a `part_of_speech` classification to the `tags` table to validate compositions using a rule engine.

## Tag Relationship Constraints

- Relationships should be semantically coherent based on their type:
	- `"is a"` and `"parent"` should enforce directionality.
	- Tag A must be more general than Tag B.
- Use a trigger or check function to enforce dominant/directed relationships.

## Purpose Tag Constraints

- In the `entity_purposes` table:
	- `purpose_tag_id` must reference a tag of type `"Purpose"`.
	- Only one row per entity should have `is_primary = TRUE`.
- Optionally enforce via SQL `CHECK` constraint or trigger.

## Unique Composite Tags

- Ensure uniqueness of compositions by enforcing:
	- `base_tag_id` + ordered `component_tag_ids` must be unique.
	- Prevent equivalent semantic composites from duplicating (e.g., `"very big car"` and `"huge car"` can be aliases, but not both standalone).

## Rating Type Normalization

- Use `is_normalized` in `rating_types` to enforce whether ratings must be in [0–1] or [1–10] range.
- Optionally enforce using a CHECK constraint on `score` column in `ratings`.

## Entity Relationships

- In `entity_relationships`:
	- `entity_a_id` must refer to the primary or source entity.
	- Deletion of a primary entity should cascade appropriately or be blocked.
