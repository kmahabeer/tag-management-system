
# Local Development - Database Persistence

This document explains how to persist Postgres data in a local development environment using **Docker Compose** and **Dev Containers**. It also describes how to safely back up, restore, and migrate data stored in a Docker volume.

## 1. Overview

By default, Docker containers are ephemeral—data inside a container is destroyed when it is removed. To preserve Postgres data between environment rebuilds, you must use either:

1. A **named Docker volume** (preferred), or
2. A **bind mount** to a local folder.

The sections below describe both approaches and how to export, back up, restore, and migrate data between environments.

## 2. Option A: Named Volume (Recommended)

In your `docker-compose.yml`:

```yaml
services:
  db:
    image: postgres:16
    restart: unless-stopped
    environment:
      POSTGRES_USER: app
      POSTGRES_PASSWORD: app
      POSTGRES_DB: app
    volumes:
      - postgres-data:/var/lib/postgresql/data
    networks:
      - devnet
    ports:
      - "5432:5432"

volumes:
  postgres-data:
    name: devcontainer_postgres_data
```

### Notes

* This creates a persistent Docker volume named `devcontainer_postgres_data`.
* The data will survive `docker compose down`, container rebuilds, and Dev Container restarts.
* Avoid running `docker compose down -v` as that deletes the volume.

To inspect the volume:

```bash
docker volume inspect devcontainer_postgres_data
```

## 3. Option B: Bind Mount (Alternative)

If you prefer to see the Postgres files directly on your host system, replace the volume with a bind mount:

```yaml
db:
  image: postgres:16
  restart: unless-stopped
  environment:
    POSTGRES_USER: app
    POSTGRES_PASSWORD: app
    POSTGRES_DB: app
  volumes:
    - ./_data/postgres:/var/lib/postgresql/data
  networks:
    - devnet
  ports:
    - "5432:5432"
```

This stores database files in a visible directory (`_data/postgres/`).
Add this folder to `.gitignore` to avoid committing binary data.

Note: Bind mounts can be slower on macOS or Windows.

## 4. Accessing and Exporting Data

You will not see files from a named volume on your host, but you can easily extract or back up data using standard Postgres tools or Docker commands.

### 4.1. Database Dump

Create a `.devcontainer/scripts/backup_db.sh` file:

```bash
#!/bin/bash
set -e

timestamp=$(date +"%Y%m%d_%H%M%S")
mkdir -p _backups

docker exec -t db pg_dump -U app -d app > "_backups/backup_${timestamp}.sql"
echo "Backup saved to _backups/backup_${timestamp}.sql"
```

To restore:

```bash
#!/bin/bash
set -e

backup_file="$1"
if [ -z "$backup_file" ]; then
  echo "Usage: $0 <path_to_backup.sql>"
  exit 1
fi

docker exec -i db psql -U app -d app < "$backup_file"
```

Make both scripts executable:

```bash
chmod +x .devcontainer/scripts/*.sh
```

Mount a `_backups` folder in your `devcontainer.json` so that backups persist:

```json
"mounts": [
  "source=${localWorkspaceFolder}/_backups,target=/workspaces/${localWorkspaceFolderBasename}/_backups,type=bind"
]
```

### 4.2. Full Volume Snapshot

To create a tarball of the entire volume:

```bash
docker run --rm -v devcontainer_postgres_data:/data -v $(pwd):/backup alpine \
  tar czf /backup/postgres_data.tar.gz -C /data .
```

To restore:

```bash
docker run --rm -v devcontainer_postgres_data:/data -v $(pwd):/backup alpine \
  sh -c "cd /data && tar xzf /backup/postgres_data.tar.gz"
```

## 5. Database Migration

Migrations allow you to modify your database schema over time without losing data. They are essential for versioning schema changes across environments (development, staging, production).

### 5.1. Using Alembic (Python)

If your stack includes FastAPI or SQLAlchemy, install Alembic:

```bash
pip install alembic
```

Initialize a migration environment:

```bash
alembic init migrations
```

Edit `alembic.ini` to set the database URL:

```bash
sqlalchemy.url = postgresql://app:app@db:5432/app
```

Generate a migration after modifying models:

```bash
alembic revision --autogenerate -m "Add new table"
```

Apply migrations:

```bash
alembic upgrade head
```

Downgrade (optional):

```bash
alembic downgrade -1
```

### 5.2. Using SQL Migration Files

If you prefer raw SQL migrations, create a folder structure like:

```txt
migrations/
  ├── 001_init.sql
  ├── 002_add_user_table.sql
  └── 003_add_index.sql
```

Apply them manually or via a script:

```bash
for f in migrations/*.sql; do
  docker exec -i db psql -U app -d app < "$f"
done
```

### 5.3. Migrating Between Environments

To migrate your local database to another environment (e.g., staging):

1. Create a SQL dump from the source:

   ```bash
   docker exec -t db pg_dump -U app -d app > staging_import.sql
   ```

2. Transfer the file to the destination host.
3. Restore it to the target database:

   ```bash
   docker exec -i staging_db psql -U app -d app < staging_import.sql
   ```

This ensures the schema and data are consistent across environments.

## 6. Common Commands

| Action                        | Command                                            |
| -- | -- |
| Start containers              | `docker compose up -d`                             |
| Stop containers (keep data)   | `docker compose down`                              |
| Stop containers (delete data) | `docker compose down -v`                           |
| Inspect Postgres volume       | `docker volume inspect devcontainer_postgres_data` |
| List all Docker volumes       | `docker volume ls`                                 |
| Delete the Postgres volume    | `docker volume rm devcontainer_postgres_data`      |
| Generate Alembic migration    | `alembic revision --autogenerate -m "message"`     |
| Apply migrations              | `alembic upgrade head`                             |

## 7. Summary

* Use a **named Docker volume** for persistent Postgres data across container rebuilds.
* Avoid deleting the volume unless you intend to reset the database.
* Use backup scripts to export data to `_backups/`.
* Use **migrations** (Alembic or SQL scripts) to evolve schema safely.
* Optionally use a bind mount for direct visibility on the host.

This configuration ensures reliable, persistent, and version-controlled Postgres data for local development.
