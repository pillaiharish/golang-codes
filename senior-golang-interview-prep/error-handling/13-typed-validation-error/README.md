# Q13. Create and inspect a ValidationError
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../12-avoid-error-string-comparison/README.md) · [Next](../14-errors-is-wrapped-sentinel/README.md)

**Level:** Beginner | **Concepts:** custom types, fields, Error method.

## Your task — code first

1. Implement ValidationError{Field,Message} with Error().
2. Use it in ValidateAge.
3. Assert structured fields with errors.As.

## Failure cases and acceptance tests

- [ ] Valid ages pass.
- [ ] Negative ages produce a typed error with Field=age.

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

