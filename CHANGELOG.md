# Changelog

All notable changes to Ultimate Bug Scanner (UBS) are documented in this file.

Versions marked **[Release]** have a corresponding [GitHub Release](https://github.com/Dicklesworthstone/ultimate_bug_scanner/releases) with artifacts (installer, checksums, SBOM, Homebrew formula). Versions marked **[Tag]** are git tags without published release assets.

Repository: <https://github.com/Dicklesworthstone/ultimate_bug_scanner>

---

## [Unreleased] (after v5.0.6)

> Commits on `master` since 2026-01-05, not yet tagged.

### New Language: C# (9th language)

- Add complete C# language module (`ubs-csharp.sh`) with test suite, bringing UBS to 9 supported languages ([`7670efc`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/7670efc40040aba5d5c1ab9c72b969f257c048da))
- Integrate async task-handle analysis, structured ast-grep ingestion, and Ruby/Swift resource lifecycle helpers for C# ([`7e755e8`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/7e755e8))
- Add C++ AST resource lifecycle helper; harden C#/Java modules; pass-through dotnet CLI flags ([`0f8f182`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/0f8f1822c71a2f29f5665ca0f642aad051fa6a72))

### New Features

- Add multi-file scanning with `--files` flag ([`9ce7471`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/9ce7471b51e24c632e9b5ff4cc0f0ad938ed3bde))
- Add TOON format output (`--format=toon`) for ~50% smaller token cost ([`5e9cee8`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/5e9cee81ab5241473fd0cf37fbed4668fe303f45))
- Add `--skip-size-check` flag and `load_ignore_patterns` for .ubsignore-aware size calculation ([`c27235e`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/c27235e219e40ea4acd99327fcd3d13b6cd1e983))
- Add Claude Code SKILL.md for automatic capability discovery ([`a59f7bb`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/a59f7bba4d66cc2759f074f89a35c66bb37cf470))
- Add CI workflow for build, test, and lint ([`e58e2c2`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/e58e2c2958dc28d6c19f0b477eaa5955a6b8f0a2))

### Bug Fixes

- Eliminate false positives in JS credential detection and Python mutable default checks ([`c54c27c`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/c54c27c8c4383ba6a9d2768172e6194c003a5aaa))
- Fall through to Python size-check when `.ubsignore` has path patterns ([`f90f948`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/f90f948bba2e13c005e38614fa8ca5dca9df60b5))
- Bypass whole-repo size guards for targeted scan modes (`--files`, `--staged`) ([`4c3bd50`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/4c3bd5071bb5d8db9050862014e98ad7ec34748c))
- Prevent Cellar path error in Homebrew installs (closes #29) ([`f00743a`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/f00743ad604bf82e1caeb9ccd54af54dc9bdcc00))
- Implement inline `ubs:ignore` suppression and apply `.ubsignore` in `--staged` mode (#24) ([`4b8a111`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/4b8a111261d3c5e6b47684759c916065ed78a401))
- Bash 4.0 version gate, `script_dir` symlink resolution, checksum refresh (#25, #26, #27) ([`c9a8316`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/c9a831650c95b95c5a6bbb3cc70d48973f635de5))
- Expand `.ubsignore` glob patterns for Bandit exclusions (#22) ([`9c81a5c`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/9c81a5c))
- Respect `.ubsignore` patterns when calculating directory size ([`d29ec78`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/d29ec7807f39d43d460a7d314a40d2fd110fbf9d))
- Add safety guards to prevent disk exhaustion (#12) ([`b9f3982`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/b9f39825f99b1eedd1d2be75a3b75ed9902c0f2f))
- Skip `du` on macOS/BSD when exclude patterns are needed ([`8cbb069`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/8cbb069d4fc8ef96dd8b6adf87b41fc45dbe4eec))
- Escape backtick `` `as` `` in Rust module to prevent command substitution ([`b742cbc`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/b742cbc0ed5daf485523da143431d662362c52a9))
- Exclude `.venv` from language detection in `detect_lang()` ([`27a0d96`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/27a0d960f15816ff74ad03acae74a624498f080a))
- Remove nonexistent `ubs-c.sh` module reference from `ALL_LANGS` ([`11dd82a`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/11dd82a))
- Correct module download URL to use master branch ([`97bb3c0`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/97bb3c0))
- Improve shell script robustness and detection patterns ([`3d9dc12`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/3d9dc12))

### Other

- Update license to MIT with OpenAI/Anthropic Rider ([`534c897`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/534c8977a2d13daa9206c408a46c6c8bb4f19971))
- Add MIT license file ([`3d6d7f3`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/3d6d7f341636220af6c56763371b0478f1c410a2))

---

## [v5.0.6] - 2026-01-05 **[Release]**

> Title: "Documentation Deep-Dive"

### Documentation

- Add 201 lines of technical deep-dive documentation covering AST rule architecture, inline suppression, cross-language async detection, helper script verification, and unified severity normalization ([`2b632f0`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/2b632f0aadb6e7aa48bdb97495f959606868ea58))

### Bug Fixes (JS Scanner)

- Fix `then-without-catch` rules and `Promise.all` consistency ([`5fc2095`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/5fc20955bfd88257ff38a50a2b43ef166f4eaff0))
- Improve `.catch()` chain detection in ast-grep rules ([`f89e0b5`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/f89e0b5))
- Add `stopBy: end` to ast-grep rules for proper ancestor traversal ([`688ea6d`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/688ea6d0e2fc407856440636ce8af7583b8c8379))

### Bug Fixes (Other)

- Repair JS `typeof` detection and Go false positive ([`2377537`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/2377537a1b78ec7feddf6ed5b131f3118a222ca8))
- Correctly parse UBS summary JSON from JSONL output in test suite ([`7868512`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/786851256046bb643dbcf93640a521df9c000d56))

**Full diff:** [`v5.0.5...v5.0.6`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/compare/v5.0.5...v5.0.6)

---

## [v5.0.5] - 2026-01-05 **[Release]**

> Title: "Documentation Release"

Documentation-only release. No code changes; all features documented were already present in v5.0.4.

### Documentation

- Comprehensive documentation for 8 previously undocumented feature systems (~386 lines added) ([`ff5b90f`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/ff5b90f91f575673a6680d8d385f9ce9bfaa739d)):
  - Safety guards for AI coding agents (`git_safety_guard.py`)
  - Extended agent detection (12+ coding agents)
  - AST-based type narrowing analysis
  - ast-grep auto-provisioning with SHA-256 verification
  - Maintenance commands (`ubs doctor`, `ubs sessions`)
  - Test suite infrastructure (manifest-driven `run_manifest.py`)
  - Auto-update system (`--update`, `--no-auto-update`)
  - Beads/Strung JSONL integration (`--beads-jsonl`)

### Bug Fixes

- Multiple helper file bugs found via code exploration ([`61fdbde`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/61fdbdea9ab3ff4225b00319b36a50b7f9189b36))
- Use ast-grep for var declaration and division finding display ([`dc55bee`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/dc55bee01843e8b40fbbce3455ed1da3226e64e8), [`aa834ee`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/aa834ee0b39de0bb78cf7129c92e64e27ed7f52d))

**Full diff:** [`v5.0.4...v5.0.5`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/compare/v5.0.4...v5.0.5)

---

## [v5.0.4] - 2026-01-05 **[Release]**

> Security patch release.

### Security

- **Critical:** Close `/bin/rm -rf` bypass in `git_safety_guard.py`. Before this fix, absolute paths to `rm` (e.g., `/bin/rm -rf /important`) were not blocked. Added `_is_rm_command()` helper that recognizes both `rm` and any path ending in `/rm` ([`8907eec`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/8907eec8c0ca08e0c04bea467f07bb5d9ab7e65d))

### Bug Fixes (CI)

- Fix output variables bug in checksum-health workflow ([`f43f420`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/f43f420))
- Address bugs found in code review ([`c2e954a`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/c2e954a))

### New

- Add defense-in-depth checksum monitoring CI workflow ([`55bdf23`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/55bdf23cb68f593bd5e12345b788f18f67b324ec))

**Full diff:** [`v5.0.3...v5.0.4`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/compare/v5.0.3...v5.0.4)

---

## [v5.0.3] - 2025-12-30 **[Release]**

### Bug Fixes

- Handle uppercase `-R`/`-F` flags in `rm_rf_targets_are_safe` ([`de62531`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/de625315421a03e363028411fdca78d0e9dd6230))
- Catch `rm` bypass variants in `git_safety_guard.py` ([`ec72ce7`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/ec72ce7de7a20275bcabdfa4c7df437cadd56387))
- Include version tag in Docker image tags ([`162ed53`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/162ed537bc9e3fad534dc9a7846ff7db3af2dd8d))

**Full diff:** [`v5.0.2...v5.0.3`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/compare/v5.0.2...v5.0.3)

---

## [v5.0.2] - 2025-12-30 **[Release]**

### Bug Fixes

- Skip SBOM generation for PR builds in OCI workflow ([`a468726`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/a468726))

Version bump for ARM64 Docker release.

**Full diff:** [`v5.0.1...v5.0.2`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/compare/v5.0.1...v5.0.2)

---

## [v5.0.1] - 2025-12-30 **[Release]**

### New

- Add multi-platform Docker builds (amd64 + arm64) ([`420370a`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/420370adf0fc40043aa5acbaa0f59d0716f19c98))

### Bug Fixes

- Make `git_safety_guard` `rm -rf` allowlist non-bypassable ([`096e1ee`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/096e1ee24a9e10eabb42e2615339d4e9f3ea6f05))
- Harden supply-chain verification and machine output ([`f2e320a`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/f2e320ad18962cecaa995a6f6ab9f4058eb24b48))
- Remove tautological return expressions in type narrowing helpers ([`e401c5b`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/e401c5b))
- Correct ripgrep TypeScript file detection in JS module ([`355a78d`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/355a78d))
- Address false positives in Go heuristics; add Bun runtime support ([`e7a6cd8`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/e7a6cd8))
- Reduce installer noise and improve Node.js detection ([`e089b32`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/e089b32))
- Allow `rm -rf` on temp directories in hooks ([`3a369ed`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/3a369ed))

### Other

- Add Codex CLI v0.77.0+ migration guidance and directory-format rules ([`cc9c6a9`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/cc9c6a9))
- Add contribution policy section to README ([`efa70e2`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/efa70e2))

**Full diff:** [`v5.0.0...v5.0.1`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/compare/v5.0.0...v5.0.1)

---

## [v5.0.0] - 2025-12-19 **[Release]**

> 105 commits, 10,600+ line additions across 50 files since v4.6.5. Major themes: Windows compatibility, mandatory AST-powered JS/TS analysis, all language modules upgraded to v3.0+ standards.

### Breaking Changes

1. **ast-grep is now required** for JS/TS scanning. Install via `brew install ast-grep`, `cargo install ast-grep`, or `npm i -g @ast-grep/cli`.
2. **Exit code 2** now indicates environment errors (missing required tools).
3. **Windows paths** now parsed correctly. Downstream tooling that depended on broken parsing may need adjustment.

### Windows Compatibility

- Add `.ubsignore` fallback chain when rsync unavailable: rsync -> tar -> Python ([`db01365`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/db013652d59ae2b3b3f490e446453609307a543d))
- Fix Windows path parsing (`C:/path:line:code`) across all 8 language modules ([`50b23db`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/50b23dba869d40eda4d7cf06a9fc91172a131402))
- Portable SHA-256 verification (sha256sum/shasum/openssl) ([`be2dc39`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/be2dc399baf342f21535516463cb268fef83070e))

### AST Accuracy Enforcement

- JS/TS scanner now requires ast-grep; noisy grep-based patterns removed ([`a17b0f7`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/a17b0f74d845ac252d8804025808dd4413a41882))
- Type narrowing properly gated to TypeScript inputs only ([`fed9b86`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/fed9b86))
- AST-based division/modulo detection with regex fallback ([`9583f87`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/9583f87691585f584bd32d2127e0231237046ff8))
- AST-based bare `var` declaration detection ([`73202c9`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/73202c902eac72535fa03c63ff5ca914919cab2d))
- AST `await-without-try` rule for async false-positive reduction ([`24a67a0`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/24a67a0debd7006a1341f4e867883612cc14a7a5))
- Stricter async via AST: dangling-promise uses AST context ([`1f3cbd2`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/1f3cbd2fcd87a675da429610cb04b794202f07fd))

### Module Upgrades to v3.0+ Standards

All 8 language modules upgraded ([`0bc391f`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/0bc391f0a5b1aaebb061f7b86c391671bfc009fe)):

| Module | Version | Highlights |
|--------|---------|------------|
| ubs-js.sh | v3.0 | React hook dependency analysis, AST-powered |
| ubs-python.sh | v3.0 | Scope-aware resource lifecycle |
| ubs-golang.sh | v7.1 | Return/defer handling, context leak detection |
| ubs-cpp.sh | v7.1 | Modern C++20 patterns, RAII checks |
| ubs-rust.sh | v3.0 | Type narrowing, cargo integration |
| ubs-java.sh | v3.0 | JDBC lifecycle, CompletableFuture checks |
| ubs-ruby.sh | v3.0 | Block/ensure analysis |
| ubs-swift.sh | v1.8.0 | Guard-let analysis, AST rules |

### Security & Supply Chain

- Add Claude Code hook (`git_safety_guard.py`) to block destructive git commands ([`38d1038`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/38d1038bedfbb52d317a2346d20d6f8594b16636))
- SHA-256 checksums auto-synced via GitHub Action ([`9358e4f`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/9358e4f))
- Multi-layered checksum verification system ([`4e61674`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/4e616749210575adb072e94f6e1f087200b43053))
- Inline suppression: `ubs:ignore`, `nolint`, `noqa` markers respected ([`44c9899`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/44c9899))

### Infrastructure

- JSONL output mode and `--version` flag ([`ed4b233`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/ed4b233ea5f763216b83c5524a4b2ae3d75cd0f5))
- Resource lifecycle analyzers now scope-aware (handles return/yield/defer) ([`3d5ec19`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/3d5ec19), [`9133d97`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/9133d97))
- AI agent quality guardrails for Cursor, Codex, Gemini, Windsurf, Cline ([`fbb479a`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/fbb479a))

**Full diff:** [`v4.6.5...v5.0.0`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/compare/v4.6.5...v5.0.0)

---

## [v4.6.5] - 2025-11-22 **[Release]**

CI release-pipeline fix. No scanner changes.

- Avoid duplicate asset names in release publish ([`952df6f`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/952df6f))

**Full diff:** [`v4.6.4...v4.6.5`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/compare/v4.6.4...v4.6.5)

---

## [v4.6.4] - 2025-11-22 **[Release]**

CI release-pipeline fix. No scanner changes.

- Disable uv cache to avoid missing cache path error ([`9135bb7`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/9135bb7))

**Full diff:** [`v4.6.3...v4.6.4`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/compare/v4.6.3...v4.6.4)

---

## [v4.6.3] - 2025-11-22 **[Tag]**

CI release-pipeline fix. No scanner changes.

- Fix setup-uv input name ([`c04abca`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/c04abca))

**Full diff:** [`v4.6.2...v4.6.3`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/compare/v4.6.2...v4.6.3)

---

## [v4.6.2] - 2025-11-22 **[Tag]**

CI release-pipeline fix. No scanner changes.

- Fix jq download URL for release workflow ([`e10b60c`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/e10b60c))

**Full diff:** [`v4.6.1...v4.6.2`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/compare/v4.6.1...v4.6.2)

---

## [v4.6.1] - 2025-11-22 **[Tag]**

- Fix release workflow YAML and bump version ([`592486d`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/592486d))

**Full diff:** [`v4.6.0...v4.6.1`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/compare/v4.6.0...v4.6.1)

---

## [v4.6.0] - 2025-11-22 **[Tag]**

> First tagged version. Encompasses all development from the initial commit (2025-11-16) through the v5.0.0-era preparation work including the polyglot architecture, 8-language module system, and the full test suite.

### Initial Capabilities (v4.4 through v4.6.0)

#### Multi-Language Scanner Architecture

- Refactor from monolithic JS-only scanner to modular polyglot architecture supporting 8 languages: JavaScript/TypeScript, Python, Go, Rust, C++, Java, Ruby, Swift ([`d392951`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/d392951bc2dc67bd44b5502534f67f1768c23182))
- Each language module runs as an independent shell script with standardized output format

#### Language Modules

- **JavaScript/TypeScript** (`ubs-js.sh`): 1000+ bug patterns including `=== NaN`, missing `await`, `innerHTML` XSS, `parseInt` radix, React hook dependencies, `var` declarations, dangling promises
- **Python** (`ubs-python.sh`): mutable defaults, bare `except`, `eval()`/`exec()`, resource lifecycle analysis
- **Go** (`ubs-golang.sh`): goroutine leaks, context cancellation, `defer` in loops, `sync.Mutex` misuse, AST-based resource lifecycle analyzer ([`790a968`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/790a968))
- **Rust** (`ubs-rust.sh`): `.unwrap()` panics, async error detection, type narrowing
- **C++** (`ubs-cpp.sh`): `strcpy`/`sprintf` buffer overflows, raw pointer misuse, RAII checks, modern C++20 patterns
- **Java** (`ubs-java.sh`): JDBC resource lifecycle with Statement/ResultSet/CallableStatement tracking, AST-grep integration via Python parser ([`2b57508`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/2b5750862ba37053e95cbe9a4c958550daf7a983))
- **Ruby** (`ubs-ruby.sh`): unsafe YAML deserialization, command injection, block/ensure analysis
- **Swift** (`ubs-swift.sh`): guard-let analysis, Objective-C bridging patterns, 23 analysis categories ([`a0032c3`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/a0032c3))

#### AST-Powered Analysis

- AST-based resource lifecycle analyzers for Python ([`242c25f`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/242c25f169f110095fc7eecc5cd0fe03a7295dec)) and Go ([`790a968`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/790a968))
- ast-grep integration for Java via Python-based result parser ([`2b57508`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/2b5750862ba37053e95cbe9a4c958550daf7a983))
- Type narrowing helpers for TypeScript, Rust, Kotlin, Swift

#### Core Features

- Git-aware scanning: `--staged` mode for pre-commit checks
- Strictness profiles: `--profile=strict`, `--profile=lenient`
- Machine-readable output: `--format=json`
- Smart silence: suppress known-clean patterns
- `.ubsignore` file support for excluding paths and patterns
- `ubs doctor` diagnostics command
- `ubs sessions` session management

#### Installer

- curl|bash one-liner installer with ripgrep auto-detection ([`6c5bf4a`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/6c5bf4a))
- Homebrew bash detection and daily auto-updates ([`376feb6`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/376feb6))
- Bun support for TypeScript installation ([`250fd82`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/250fd82))
- Non-interactive uninstall support ([`81af909`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/81af909))
- Automatic bash upgrade on macOS ([`8b5b925`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/8b5b925))
- WSL detection and BSD platform support ([`60949b1`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/60949b1))
- Concurrency protection with lock file ([`263479c`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/263479c))

#### Test Suite

- Comprehensive polyglot test suite with buggy and clean code examples ([`a06ebce`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/a06ebcee951c53f23757d3bc5dbbb1f3867b723d))
- 275+ additional bug patterns across 12 test files ([`2a7ac95`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/2a7ac95))
- Framework-specific anti-patterns for React and Node.js/Express ([`7fa23cf`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/7fa23cf))
- Manifest-driven test automation infrastructure ([`bd09105`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/bd09105))
- Async error path coverage fixtures for all languages ([`12d3507`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/12d3507), [`f59787f`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/f59787f))
- Resource lifecycle test cases for all languages ([`5cfebca`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/5cfebca))

#### Metrics & Reporting

- Comprehensive metrics collection infrastructure ([`76b5024`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/76b5024))
- Metrics collection in JS, Python, and Ruby modules ([`fd19ad4`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/fd19ad4))

#### Documentation

- 12 installer bugs fixed ([`92b0c85`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/92b0c85a3330d8e92664a45787098a3484b9376f))
- Project justification and rationale section with FAQ ([`c841e66`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/c841e66))
- Release workflows and supply chain documentation ([`8144864`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/8144864))
- Minisign public key published in install instructions ([`2c549eb`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/2c549eb))

---

## Pre-v4.6.0 (v4.4 initial commit)

- **2025-11-16**: Initial release of Ultimate Bug Scanner v4.4 as a JavaScript/TypeScript-focused bug scanner ([`689c581`](https://github.com/Dicklesworthstone/ultimate_bug_scanner/commit/689c58100a0bd40d8bef0357aaa50c7a3a19f6f2))

---

<!-- Link references -->
[Unreleased]: https://github.com/Dicklesworthstone/ultimate_bug_scanner/compare/v5.0.6...HEAD
[v5.0.6]: https://github.com/Dicklesworthstone/ultimate_bug_scanner/releases/tag/v5.0.6
[v5.0.5]: https://github.com/Dicklesworthstone/ultimate_bug_scanner/releases/tag/v5.0.5
[v5.0.4]: https://github.com/Dicklesworthstone/ultimate_bug_scanner/releases/tag/v5.0.4
[v5.0.3]: https://github.com/Dicklesworthstone/ultimate_bug_scanner/releases/tag/v5.0.3
[v5.0.2]: https://github.com/Dicklesworthstone/ultimate_bug_scanner/releases/tag/v5.0.2
[v5.0.1]: https://github.com/Dicklesworthstone/ultimate_bug_scanner/releases/tag/v5.0.1
[v5.0.0]: https://github.com/Dicklesworthstone/ultimate_bug_scanner/releases/tag/v5.0.0
[v4.6.5]: https://github.com/Dicklesworthstone/ultimate_bug_scanner/releases/tag/v4.6.5
[v4.6.4]: https://github.com/Dicklesworthstone/ultimate_bug_scanner/releases/tag/v4.6.4
[v4.6.3]: https://github.com/Dicklesworthstone/ultimate_bug_scanner/compare/v4.6.2...v4.6.3
[v4.6.2]: https://github.com/Dicklesworthstone/ultimate_bug_scanner/compare/v4.6.1...v4.6.2
[v4.6.1]: https://github.com/Dicklesworthstone/ultimate_bug_scanner/compare/v4.6.0...v4.6.1
[v4.6.0]: https://github.com/Dicklesworthstone/ultimate_bug_scanner/compare/689c581...v4.6.0
