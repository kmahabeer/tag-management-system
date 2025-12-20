# Tag Management System

The **Tag Management System** is designed to power a flexible and context-aware user interface (UI) for entities like images, film stills, documents, and more. Each entity may serve different "purposes" (e.g. a film still, a figure study, a reference photo) and should dynamically display its associated metadata (tags) in a structured way according to those purposes.

At the heart of the system is a flexible tagging model with many-to-many relationships, hierarchical categorization, and purpose-specific rendering logic. Tags can represent people, roles, vehicles, genres, purposes, and any other useful classification. These tags can be reused and linked across entities like films or subjects.

Frontend clients, such as a web-based graphical user interface (GUI) will read structured JSON according to predefined layouts for each purpose (e.g., "Film Still" might expect to show Directors, Writers, Actors, etc.). Tags that do not fit into this predefined structure are shown in a catch-all "Other Tags" section.

## Features

- Tag entities with hierarchical, typed, and composable labels
- Group tags by UI logic using types and display controls
- Store purposes for each entity to customize rendering
- Link tags to multiple entities with support for aliases and relationships
- Supports modifier phrases like “very tall” or “bright red”

## Getting Started

### Prerequisites

- Python 3.12+
- Node.js 18+
- PostgreSQL (for production) or Docker (for local dev)

### Backend Setup

1. Install Python dependencies using `uv`:

   ```bash
   uv sync
   ```

2. Set up the database:
   - For local development, ensure PostgreSQL is running.
   - The app will automatically create tables on startup.

3. Run the backend:

   ```bash
   uv run uvicorn app.main:app --reload --host 0.0.0.0 --port 8100
   ```

   The API will be available at `http://localhost:8100/api/v1`.

### Frontend Setup

1. Install dependencies:

   ```bash
   cd frontend
   npm install
   ```

2. Run the development server:

   ```bash
   npm run dev
   ```

   The frontend will be available at `http://localhost:5175`.

### Development

- Backend API docs: `http://localhost:8100/docs`
- Frontend connects to backend at `http://localhost:8100/api/v1`
- Use `make routes` to regenerate API routes from OpenAPI spec
- Use `make schemas` to regenerate Pydantic models from OpenAPI

## Usage

The system exposes a RESTful API for managing tags, relationships, and entity tagging.

See the [documentation](./docs/index.md) for detailed schema, logic, and integration guides.
