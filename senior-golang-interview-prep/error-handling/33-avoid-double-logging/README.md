# Q33. Choose one error logging boundary
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../32-table-tests-validation/README.md) · [Next](../34-structured-slog-errors/README.md)

**Level:** Upper intermediate | **Concepts:** log once, handlers, error propagation.

## Your task — code first

1. Implement a lower-layer function returning error with context.
2. Create an outer handler that logs once through an injected logger.
3. Test one log record for one failure.

## Failure cases and acceptance tests

- [ ] Failure produces one structured record.
- [ ] Error identity is preserved.

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

