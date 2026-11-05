# Q57. Implement a minimal circuit breaker state machine
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../56-debug-error-storm/README.md) · [Next](../58-benchmark-error-cost/README.md)

**Level:** Senior | **Concepts:** closed/open/half-open, thresholds, timeouts, bulkheads.

## Your task — code first

1. Open a circuit after a configurable consecutive failure threshold.
2. Allow one half-open probe after the cooldown.
3. Explain where timeouts, retries, load shedding and fallbacks fit.

## Failure cases and acceptance tests

- [ ] Threshold opens the breaker.
- [ ] Cooldown enters half-open; a successful probe closes.

## Explain before revealing

- What is the behavior on both success and failure?
- Which error identity or typed details can the caller inspect?
- Where does wrapping, logging, cancellation or retry policy belong?
- Which apparently reasonable implementation could conceal a failure?

## Practice workflow

1. Implement independently without opening the answer.
2. Write tests for expected failures and successes.
3. Compare with the [worked solution](solution/README.md), then refactor.
4. Run tests and static checks; record evidence in [revision notes](../NOTES.md).

**Mastery is earned by reproducing the solution and explaining its trade-offs**, not by reading the reference.

