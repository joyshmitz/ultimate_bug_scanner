#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd -- "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd -- "$ROOT_DIR/.." && pwd)"
UBS_BIN="${UBS_BIN:-$REPO_ROOT/ubs}"

if [[ ! -x "$UBS_BIN" ]]; then
  echo "Unable to find executable UBS binary at $UBS_BIN" >&2
  exit 127
fi

cases=(
  "js-buggy:$ROOT_DIR/buggy:fail"
  "js-clean:$ROOT_DIR/clean:pass"
  "python-buggy:$ROOT_DIR/python/buggy:fail"
  "python-clean:$ROOT_DIR/python/clean:pass"
  "cpp-buggy:$ROOT_DIR/cpp/buggy:fail"
  "cpp-clean:$ROOT_DIR/cpp/clean:pass"
  "rust-buggy:$ROOT_DIR/rust/buggy:fail"
  "rust-clean:$ROOT_DIR/rust/clean:pass"
  "golang-buggy:$ROOT_DIR/golang/buggy:fail"
  "golang-clean:$ROOT_DIR/golang/clean:pass"
  "java-buggy:$ROOT_DIR/java/buggy:fail"
  "java-clean:$ROOT_DIR/java/clean:pass"
  "ruby-buggy:$ROOT_DIR/ruby/buggy:fail"
  "ruby-clean:$ROOT_DIR/ruby/clean:pass"
)

failures=0

run_case() {
  local name="$1" path="$2" expectation="$3" extra="$4"
  echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
  echo "Running ${name} (${path})"
  echo "Command: $UBS_BIN --ci ${extra} ${path}"
  echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"

  local cmd=("$UBS_BIN" --ci)
  if [[ -n "$extra" ]]; then
    # shellcheck disable=SC2206 # we want word splitting for arguments
    cmd+=($extra)
  fi
  cmd+=("$path")

  if "${cmd[@]}"; then
    status=0
  else
    status=$?
  fi

  case "$expectation" in
    pass)
      if [[ $status -eq 0 ]]; then
        echo "✅ ${name} passed"
      else
        echo "❌ ${name} expected success but exited $status"
        failures=$((failures + 1))
      fi
      ;;
    fail)
      if [[ $status -ne 0 ]]; then
        echo "✅ ${name} correctly failed (exit $status)"
      else
        echo "❌ ${name} expected non-zero exit"
        failures=$((failures + 1))
      fi
      ;;
    *)
      echo "Unknown expectation '$expectation'" >&2
      failures=$((failures + 1))
      ;;
  esac
  echo
}

for case_entry in "${cases[@]}"; do
  IFS=":" read -r name path expectation extra <<<"$case_entry"
  run_case "$name" "$path" "$expectation" "$extra"
done

if [[ $failures -ne 0 ]]; then
  echo "Test suite finished with $failures failure(s)."
  exit 1
fi

echo "All suites completed successfully."
