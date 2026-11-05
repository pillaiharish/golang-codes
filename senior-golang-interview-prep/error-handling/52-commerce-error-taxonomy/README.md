# Q52. Design a domain error taxonomy for commerce
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../51-team-error-policy/README.md) · [Next](../53-backward-compatible-error-types/README.md)

**Level:** Senior | **Concepts:** client errors, server errors, retry classes, SLO.

## Your task — code first

1. Enumerate validation/authentication/authorization/not-found/conflict/quota/rate-limit/dependency/internal errors.
2. Implement mapping to HTTP and retryable status.
3. Explain which classes should contribute to an availability SLO.

## Failure cases and acceptance tests

- [ ] Validation returns 400 and no retry.
- [ ] Dependency unavailable returns 503 and is retryable only for safe operations.

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

