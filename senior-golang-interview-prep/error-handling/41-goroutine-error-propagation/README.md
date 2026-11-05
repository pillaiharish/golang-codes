# Q41. Collect errors from goroutines safely
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../40-static-error-analysis/README.md) · [Next](../42-errgroup-fail-fast-cancellation/README.md)

**Level:** Advanced | **Concepts:** buffered channel, WaitGroup, errors.Join.

## Your task — code first

1. Launch three failing/succeeding jobs.
2. Use a buffered error channel with enough capacity and wait for completion.
3. Return all errors with errors.Join and explain first-error alternatives.

## Failure cases and acceptance tests

- [ ] Two causes are inspectable.
- [ ] No sender blocks when the receiver waits for all workers.

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

