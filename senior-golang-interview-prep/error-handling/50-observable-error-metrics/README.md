# Q50. Count errors without high-cardinality labels
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../49-translate-sql-errors/README.md) · [Next](../51-team-error-policy/README.md)

**Level:** Advanced | **Concepts:** metrics, structured categories, alerting.

## Your task — code first

1. Count failures by bounded category.
2. Use concurrency-safe counter storage.
3. Explain which errors deserve pages versus dashboard-only counts.

## Failure cases and acceptance tests

- [ ] Counts increment for each kind.
- [ ] Concurrent updates are race-free.

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

