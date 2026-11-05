# Q34. Log failures with structured slog fields
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../33-avoid-double-logging/README.md) · [Next](../35-context-cancellation-errors/README.md)

**Level:** Upper intermediate | **Concepts:** slog, identifiers, safe fields.

## Your task — code first

1. Log user_id, request_id, and error as separate structured fields.
2. Do not put secrets in log records.
3. Test fields with a JSON handler and an in-memory writer.

## Failure cases and acceptance tests

- [ ] Fields are emitted.
- [ ] No API key or private token is logged.

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

