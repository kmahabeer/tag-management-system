# Tagging Service

This system is designed to power a flexible and context-aware Inspector UI for entities like images, film stills, documents, and more. Each asset may serve different "purposes" (e.g. a film still, a figure study, a reference photo) and should dynamically display its associated metadata (tags) in a structured way according to those purposes.

At the heart of the system is a flexible tagging model with many-to-many relationships, hierarchical categorization, and purpose-specific rendering logic. Tags can represent people, roles, vehicles, genres, purposes, and any other useful classification. These tags can be reused and linked across entities like films or subjects.

Frontend clients, such as a web-based graphical user-interface (UI) will read structured JSON formatted according to predefined layouts for each purpose (e.g., "Film Still" might expect to show Directors, Writers, Actors, etc.). Tags that do not fit into this predefined structure are shown in a catch-all "Other Tags" section.

## Features

- Tag entities with hierarchical, typed, and composable labels
- Group tags by UI logic using types and display controls
- Store purposes for each entity to customize rendering
- Link tags to multiple entities with support for aliases and relationships
- Supports modifier phrases like “very tall” or “bright red”

## Usage

The system is designed to expose a RESTful API for tag creation, updates, querying, and relationship management.

See the [documentation](./docs/_index.md) for detailed schema, logic, and integration guides.
