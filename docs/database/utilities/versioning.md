# Entity Versioning Logic

This document defines how versioning is handled within the `entities` table using the `entity_relationships` join table.

## Key Concepts

- **Primary Entity**: The canonical or "main" version of a digital artifact (e.g., the final render of an image).
- **Alternate Entity**: A variant, earlier version, or derived file from the primary entity.

## `entity_relationships` Rules

- `entity_a_id` always refers to the **source** or **primary** entity.
- `entity_b_id` refers to the **alternate** or **derived** entity.

This unidirectional model enables consistent traversal and grouping of related files.

### Examples

- `"Film Still A"` → `"alternate version of"` → `"Film Still B"`
- `"Sketch X"` → `"belongs to set"` → `"Study Group A"`

## Constraints

- A primary entity **can have many** alternates.
- An alternate entity **must not be reused** as the primary in another relationship — prevent circular versioning unless explicitly allowed.
- Deleting the primary entity should either:
	- Cascade delete its alternates **(if derived and non-critical)**, or
	- Be blocked unless all alternates are reassigned or deleted.

> [!NOTE] Future Consideration
> Support recursive resolution of all versions via a materialized view or recursive CTE for UI grouping and search optimization.
