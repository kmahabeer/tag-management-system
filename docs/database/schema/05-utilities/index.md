---
title: Utilities
parent: Database Schema
has_children: true
nav_order: 5
---
# Utilities

The **utilities** schema defines shared reference tables that provide semantic structure and consistency across tags, entities, and relationships.

## API Endpoint Summary

| Category | Endpoint | CRUD Coverage | Description |
|-----------|-----------|----------------|--------------|
| **Contexts** | `/contexts`, `/contexts/{id}` | **GET**, **POST**, **PATCH**, **DELETE** | Manage semantic contexts used to group or evaluate tags and entities (e.g., “style,” “subject,” “emotion”). |
| **Parts of Speech** | `/parts-of-speech`, `/parts-of-speech/{id}` | **GET**, **POST**, **PATCH**, **DELETE** | Manage grammatical classifications for tags (e.g., noun, adjective, adverb). |
| **Ratings** | `/ratings`, `/ratings/{id}` | **GET**, **POST**, **PATCH**, **DELETE** | Define rating values representing qualitative or quantitative scores. |
| **Rating Types** | `/rating-types`, `/rating-types/{id}` | **GET**, **POST**, **PATCH**, **DELETE** | Define categories of ratings (e.g., “Clarity,” “Confidence,” “Likeness”). |

### Coverage Summary

| Utility | CRUD | Description |
|----------|------|--------------|
| Contexts | ✅ Full | Defines evaluative or semantic dimensions for classification and rating. |
| Parts of Speech | ✅ Full | Defines grammatical roles used to validate tag compositions. |
| Ratings | ✅ Full | Defines reusable rating values used by tags, entities, and relationships. |
| Rating Types | ✅ Full | Defines rating categories and normalization rules. |
