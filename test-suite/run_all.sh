#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd -- "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$ROOT_DIR"

if command -v uv >/dev/null 2>&1; then
  uv run python ./run_manifest.py "$@"
  uv run python shareable/test_shareable_reports.py
  uv run python python/tests/test_resource_helper.py
  uv run python java/tests/test_resource_lifecycle_helper.py
else
  echo "[warn] uv not found – falling back to system python3. Run 'uv sync --python 3.13' for the supported toolchain." >&2
  python3 ./run_manifest.py "$@"
  python3 shareable/test_shareable_reports.py
  python3 python/tests/test_resource_helper.py
  python3 java/tests/test_resource_lifecycle_helper.py
fi
