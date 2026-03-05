# C# Fixtures

- `buggy/` contains intentionally unsafe patterns for `ubs-csharp.sh`.
- `clean/` provides counterexamples that should stay free of critical/warning findings.
- Manifest cases run with `--no-dotnet` so scanner regressions stay stable even when the .NET SDK is absent.
