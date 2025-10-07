#!/usr/bin/env bash
set -euxo pipefail

# Synchronize backend packages with Python virtual environment
cd app
uv sync
cd ..