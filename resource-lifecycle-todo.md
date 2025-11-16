# Resource Lifecycle Correlation TODO

1. ✅ Capture plan requirements from PLAN_FOR_NEXT_FEATURES for ast-grep-based resource correlation.
2. ✅ Add ast-grep resource-correlation rules in every module (js, python, golang, cpp, rust, java, ruby).
3. ✅ Wire new rule files into module rule-pack builders so `run_ast_rules` emits correlation hits.
4. ✅ Extend each per-language buggy fixture set with a minimal resource-leak sample and clean counterparts if needed.
5. ✅ Update `test-suite/manifest.json` to add cases validating the new resource lifecycle category/order for every language.
6. ☐ Run representative UBS scans (or targeted manifest runner) to ensure new rules execute without crashing and report expected findings.
7. ☐ Mark TODO entries complete as each step finishes.
