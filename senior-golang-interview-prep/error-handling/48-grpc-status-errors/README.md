# Q48. Return idiomatic gRPC status codes
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../47-http-domain-error-responses/README.md) · [Next](../49-translate-sql-errors/README.md)

**Level:** Advanced | **Concepts:** grpc/status, codes, retry semantics.

## Your task — code first

1. Return codes.NotFound, codes.InvalidArgument and codes.Internal.
2. Check status.Code without comparing text.
3. Explain how to add structured details and select retryable codes.

## Failure cases and acceptance tests

- [ ] NotFound is classifiable.
- [ ] InvalidArgument remains distinct from Internal.

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

