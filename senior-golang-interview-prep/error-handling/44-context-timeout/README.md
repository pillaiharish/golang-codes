# Q44. Distinguish request timeout from dependency failure
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../43-concurrent-multierror-validation/README.md) · [Next](../45-bounded-retry-backoff/README.md)

**Level:** Advanced | **Concepts:** context.WithTimeout, deadlines, retry decisions.

## Your task — code first

1. Create a 500ms timeout around slow work.
2. Check for context.DeadlineExceeded using errors.Is.
3. Explain client cancellations versus server-side deadlines.

## Failure cases and acceptance tests

- [ ] Fast work can complete.
- [ ] Slow work under a short deadline times out.

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

