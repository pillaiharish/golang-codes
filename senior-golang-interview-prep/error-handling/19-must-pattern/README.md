# Q19. Implement a documented Must helper
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../18-recover-deferred-boundary/README.md) · [Next](../20-table-driven-error-tests/README.md)

**Level:** Beginner | **Concepts:** initialization, MustDivide, panics.

## Your task — code first

1. Implement MustDivide using Divide.
2. Use it for a trusted constant during startup.
3. Test that dividing by zero panics instead of returning a value.

## Failure cases and acceptance tests

- [ ] Safe division succeeds.
- [ ] Zero divisor panics.

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

