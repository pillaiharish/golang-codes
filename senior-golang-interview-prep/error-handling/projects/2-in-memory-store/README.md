# Project 2: In-memory user store

[Project index](../README.md) · [Questions](../../README.md) · [Study plan](../../STUDY-PLAN.md)

**Level:** Beginner. **Status:** unsolved practice challenge.

## Requirements

Implement User, Store{users map[int]User}, Get and Add; stable ErrNotFound and ErrDuplicateUser; typed ValidationError; callers use errors.Is and errors.As.

## Required tests and failure injection

Add/find success; missing ID; duplicate ID; invalid ID/name; verify nested error causes and no mutation after failed add.

## Acceptance

- [ ] Independently implemented with runnable Go code and tests.
- [ ] All relevant failure cases have deterministic tests using fakes and temporary resources.
- [ ] Errors are classified structurally with errors.Is/As rather than text comparison.
- [ ] One boundary owns logging/exit/HTTP response; internal details do not leak.
- [ ] Cancellation, retries and partial failures have an explicit policy where applicable.
- [ ] Run go test, go vet, staticcheck and errcheck in the project's module.
- [ ] Record design trade-offs and actual evidence in [revision notes](../../NOTES.md).

