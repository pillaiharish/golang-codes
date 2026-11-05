# Q43. Join validation errors from a worker pool
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../42-errgroup-fail-fast-cancellation/README.md) · [Next](../44-context-timeout/README.md)

**Level:** Advanced | **Concepts:** concurrency, errors.Join, bounded workers.

## Your task — code first

1. Validate several independent items using at most two workers.
2. Collect all field validation errors.
3. Explain why collect-all can be preferable to fail-fast for input validation.

## Failure cases and acceptance tests

- [ ] All invalid items contribute an error.
- [ ] Empty/valid input returns nil.

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

