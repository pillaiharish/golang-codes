# Q54. Design safe errors across service boundaries
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../53-backward-compatible-error-types/README.md) · [Next](../55-slo-error-budget-accounting/README.md)

**Level:** Senior | **Concepts:** stable codes, request IDs, internal logging.

## Your task — code first

1. Define a JSON wire error contract with code, message, retryable and request_id.
2. Map an internal cause to a safe wire code.
3. Decode it on a client without parsing message text.

## Failure cases and acceptance tests

- [ ] Stable code survives JSON roundtrip.
- [ ] Internal dependency text is absent from serialized response.

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

