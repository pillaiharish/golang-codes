# Q16. Reproduce and fix a typed-nil error
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../15-errors-as-typed/README.md) · [Next](../17-panic-only-for-invariants/README.md)

**Level:** Beginner | **Concepts:** interfaces, typed nil, panic hazards.

## Your task — code first

1. Return a nil *MyError through an error interface and inspect err!=nil.
2. Implement a fixed function that returns a genuinely nil interface.
3. Explain why calling Error() on a nil pointer receiver may panic.

## Failure cases and acceptance tests

- [ ] Broken function returns non-nil interface.
- [ ] Fixed function returns nil.

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

