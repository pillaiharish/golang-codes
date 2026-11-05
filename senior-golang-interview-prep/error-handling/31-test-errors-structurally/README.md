# Q31. Assert error identity and type in tests
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../30-debug-errors-is-failure/README.md) · [Next](../32-table-tests-validation/README.md)

**Level:** Upper intermediate | **Concepts:** errors.Is, errors.As, message stability.

## Your task — code first

1. Build a function returning a wrapped sentinel and typed context.
2. Test classification without exact message matching.
3. Explain when a full-message assertion is actually a public contract.

## Failure cases and acceptance tests

- [ ] Sentinel is present.
- [ ] Typed context is extractable.

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

