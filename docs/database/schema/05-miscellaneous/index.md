---
title: Miscellaneous
parent: Database Schema
has_children: true
nav_order: 5
---

## API Endpoint Summary

| Category            | Endpoint                | CRUD Coverage                  | Description                                                                                          |
| ------------------- | ----------------------- | ------------------------------ | ---------------------------------------------------------------------------------------------------- |
| **Contexts**        | `/contexts`             | **GET**, **POST**              | Retrieve or create semantic contexts used to group, rate, or interpret tags and entities.            |
|                     | `/contexts/{id}`        | **GET**, **PATCH**, **DELETE** | Retrieve, update, or delete a specific context definition.                                           |
| **Parts of Speech** | `/parts-of-speech`      | **GET**, **POST**              | Retrieve or create grammatical classifications used for tag composition (e.g., noun, adjective).     |
|                     | `/parts-of-speech/{id}` | **GET**, **PATCH**, **DELETE** | Retrieve, update, or delete a part of speech record.                                                 |
| **Ratings**         | `/ratings`              | **GET**, **POST**              | Retrieve or create system-wide rating values that represent qualitative or quantitative evaluations. |
|                     | `/ratings/{id}`         | **GET**, **PATCH**, **DELETE** | Retrieve, update, or delete a specific rating value.                                                 |
| **Rating Types**    | `/rating-types`         | **GET**, **POST**              | Retrieve or create rating categories that define semantic meaning (e.g., “Clarity,” “Confidence”).   |
|                     | `/rating-types/{id}`    | **GET**, **PATCH**, **DELETE** | Retrieve, update, or delete a specific rating type definition.                                       |

## Coverage Summary

| Utility         | CRUD   | Description                                                                   |
| --------------- | ------ | ----------------------------------------------------------------------------- |
| Contexts        | ✅ Full | Defines evaluative or semantic dimensions like *style*, *mood*, or *content*. |
| Parts of Speech | ✅ Full | Defines grammatical roles for tag composition logic.                          |
| Ratings         | ✅ Full | Defines reusable rating values used by tags, entities, and relationships.     |
| Rating Types    | ✅ Full | Defines rating categories and normalization rules.                            |
