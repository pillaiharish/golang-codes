# Project 4: File processor CLI

[Project index](../README.md) · [Questions](../../README.md) · [Study plan](../../STUDY-PLAN.md)

**Level:** Upper intermediate. **Status:** unsolved practice challenge.

## Requirements

Process paths with context cancellation and slog; continue on missing, permission-denied or invalid-content items; stop on internal fatal error; report structured counts and nonzero fatal status.

## Required tests and failure injection

Success batch; missing file continues; synthetic permission error; invalid file continues; cancellation stops work; no duplicate logging; no leaked descriptors.

## Acceptance

- [ ] Independently implemented with runnable Go code and tests.
- [ ] All relevant failure cases have deterministic tests using fakes and temporary resources.
- [ ] Errors are classified structurally with errors.Is/As rather than text comparison.
- [ ] One boundary owns logging/exit/HTTP response; internal details do not leak.
- [ ] Cancellation, retries and partial failures have an explicit policy where applicable.
- [ ] Run go test, go vet, staticcheck and errcheck in the project's module.
- [ ] Record design trade-offs and actual evidence in [revision notes](../../NOTES.md).

