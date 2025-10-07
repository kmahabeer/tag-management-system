---
title: Entity-Tag Relationships
---
# Entity Tagging

## `entity_tags` Table

Many-to-many join table which links [`tags`](./tags.md) to [`entities`](./entities.md).

| Column       | Type  | Description                                                                                                                          |
| ------------ | ----- | ------------------------------------------------------------------------------------------------------------------------------------ |
| `id`         | UUID  | Primary key                                                                                                                          |
| `entity_id`  |       | Foreign key to the [`entities`](./entities.md) table                                                                                 |
| `tag_id`     |       | Foreign key to the [`tags`](./tags.md) table                                                                                         |
| `context_id` |       | Foreign key to the [`contexts`](ui_configurations.md#contexts) table                                                                 |
| `metadata`   | JSONB | Arbitrary key-value data about the tag assignment (e.g., `{ "source": "Label Studio", "confidence": 0.92, "annotator": "user123" }`) |

`Metadata` in this context informs about how the tag is used on an entity.

## Tagging Contexts

Tags can be assigned to entities in **different semantic roles**, depending on how they relate to the entity. For example, a tag like `"Dog"` could indicate that a dog is **depicted in** the image (content), that the image is **about** dogs (subject), or that the tag is part of **annotation metadata**.

To support this, the `entity_tags` table includes a `context` field that captures the *purpose or meaning* of each tag assignment.

### `contexts` Table

| Column                | Type    | Description                 |
| --------------------- | ------- | --------------------------- |
| `id`                  | UUID    | Primary key                 |
| `name`                | TEXT    |                             |
| `classification_type` |         | `subjective` or `objective` |
| `description`         | TEXT    |                             |
| `is_active`           | BOOLEAN |                             |

This enables:

- Assigning the **same tag** to the **same entity** in multiple ways (e.g. an image of Obama having the tag `"Obama"` as both "content" and "subject").
- **Filtering, displaying, or editing tags** differently depending on the workflow or UI configuration.
- **Improved semantic clarity** in how tags are interpreted by downstream systems.

The list of allowed contexts is defined and enforced at the schema level, and can be extended over time. See the table below for currently supported contexts. The following table defines the **Taxonomy of Tagging Contexts**, grouped by *objective* and *subjective* classifications.

### Objective Context Classifications

| Context    | Description                                                                 | Use Case                                   | Example Tags               |
| ---------- | --------------------------------------------------------------------------- | ------------------------------------------ | -------------------------- |
| `subject`  | What the entity is **about**<br>(e.g., themes, topics, high-level meaning)  | A documentary image *about* climate issues | "Climate Change"           |
| `content`  | What is **in** the entity<br>(e.g., visual elements, objects, people shown) | A person or animal shown *in* a photo      | "Person", "Dog"            |
| `metadata` | Descriptive info not from the content, but about the file or process        | Indicates origin or status info            | "Generated", "Low Quality" |

### Subjective Context Classifications

| Context      | Description                                                   | Use Case                                  | Example Tags                  |
| ------------ | ------------------------------------------------------------- | ----------------------------------------- | ----------------------------- |
| `style`      | Artistic or aesthetic style the entity belongs to             | Art, photography, or design               | "Impressionism", "Monochrome" |
| `emotion`    | Emotional tone conveyed by the entity                         | Classifying mood of a film still or image | "Sad", "Joyful"               |
| `annotation` | Notes made by an annotator; subjective or observational input | For human-led review or labeling projects | "Crowd", "Blurry"             |
