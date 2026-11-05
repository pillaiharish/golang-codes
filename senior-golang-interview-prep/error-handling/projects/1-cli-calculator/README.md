# Project 1: CLI calculator

[Project index](../README.md) · [Questions](../../README.md) · [Study plan](../../STUDY-PLAN.md)

**Level:** Starter. **Status:** unsolved practice challenge.

## Requirements

Parse add/divide flags or positional arguments; return contextual errors for missing or non-numeric input; divide by zero is a checked error; main alone chooses nonzero exit status.

## Required tests and failure injection

Success: 2+3 and 10/2; failures: no operation, bad number, division by zero, unexpected argument; assert exit statuses and stderr without calling os.Exit from testable logic.

## Acceptance

- [ ] Independently implemented with runnable Go code and tests.
- [ ] All relevant failure cases have deterministic tests using fakes and temporary resources.
- [ ] Errors are classified structurally with errors.Is/As rather than text comparison.
- [ ] One boundary owns logging/exit/HTTP response; internal details do not leak.
- [ ] Cancellation, retries and partial failures have an explicit policy where applicable.
- [ ] Run go test, go vet, staticcheck and errcheck in the project's module.
- [ ] Record design trade-offs and actual evidence in [revision notes](../../NOTES.md).

