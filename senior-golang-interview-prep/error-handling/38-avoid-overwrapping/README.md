# Q38. Avoid redundant context in error chains
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../37-translate-third-party-errors/README.md) · [Next](../39-classify-app-errors/README.md)

**Level:** Upper intermediate | **Concepts:** layer boundaries, context density, message noise.

## Your task — code first

1. Build one useful wrapper at a boundary.
2. Return the same error unchanged from a pass-through function.
3. Explain when preserving exact identity is preferable to another wrapper.

## Failure cases and acceptance tests

- [ ] Cause remains matchable.
- [ ] Pass-through does not add noise.

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

