#!/usr/bin/env python3
"""
Update checksums for UBS modules in the meta-runner script.
"""
import hashlib
import re
import sys
from pathlib import Path

def compute_sha256(path: Path) -> str:
    if not path.exists():
        return ""
    return hashlib.sha256(path.read_bytes()).hexdigest()

def main():
    root = Path(__file__).resolve().parent.parent
    ubs_script = root / "ubs"
    modules_dir = root / "modules"

    if not ubs_script.exists():
        print(f"Error: ubs script not found at {ubs_script}", file=sys.stderr)
        sys.exit(1)

    if not modules_dir.exists():
        print(f"Error: modules directory not found at {modules_dir}", file=sys.stderr)
        sys.exit(1)

    print("Updating module checksums in ubs...")
    
    # Map lang to filename
    # bash associative array keys in ubs: js, python, cpp, rust, golang, java, ruby, swift
    # filenames: ubs-js.sh, ubs-python.sh, etc.
    
    lang_map = {
        "js": "ubs-js.sh",
        "python": "ubs-python.sh",
        "cpp": "ubs-cpp.sh",
        "rust": "ubs-rust.sh",
        "golang": "ubs-golang.sh",
        "java": "ubs-java.sh",
        "ruby": "ubs-ruby.sh",
        "swift": "ubs-swift.sh"
    }

    new_checksums = {}
    
    for lang, filename in lang_map.items():
        path = modules_dir / filename
        if not path.exists():
            print(f"Warning: Module {filename} not found for {lang}")
            continue
            
        checksum = compute_sha256(path)
        print(f"  {lang}: {checksum}")
        new_checksums[lang] = checksum

    # Read ubs script
    content = ubs_script.read_text(encoding="utf-8")
    
    # Regex to find the MODULE_CHECKSUMS array block
    # It looks like:
    # declare -A MODULE_CHECKSUMS=(
    #   [js]='...'
    #   ...
    # )
    
    pattern = re.compile(r"(declare -A MODULE_CHECKSUMS=\s*\()([\s\S]*?)(\))", re.MULTILINE)
    
    def replace_checksums(match):
        prefix = match.group(1)
        suffix = match.group(3)
        
        lines = []
        for lang in sorted(lang_map.keys()): # Sort for stability
            if lang in new_checksums:
                # Preserve indentation
                lines.append(f"  [{lang}]='{new_checksums[lang]}'")
        
        return f"{prefix}\n" + "\n".join(lines) + f"\n{suffix}"

    new_content = pattern.sub(replace_checksums, content)
    
    if new_content != content:
        ubs_script.write_text(new_content, encoding="utf-8")
        print("✓ ubs script updated with new checksums.")
    else:
        print("✓ No changes needed.")

if __name__ == "__main__":
    main()

