# Q56. Triage a downstream failure without log flooding
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../55-slo-error-budget-accounting/README.md) · [Next](../57-circuit-breaker-policy/README.md)

**Level:** Senior | **Concepts:** error storms, sampling, upstream tracing, load shedding.

## Your task — code first

1. Write a triage plan comparing dependency latency, error rate, saturation and deployment changes.
2. Implement a small log sampling decision by occurrence count.
3. Explain why metrics must retain full counts even when logs are sampled.

## Failure cases and acceptance tests

- [ ] First failure is logged.
- [ ] Subsequent events are sampled while metrics can still count all.

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

