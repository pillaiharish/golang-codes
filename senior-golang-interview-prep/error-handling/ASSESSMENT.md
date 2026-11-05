# Senior error-handling readiness checklist

[Question index](README.md) · [Revision notes](NOTES.md) · [Capstone](projects/README.md)

For each row, explain one real failure scenario and point to a test, benchmark, or code review demonstrating the property.

| Area | What you should be able to defend |
|---|---|
| Creation | Sentinel versus typed versus contextual error; nil/interface semantics |
| Propagation | Where errors are wrapped, returned, joined, or intentionally translated |
| Inspection | errors.Is/As, Unwrap(), Unwrap() []error, custom Is |
| Testing | Table-driven failure tests, fake dependency failures, panic and cancellation boundaries |
| Logging | One owner per event, structured safe fields, bounded volume |
| Context | Cancellation and deadlines propagate; work actually terminates |
| Concurrency | Failure collection vs fail-fast; no blocked goroutines on early exit |
| Resilience | Retry budget, backoff + jitter, idempotency, timeout, breaker, shedding |
| APIs | Stable wire codes, HTTP/gRPC mapping, no internal leaks, backward compatibility |
| Reliability | SLI denominator, user-impact alerts, partial failures and shutdown |
| Team policy | Code review requirements, analyzers, migration and deprecation guidance |

## Design-review prompts

1. A payment gateway times out after accepting a charge. Is retry safe? What idempotency key and reconciliation policy are necessary?
2. A new custom error type replaces a sentinel. Which callers use errors.Is and which use direct equality? How will you migrate?
3. A downstream outage causes 100x retries and log floods. What are the immediately safe mitigations?
4. A service returns 500 for caller cancellation. What should its request metrics, traces, and handler semantics do instead?
5. A shared counter increments from many goroutines. Which test and analyzer would demonstrate safety?

## Exit bar

- [ ] Implement selected cold-recall questions without copying.
- [ ] Demonstrate tests for both failure and success paths.
- [ ] Explain a broken error chain, typed-nil interface, and multi-error traversal.
- [ ] Defend a bounded, idempotent retry policy.
- [ ] Review a service's error taxonomy, API contract and observability.
- [ ] Finish one capstone with failure injection and document trade-offs.

