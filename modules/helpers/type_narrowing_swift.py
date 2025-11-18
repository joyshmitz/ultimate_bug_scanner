#!/usr/bin/env python3
"""Detect Swift guard let / optional bindings that continue without exiting."""
from __future__ import annotations

import re
import sys
from pathlib import Path

SKIP_DIRS = {".git", ".hg", ".svn", "build", "DerivedData", ".swiftpm", ".idea", "node_modules"}
GUARD_PATTERN = re.compile(r"guard\s+let\s+([A-Za-z_][\w]*)\s*=\s*[^\\n]+\s+else\s*\{", re.MULTILINE)
EXIT_KEYWORDS = ("return", "throw", "break", "continue", "fatalError", "preconditionFailure")


def iter_swift_files(root: Path):
    if root.is_file():
        if root.suffix == ".swift" and not any(part in SKIP_DIRS for part in root.parts):
            yield root
        return
    for path in root.rglob("*.swift"):
        if any(part in SKIP_DIRS for part in path.parts):
            continue
        if path.is_file():
            yield path


def find_block_end(text: str, brace_start: int) -> int:
    depth = 0
    for idx in range(brace_start, len(text)):
        ch = text[idx]
        if ch == "{":
            depth += 1
        elif ch == "}":
            depth -= 1
            if depth == 0:
                return idx
    return len(text) - 1


def block_has_exit(block: str) -> bool:
    lower = block.lower()
    for keyword in EXIT_KEYWORDS:
        if keyword.lower() in lower:
            return True
    return False


def line_col(text: str, pos: int) -> tuple[int, int]:
    line = text.count("\n", 0, pos) + 1
    last_newline = text.rfind("\n", 0, pos)
    if last_newline == -1:
        col = pos + 1
    else:
        col = pos - last_newline
    return line, col


def analyze_file(path: Path):
    text = path.read_text(encoding="utf-8", errors="ignore")
    issues = []
    for match in GUARD_PATTERN.finditer(text):
        name = match.group(1)
        block_start = match.end()
        block_end = find_block_end(text, block_start)
        block_text = text[block_start:block_end]
        if block_has_exit(block_text):
            continue
        line, col = line_col(text, block_start)
        message = f"guard let '{name}' else-block does not exit before continuing"
        issues.append((line, col, message))
    return issues


def main() -> int:
    if len(sys.argv) < 2:
        print("Usage: type_narrowing_swift.py <project_dir>", file=sys.stderr)
        return 1
    root = Path(sys.argv[1]).resolve()
    if not root.exists():
        return 0
    any_output = False
    for path in iter_swift_files(root):
        try:
            issues = analyze_file(path)
        except OSError:
            continue
        for line, col, message in issues:
            any_output = True
            print(f"{path}:{line}:{col}\t{message}")
    return 0 if any_output or root.is_dir() else 0


if __name__ == "__main__":
    raise SystemExit(main())
