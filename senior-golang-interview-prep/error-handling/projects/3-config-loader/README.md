# Project 3: Configuration loader

[Project index](../README.md) · [Questions](../../README.md) · [Study plan](../../STUDY-PLAN.md)

**Level:** Intermediate. **Status:** unsolved practice challenge.

## Requirements

Implement LoadConfig(path) returning *Config; map missing file to ErrConfigNotFound, malformed/invalid content to ErrConfigInvalid, and combine field validation with errors.Join.

## Required tests and failure injection

Valid config; nonexistent path; syntax error; multiple invalid fields; errors.Is for domain categories and errors.As for a field error; include path in message.

## Acceptance

- [ ] Independently implemented with runnable Go code and tests.
- [ ] All relevant failure cases have deterministic tests using fakes and temporary resources.
- [ ] Errors are classified structurally with errors.Is/As rather than text comparison.
- [ ] One boundary owns logging/exit/HTTP response; internal details do not leak.
- [ ] Cancellation, retries and partial failures have an explicit policy where applicable.
- [ ] Run go test, go vet, staticcheck and errcheck in the project's module.
- [ ] Record design trade-offs and actual evidence in [revision notes](../../NOTES.md).

