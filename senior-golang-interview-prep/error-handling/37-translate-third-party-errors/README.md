# Q37. Translate third-party errors into domain errors
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../36-wrap-context-errors/README.md) · [Next](../38-avoid-overwrapping/README.md)

**Level:** Upper intermediate | **Concepts:** adapters, error codes, boundaries.

## Your task — code first

1. Simulate a library returning typed/code-bearing errors.
2. Map library missing-object to a domain ErrNotFound.
3. Propagate other failures without relying on text comparisons.

## Failure cases and acceptance tests

- [ ] Missing maps to domain sentinel.
- [ ] Other library errors preserve their original cause.

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

