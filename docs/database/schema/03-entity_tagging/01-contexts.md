---
title: Contexts
parent: Entity-Tag Relationships
nav_order: 1
---
## Tagging Contexts

Tags can be assigned to entities in **different semantic roles**, depending on how they relate to the entity.  
For example, a tag like `"Dog"` could indicate that a dog is **depicted in** the image (content), that the image is **about** dogs (subject), or that the tag is part of **annotation metadata**.

To support this, the [`entity_tags`](../index.md#entity_tags-table) table includes a `context_id` field that captures the *purpose or meaning* of each tag assignment.

### `contexts` Table

| Column | Type | Description |
|--------|------|--------------|
| `id` | UUID | Primary key. |
| `name` | TEXT | Human-readable name of the tagging context (e.g., `"content"`, `"subject"`). |
| `classification_type` | TEXT | Either `"subjective"` or `"objective"`, defining how the context should be interpreted. |
| `description` | TEXT | Optional description explaining how this context is used. |
| `is_active` | BOOLEAN | Whether this context is active and assignable. |

## Purpose

The contexts table enables:

- Assigning the **same tag** to the **same entity** in multiple ways (e.g., tagging `"Obama"` as both `"subject"` and `"content"`).  
- **Filtering, displaying, or editing tags** differently depending on workflow, UI layout, or context-specific logic.  
- **Improved semantic clarity** for downstream systems interpreting tagging meaning.

The list of allowed contexts is defined and enforced at the schema level, and can be extended over time. See the table below for currently supported contexts.

## Taxonomy of Tagging Contexts

Contexts are grouped into two high-level classifications — **objective** and **subjective** — based on whether they describe factual characteristics or human interpretation.

### Objective Context Classifications

| Context | Description | Use Case | Example Tags |
|----------|-------------|-----------|---------------|
| `subject` | What the entity is **about** (themes, topics, high-level meaning). | A documentary image *about* climate issues. | `"Climate Change"` |
| `content` | What is **in** the entity (visual elements, objects, or people depicted). | A person or animal shown *in* a photo. | `"Person"`, `"Dog"` |
| `metadata` | Descriptive information about the file or process, not from content itself. | Indicates origin or status of the file. | `"Generated"`, `"Low Quality"` |

### Subjective Context Classifications

| Context | Description | Use Case | Example Tags |
|----------|-------------|-----------|---------------|
| `style` | Artistic or aesthetic style the entity belongs to. | Used for art, photography, or design. | `"Impressionism"`, `"Monochrome"` |
| `emotion` | Emotional tone or mood conveyed by the entity. | Used for classifying film stills or images by mood. | `"Sad"`, `"Joyful"` |
| `annotation` | Notes made by annotators or reviewers; subjective or observational. | Used in human-led labeling projects. | `"Crowd"`, `"Blurry"` |

## API Endpoint Summary

| Category | Endpoint | CRUD Coverage | Description |
|-----------|-----------|----------------|--------------|
| **Contexts** | `/contexts` | **GET**, **POST** | Retrieve or create semantic contexts used to group or interpret tags and entities. |
| | `/contexts/{id}` | **GET**, **PATCH**, **DELETE** | Retrieve, update, or delete a specific context definition. |

### CRUD Parity Overview

| Scope | CRUD | Description |
|--------|------|--------------|
| Contexts | ✅ Full | Manage the complete lifecycle of context definitions used across the tagging system. |

✅ = Full REST-compliant CRUD support.

## Summary

Contexts serve as the **semantic bridge** between entities, tags, and UI configurations. They define *how* a tag applies to an entity — as content, subject, emotion, or metadata — and ensure that meaning is explicit, structured, and queryable. This design enables flexible, human-readable tagging while preserving machine-interpretability and consistency across the entire system.
