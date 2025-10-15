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

# # Start backend API
# uvicorn app.main:app --reload # Needs a way to wait until uv installs all packages