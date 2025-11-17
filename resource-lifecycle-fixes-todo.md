# Resource Lifecycle + Java Clean TODO

## Diagnosis & Planning
- [x] Capture current failures for `rust-resource-lifecycle`, `java-clean`, `java-resource-lifecycle` (expected vs actual).
- [x] Identify which module heuristics/fixtures need adjustment per case (none needed after rerun; cases already passing).

## Rust Resource Lifecycle
- [x] Inspect `test-suite/rust/buggy/resource_lifecycle.rs` output (case currently passes; no changes).
- [x] Verify `modules/ubs-rust.sh` resource lifecycle logic counts warnings (confirmed via manifest case).
- [x] Implement code/fixture adjustments so warning count >= 1 (not needed; already satisfied).

## Java Clean Fixture
- [x] Review `test-suite/java/clean/...` output (case passes now; earlier failure not reproducible).
- [x] Update clean fixture or adjust module severity (not required).

## Java Resource Lifecycle
- [x] Ensure `test-suite/java/buggy/ResourceLifecycle.java` still triggers warnings (case passes).
- [x] Align module heuristic (no action needed).

## Verification & Wrap-up
- [x] Rerun targeted manifest cases for the three suites.
- [ ] Rerun full manifest (optional) or at least relevant subsets to ensure no regressions (full run timed out at 120s; rerun later if needed with longer window).
- [x] Update TODO file and summarize results in final report.
