---
title: Parts of Speech
parent: Miscellaneous
nav_order: 1
---
# Parts of Speech

Tags in the system can optionally be classified by their grammatical role using **part-of-speech** labels. This enables the construction of semantically valid composite tags (e.g., `"very big car"`) and enforcement of grammar rules during tag creation.

This document defines the `part_of_speech` reference table and its role in validating tag compositions.

## Purpose

The primary use of this table is to:

- Enforce compositional grammar rules (e.g., only combine an adverb with an adjective, then a noun).
- Help prevent semantically invalid tags such as `"very big"` (which lacks a noun).
- Enable intelligent filtering or generation of tag suggestions in the UI.

## `parts_of_speech` Table

| Field        | Type | Description                                             |
|--------------|------|---------------------------------------------------------|
| `id`         | UUID | Unique identifier                                       |
| `name`       | TEXT | Part of speech label (e.g., `noun`, `adjective`, `verb`) |
| `description`| TEXT | Explanation of the grammatical role                    |
| `is_active`  | BOOL | Indicates if the POS type is currently in use          |

### Recommended Values

| name        | description                                                   |
|-------------|---------------------------------------------------------------|
| `noun`      | A person, place, thing, or concept                            |
| `adjective` | A word that modifies a noun                                   |
| `adverb`    | A word that modifies an adjective, verb, or another adverb    |
| `verb`      | An action or state                                            |
| `modifier`  | General-purpose modifier (fallback)                           |
| `preposition` | A word showing relationship between elements (e.g., "under")|
| `conjunction` | A linking word such as "and", "or"                          |

## Integration

- Each tag may have an optional `part_of_speech_id` field referencing this table.
- The system may use this classification to validate:
	- Composite tag structure
	- Tag search disambiguation
	- Filtering logic by POS (e.g., show only adjectives)

## Example Usage

- `"very"` → `adverb`
- `"big"` → `adjective`
- `"car"` → `noun`

Composite tag: `"very big car"`
Validation rule: ✅ `adverb` → `adjective` → `noun`

Invalid tag: `"very big"`
Validation rule: ❌ missing required `noun` root
