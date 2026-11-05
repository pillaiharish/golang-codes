# Q23. Trace errors.Is through multiple wrappers
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../22-custom-context-wrapper/README.md) · [Next](../24-extract-db-error/README.md)

**Level:** Intermediate | **Concepts:** chain traversal, broken chains.

## Your task — code first

1. Wrap a base sentinel twice with %w.
2. Show that errors.Is finds the deepest cause.
3. Replace one %w with %v to demonstrate the broken chain.

## Failure cases and acceptance tests

- [ ] Intact chain matches base.
- [ ] Broken middle link stops matching.

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

