# Q45. Implement bounded, cancellable retries
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../44-context-timeout/README.md) · [Next](../46-retryable-error-type/README.md)

**Level:** Advanced | **Concepts:** transient classification, exponential backoff, jitter, idempotency.

## Your task — code first

1. Implement WithRetry(ctx, attempts, backoff, fn).
2. Retry only errors marked transient, with bounded exponential backoff and jitter.
3. Stop on context cancellation, success, or a permanent error.

## Failure cases and acceptance tests

- [ ] Transient failures retry to success.
- [ ] Validation failures do not retry; cancellation stops waiting.

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

