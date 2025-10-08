---
title: Parts of Speech
parent: Utilities
nav_order: 1
---
# Parts of Speech

Tags in the system can optionally be classified by their grammatical role using **part-of-speech (POS)** labels.  
This enables validation of semantically valid composite tags (e.g., `"very big car"`) and ensures grammatically coherent tag creation.

## Purpose

The `parts_of_speech` table provides grammatical metadata that helps:

- Enforce compositional grammar rules (e.g., only combine an *adverb* with an *adjective*, then a *noun*).  
- Prevent semantically invalid tags such as `"very big"` (missing a noun).  
- Enable grammar-aware autocomplete, tag filtering, and intelligent UI suggestions.  

Each tag may optionally reference a part of speech via its `part_of_speech_id` column.

## `parts_of_speech` Table

| Field | Type | Description |
|--------|------|--------------|
| `id` | UUID | Primary key. |
| `name` | TEXT | Part of speech label (e.g., `"noun"`, `"adjective"`, `"verb"`). |
| `description` | TEXT | Explanation of the grammatical role. |
| `is_active` | BOOLEAN | Whether this part of speech is active and valid for tagging. |

### Recommended Values

| name        	| description                                                   |
|---------------|---------------------------------------------------------------|
| `noun`      	| A person, place, thing, or concept                            |
| `adjective` 	| A word that modifies a noun                                   |
| `adverb`    	| A word that modifies an adjective, verb, or another adverb    |
| `verb`      	| An action or state                                            |
| `modifier`  	| General-purpose modifier (fallback)                           |
| `preposition` | A word showing relationship between elements (e.g., "under")|
| `conjunction` | A linking word such as "and", "or"                          |

## Integration

Each tag can optionally reference this table via its `part_of_speech_id`. This allows the system to:

- Validate the **structure of composite tags** (see [`tag_compositions`](../../tags/tag_compositions.md)).  
- Disambiguate search results (e.g., find only nouns or adjectives).  
- Filter or group tags in the UI based on linguistic role.

## Example Usage

| Tag | POS | Example Composite |
|------|------|------------------|
| `"very"` | `adverb` | `"very big car"` |
| `"big"` | `adjective` | `"very big car"` |
| `"car"` | `noun` | `"very big car"` |

Validation: ✅ `adverb` → `adjective` → `noun`  
Invalid: ❌ `"very big"` (missing required `noun` root)

## API Endpoint Summary

| Category | Endpoint | CRUD Coverage | Description |
|-----------|-----------|----------------|--------------|
| **Parts of Speech** | `/parts-of-speech` | **GET**, **POST** | Retrieve or create part-of-speech definitions used for tag classification. |
| | `/parts-of-speech/{id}` | **GET**, **PATCH**, **DELETE** | Retrieve, update, or delete a specific part-of-speech record. |

### CRUD Parity Overview

| Scope | CRUD | Description |
|--------|------|--------------|
| Parts of Speech | ✅ Full | Manage grammatical classifications that influence tag composition and validation. |

✅ = Full REST-compliant CRUD support.

## Summary

Parts of speech provide a **grammatical backbone** to the tagging system, enabling structured and linguistically valid compositions.  
By embedding basic linguistic logic into the data model, the system gains the ability to generate, validate, and interpret complex tags intelligently and consistently.
