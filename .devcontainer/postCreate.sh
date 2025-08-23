#!/usr/bin/env bash
set -euxo pipefail

# Start backend API
uvicorn app.main:app --reload # Needs a way to wait until uv installs all packages