#!/usr/bin/env bash
set -e
# Wrapper to run the python update script using the project venv if available
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(dirname "$SCRIPT_DIR")"

# Ensure we are in the root
cd "$ROOT_DIR"

if command -v uv >/dev/null 2>&1; then
    uv run python scripts/update_checksums.py
elif [[ -f ".venv/bin/python3" ]]; then
    .venv/bin/python3 scripts/update_checksums.py
else
    python3 scripts/update_checksums.py
fi

