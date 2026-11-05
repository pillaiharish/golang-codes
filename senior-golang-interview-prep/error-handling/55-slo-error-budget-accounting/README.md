# Q55. Decide what counts against availability SLOs
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../54-cross-service-error-contract/README.md) · [Next](../56-debug-error-storm/README.md)

**Level:** Senior | **Concepts:** SLI denominator, client errors, retries.

## Your task — code first

1. Define a success-rate SLI from external requests.
2. Implement an example availability counter that excludes client validation errors.
3. Explain policy for cancellations, rate limiting and retried attempts.

## Failure cases and acceptance tests

- [ ] 5xx count as failure.
- [ ] 400 validation excluded from the example denominator.

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

