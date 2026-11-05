# Q27. Collect field errors with errors.Join
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../26-layered-error-messages/README.md) · [Next](../28-join-nil-errors/README.md)

**Level:** Intermediate | **Concepts:** multi-error, ValidationError, Is and As.

## Your task — code first

1. Validate name and age independently.
2. Combine all errors using errors.Join.
3. Prove errors.As retrieves one typed error and callers can inspect all by multi-unwrapping.

## Failure cases and acceptance tests

- [ ] Valid user returns nil.
- [ ] Two invalid fields produce two children.

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

