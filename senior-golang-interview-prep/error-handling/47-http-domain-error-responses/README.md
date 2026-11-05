# Q47. Map domain errors to safe HTTP JSON
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../46-retryable-error-type/README.md) · [Next](../48-grpc-status-errors/README.md)

**Level:** Advanced | **Concepts:** HTTP statuses, public codes, request ID, leak prevention.

## Your task — code first

1. Create an HTTP error handler that maps missing, validation, conflict, unauthorized and internal errors.
2. Return safe JSON with request_id.
3. Do not expose raw storage errors to clients.

## Failure cases and acceptance tests

- [ ] Not found maps to 404.
- [ ] An internal database error maps to 500 without leaking text.

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

