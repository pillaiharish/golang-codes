# Q18. Recover only in the panicking goroutine
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../17-panic-only-for-invariants/README.md) · [Next](../19-must-pattern/README.md)

**Level:** Beginner | **Concepts:** recover, deferred functions, panic boundaries.

## Your task — code first

1. Trigger a panic and recover using a deferred function.
2. Convert the recovered value into a return result.
3. Explain why recovery in another goroutine cannot catch it.

## Failure cases and acceptance tests

- [ ] Panic value is captured.
- [ ] No panic returns nil recovery value.

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

