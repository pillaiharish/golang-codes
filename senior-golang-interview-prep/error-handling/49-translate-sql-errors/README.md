# Q49. Translate sql.ErrNoRows at repository boundary
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../48-grpc-status-errors/README.md) · [Next](../50-observable-error-metrics/README.md)

**Level:** Advanced | **Concepts:** database/sql, domain errors, cause preservation.

## Your task — code first

1. Translate sql.ErrNoRows to a domain ErrNotFound.
2. Wrap unknown DB errors with operation context.
3. Explain when raw driver errors are an implementation detail.

## Failure cases and acceptance tests

- [ ] No rows maps to ErrNotFound.
- [ ] Connection failure remains inspectable.

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

