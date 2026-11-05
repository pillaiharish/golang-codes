# Q15. Extract typed errors with errors.As
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../14-errors-is-wrapped-sentinel/README.md) · [Next](../16-typed-nil-error/README.md)

**Level:** Beginner | **Concepts:** typed errors, pointer to target, wrapping.

## Your task — code first

1. Create a typed validation error and wrap it.
2. Use var target *ValidationError and errors.As(err, &target).
3. Explain why a pointer to the target variable is required.

## Failure cases and acceptance tests

- [ ] errors.As finds the typed error through wrapping.
- [ ] Its Field is available.

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

