# API Reference

The **Tagging Service API** exposes a RESTful interface for managing all tag-related operations: creation, updates, deletion, linking, and querying of tags and their relationships. It also supports metadata assignment for tagged entities such as images, documents, or video stills.

This document provides an overview of the structure, conventions, and usage patterns for the API.

- [Endpoints](./endpoints/_index.md)
	- [Entities](./endpoints/endpoint-entities.md)
	- [Tags](./endpoints/endpoint-tags.md)
	- [Utilities](./endpoints/endpoint-utilities.md)

## Principles

- **RESTful architecture**: Resource-based URIs and standard HTTP verbs (GET, POST, PUT, DELETE).
- **JSON payloads**: All requests and responses are structured in JSON for consistency and ease of use.
- **Frontend-controlled authentication**: Only frontend clients (CLI, UI) interact with the Tagging Service API through validated tokens; the only exception is for development and testing purposes.

## URL Structure

### Base URL

All API endpoints are accessible under the following base path (example for local development):

```bash
http://localhost:8000/api
```

API version number is appended to the base URL in the path (e.g., `/v1`, `/v2`, etc.) to support compatibility.

## RESTful Conventions

The API follows REST conventions:

- **GET**: Retrieve resources
- **POST**: Create new resources
- **PATCH**: Partially update existing resources
- **DELETE**: Remove resources

Endpoint paths are structured around **noun-based resources**, e.g., `/tags` and `/entities`.
