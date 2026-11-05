# Q14. Match wrapped sentinels with errors.Is
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../13-typed-validation-error/README.md) · [Next](../15-errors-as-typed/README.md)

**Level:** Beginner | **Concepts:** error wrapping, identity traversal.

## Your task — code first

1. Create ErrNotFound and wrap it at a repository boundary.
2. Show direct equality fails for the wrapper.
3. Show errors.Is succeeds.

## Failure cases and acceptance tests

- [ ] Wrapped error is not directly equal to sentinel.
- [ ] errors.Is finds it.

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

