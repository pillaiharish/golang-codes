# Q39. Build a small application error taxonomy
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../38-avoid-overwrapping/README.md) · [Next](../40-static-error-analysis/README.md)

**Level:** Upper intermediate | **Concepts:** AppError, error kinds, policy.

## Your task — code first

1. Define validation, not-found, conflict, internal and transient kinds.
2. Wrap underlying errors in AppError.
3. Map each kind to an appropriate HTTP status and decide retry policy.

## Failure cases and acceptance tests

- [ ] Validation maps to 400.
- [ ] Transient maps to 503; cause is inspectable.

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

