#!/usr/bin/env bash
set -euxo pipefail

export DEBIAN_FRONTEND=noninteractive

apt update
apt -y dist-upgrade
apt -y autoremove --purge
apt -y autoclean
apt -y clean
rm -rf /var/lib/apt/lists/*