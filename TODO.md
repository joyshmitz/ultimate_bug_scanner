# Ultimate Bug Scanner – Work Log & TODOs

All tasks reference Beads issue IDs so progress stays traceable. Update this list whenever you discover new work or finish a sub-task.

## 1. Root UBS triage (ultimate_bug_scanner-9dk)
- [x] Run per-language UBS scans to get baseline counts. (`./ubs --format=json --ci --only=<lang> .`)
- [x] Record baseline in notes/root-scan-2025-11-16.md.
- [ ] Triage JS critical categories (group by filename / category, identify quick wins).
- [ ] Create Beads sub-issues for JS hotspots (null safety, math pitfalls, parsing, security).
- [ ] Repeat triage for Python, Go, Rust, C++, Java, Ruby after JS plan is in place.

## 2. Manifest coverage expansion (ultimate_bug_scanner-d5z, ultimate_bug_scanner-aqd, ultimate_bug_scanner-o3l)
- [x] Go: enable buggy/clean manifest entries with `--only=golang` + substring requirement.
- [ ] Add substring/rule expectations for Rust/C++/Java/Ruby fixtures once ready.
- [ ] Create dedicated manifest cases for each language (buggy + clean).
- [ ] Add edge-case directories (unicode/timezone/fp) with explicit thresholds.
- [ ] Wire `test-suite/run_manifest.py` into CI so regressions fail PRs.

## 3. Framework & module hygiene (ultimate_bug_scanner-dmo)
- [ ] Fix `modules/ubs-js.sh` file counting so both module + meta-runner agree.
- [ ] Audit other modules for similar counting or summary issues.

## 4. Documentation / developer experience
- [ ] Merge Beads instructions + manifest workflow into README sections as they mature.
- [ ] Ensure AGENTS.md references Beads issue IDs whenever handoffs occur.

## 5. Resource lifecycle fixtures (ultimate_bug_scanner-6ig)
- [x] Investigate `modules/ubs-python.sh` single-file runs (resource_lifecycle) reporting zero files/warnings.
- [x] Do the same for Go and Java fixtures (confirm detection logic).
- [x] Restore warnings so `--fail-on-warning` triggers and manifest passes.

_Last updated: 2025-11-16 22:58 UTC_
