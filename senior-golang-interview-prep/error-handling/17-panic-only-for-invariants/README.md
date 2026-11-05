# Q17. Use panic only for programmer invariants
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../16-typed-nil-error/README.md) · [Next](../18-recover-deferred-boundary/README.md)

**Level:** Beginner | **Concepts:** panic, Must prefix, recover in tests.

## Your task — code first

1. Write MustParseInt that panics on impossible programmer-controlled input.
2. Keep a regular parsing function that returns an error for user input.
3. Test panic and normal behavior.

## Failure cases and acceptance tests

- [ ] Valid value returns integer.
- [ ] Invalid Must input panics.

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

