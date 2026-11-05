# Q21. Compare %v and %w on an error chain
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../20-table-driven-error-tests/README.md) · [Next](../22-custom-context-wrapper/README.md)

**Level:** Intermediate | **Concepts:** fmt.Errorf, errors.Is, Unwrap.

## Your task — code first

1. Wrap a sentinel with %v and %w in separate functions.
2. Inspect both with errors.Is and errors.Unwrap.
3. Explain API compatibility consequences of choosing %w.

## Failure cases and acceptance tests

- [ ] %v does not expose the cause.
- [ ] %w does expose the cause.

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

