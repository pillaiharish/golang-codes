# Project 6: Order/payment service

[Project index](../README.md) · [Questions](../../README.md) · [Study plan](../../STUDY-PLAN.md)

**Level:** Senior. **Status:** unsolved practice challenge.

## Requirements

Production-harden an order workflow with domain taxonomy, idempotency keys, API mapping, safe logs, metrics, deadlines, retry budget, circuit breaker, graceful shutdown and compensation/reconciliation.

## Required tests and failure injection

Inject payment gateway timeout after charge acceptance, duplicate request, inventory conflict, database failure after partial write, cancellation after charge, downstream outage; assert no duplicate charges and safe public errors.

## Acceptance

- [ ] Independently implemented with runnable Go code and tests.
- [ ] All relevant failure cases have deterministic tests using fakes and temporary resources.
- [ ] Errors are classified structurally with errors.Is/As rather than text comparison.
- [ ] One boundary owns logging/exit/HTTP response; internal details do not leak.
- [ ] Cancellation, retries and partial failures have an explicit policy where applicable.
- [ ] Run go test, go vet, staticcheck and errcheck in the project's module.
- [ ] Record design trade-offs and actual evidence in [revision notes](../../NOTES.md).

