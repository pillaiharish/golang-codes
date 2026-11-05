# Q46. Classify retryability structurally
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../45-bounded-retry-backoff/README.md) · [Next](../47-http-domain-error-responses/README.md)

**Level:** Advanced | **Concepts:** errors.As, transient policy, context handling.

## Your task — code first

1. Define RetryableError that wraps a cause.
2. Implement IsRetryable using errors.As.
3. Ensure context cancellation is not classified as retryable even when wrapped.

## Failure cases and acceptance tests

- [ ] Wrapped retryable error matches.
- [ ] Canceled context is not retryable.

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

