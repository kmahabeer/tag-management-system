---
title: Schema Constraints
parent: Database Schema
nav_order: 6
---
# Schema Enforcement Rules

Defines integrity constraints, validations, and behavioral rules ensuring semantic correctness and data consistency across the schema.

## Tag Composition Constraints

- A composite tag must be composed of valid atomic tags.  
- Composite tags may reference other composites as bases but must form valid grammatical structures.  
- A composite must include exactly one noun (or noun phrase) as the root.  
- Enforce uniqueness on (`base_tag_id`, ordered `component_tag_ids`).  

## Tag Relationship Constraints

- Directional relationships (e.g., “is a,” “parent”) must maintain logical dominance (`tag_a_id` > `tag_b_id`).  
- Tag A cannot be both parent and child of Tag B.  
- Enforce referential integrity between `tag_relationships` and `tag_relationship_types`.  

## Purpose Tag Constraints

- Each `entity_purpose` must reference a tag classified as type “Purpose.”  
- Only one purpose per entity may have `is_primary = TRUE`.  

## Unique Composite Tags

- Ensure uniqueness of compositions by enforcing:
	- `base_tag_id` + ordered `component_tag_ids` must be unique.
	- Prevent equivalent semantic composites from duplicating (e.g., `"very big car"` and `"huge car"` can be aliases, but not both standalone).

## Rating Normalization

- `is_normalized` in `rating_types` enforces whether associated `ratings.score` must be in `[0–1]` or `[1–10]`.  
- Use a database `CHECK` constraint or validation trigger to enforce consistency.  

## Entity Relationship Rules

- `entity_a_id` must reference the source (primary) entity.  
- `entity_a_id` ≠ `entity_b_id`.  
- The combination (`entity_a_id`, `entity_b_id`, `relationship_type_id`) must be unique.  
