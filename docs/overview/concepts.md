# Concepts

The Tagging Service powers dynamic metadata assignment for creative assets. Tags are:

- **Reusable**: The same tag (e.g., "Quentin Tarantino") may apply to multiple assets or entities.
- **Typed**: Tags are classified (e.g., as a Person, Role, Vehicle).
- **Hierarchical**: Tags can be related to other tags (e.g., "Car" → "Datsun 240z").
- **Purpose-aware**: Each asset may serve multiple purposes (e.g., a film still and a reference photo), and the tag rendering changes accordingly.
- **Modifier-compatible**: Tags may be augmented by other tags (e.g., "very" + "tall").

The service uses a unified tag table, a set of tag types, and relationships to power rich, composable metadata logic.
