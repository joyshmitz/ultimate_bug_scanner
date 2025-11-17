# Resource Lifecycle + Java Clean TODO

## Diagnosis & Planning
- [ ] Capture current failures for `rust-resource-lifecycle`, `java-clean`, `java-resource-lifecycle` (expected vs actual).
- [ ] Identify which module heuristics/fixtures need adjustment per case.

## Rust Resource Lifecycle
- [ ] Inspect `test-suite/rust/buggy/resource_lifecycle.rs` to ensure it still triggers warnings (maybe fixes needed).
- [ ] Verify `modules/ubs-rust.sh` resource lifecycle logic counts files/warnings for that fixture.
- [ ] Implement code/fixture adjustments so warning count >= 1.

## Java Clean Fixture
- [ ] Review `test-suite/java/clean/...` output to see why warnings/critical > thresholds (maybe new async coverage warns?).
- [ ] Update clean fixture or adjust module severity to keep clean set quiet.

## Java Resource Lifecycle
- [ ] Ensure `test-suite/java/buggy/ResourceLifecycle.java` still triggers warnings (maybe file renamed?)
- [ ] Align module heuristic so warnings appear.

## Verification & Wrap-up
- [ ] Rerun targeted manifest cases for the three suites.
- [ ] Rerun full manifest (optional) or at least relevant subsets to ensure no regressions.
- [ ] Update TODO file and summarize results in final report.
