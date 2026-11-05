# Error-handling anti-patterns and interview scenarios

[Question index](README.md) · [Assessment](ASSESSMENT.md)

For each case, **predict the behavior, reproduce it locally, write a regression test, then fix it**. The linked questions supply examples.

| Anti-pattern | Consequence | Safer approach | Practice |
|---|---|---|---|
| Discarding a returned error | Failure silently becomes an apparent zero/success | Check and propagate or explicitly justify best-effort ignore | [Q6](6-do-not-ignore-errors/README.md) |
| Logging and returning at each layer | Duplicate incidents, noisy alerts | Wrap inside, log once at boundary | [Q9](9-return-vs-log-error/README.md), [Q33](33-avoid-double-logging/README.md) |
| Matching err.Error() | Broken by text/context changes | Sentinel, errors.Is or typed errors.As | [Q12](12-avoid-error-string-comparison/README.md) |
| Using %v instead of %w | Lost inspectable cause | Wrap deliberately at a boundary | [Q21](21-percent-v-vs-percent-w/README.md) |
| Returning a typed nil pointer as error | err != nil despite nil concrete pointer | Return literal nil from the concrete-type boundary | [Q16](16-typed-nil-error/README.md) |
| Panic for client input | Request/service crash | Return validation error; boundary recovery as last resort | [Q17](17-panic-only-for-invariants/README.md) |
| Leaking SQL/stack details via API | Security and compatibility risk | Safe public code, internal structured logs, request ID | [Q47](47-http-domain-error-responses/README.md) |
| Wrapping every layer | Verbose redundant error text | Add unique context only where useful | [Q38](38-avoid-overwrapping/README.md) |

## Five scenario drills

1. **Nil-looking error:** a function returns a nil pointer but caller gets err != nil. Draw interface dynamic type/value; reproduce [Q16](16-typed-nil-error/README.md).
2. **errors.Is suddenly fails:** inspect %w vs %v, recreated sentinels, broken Unwrap, incorrect target; repair [Q30](30-debug-errors-is-failure/README.md).
3. **HTTP 500 leaks schema details:** map an internal error to a stable public code, log internally, add correlation ID; implement [Q47](47-http-domain-error-responses/README.md).
4. **Retry storm:** identify missing attempt cap, deadline, jitter, idempotency gate, circuit breaker and load shedding; test [Q45](45-bounded-retry-backoff/README.md) and [Q57](57-circuit-breaker-policy/README.md).
5. **Goroutine panic:** a recover in the parent goroutine cannot intercept a child panic. Install a recovery boundary within the worker only when policy permits survival; review [Q18](18-recover-deferred-boundary/README.md) and the existing [concurrency Q4](../concurrency/4-goroutine-panic-and-recover/README.md).

