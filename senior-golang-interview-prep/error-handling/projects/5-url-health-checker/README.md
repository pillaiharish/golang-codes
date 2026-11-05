# Project 5: Concurrent URL health checker

[Project index](../README.md) · [Questions](../../README.md) · [Study plan](../../STUDY-PLAN.md)

**Level:** Advanced. **Status:** unsolved practice challenge.

## Requirements

Expose POST /check JSON; run bounded concurrent URL probes; timeouts and cancellation; retries for transient failures only; safe error codes/latencies in results; do not leak transport internals.

## Required tests and failure injection

Use httptest servers/fake transport, never real Internet; success, DNS/error classification, refused connection, timeout, cancellation, retry cap, invalid URL, all goroutines exit.

## Acceptance

- [ ] Independently implemented with runnable Go code and tests.
- [ ] All relevant failure cases have deterministic tests using fakes and temporary resources.
- [ ] Errors are classified structurally with errors.Is/As rather than text comparison.
- [ ] One boundary owns logging/exit/HTTP response; internal details do not leak.
- [ ] Cancellation, retries and partial failures have an explicit policy where applicable.
- [ ] Run go test, go vet, staticcheck and errcheck in the project's module.
- [ ] Record design trade-offs and actual evidence in [revision notes](../../NOTES.md).

