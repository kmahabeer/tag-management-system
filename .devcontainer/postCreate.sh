#!/usr/bin/env bash
set -euxo pipefail

export DEBIAN_FRONTEND=noninteractive

apt update
apt -y dist-upgrade
apt -y autoremove --purge
apt -y autoclean
apt -y clean
rm -rf /var/lib/apt/lists/*

# Ensure shared dev network exists
docker network inspect tag-management-system_dev_net >/dev/null 2>&1 || \
  docker network create tag-management-system_dev_net

# Start the backend API
uv run uvicorn app.main:app --reload --port 8100 # Needs a way to wait until uv installs all packages

# Start the frontend UI
cd frontend
npm run dev