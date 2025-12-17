#!/usr/bin/env bash
set -euo pipefail

# Update SHA256SUMS file with current ubs script checksum
# Run this after modifying the ubs meta-runner script

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(dirname "$SCRIPT_DIR")"

cd "$ROOT_DIR"

if [[ ! -f "ubs" ]]; then
    echo "Error: ubs script not found" >&2
    exit 1
fi

CHECKSUM=$(sha256sum ubs | awk '{print $1}')
echo "${CHECKSUM}  ubs" > SHA256SUMS

echo "Updated SHA256SUMS:"
cat SHA256SUMS
