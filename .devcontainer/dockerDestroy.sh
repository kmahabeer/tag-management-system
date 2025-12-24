#!/usr/bin/env bash
set -euxo pipefail

# stop all containers
docker stop $(docker ps -aq) || true

# remove all containers
docker rm $(docker ps -aq) || true

# prune everything without prompts
docker container prune -f
docker image prune -af
docker volume prune -f
docker network prune -f
docker system prune -af