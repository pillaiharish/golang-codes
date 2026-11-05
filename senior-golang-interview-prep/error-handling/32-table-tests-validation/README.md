# Q32. Use table-driven validation error tests
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../31-test-errors-structurally/README.md) · [Next](../33-avoid-double-logging/README.md)

**Level:** Upper intermediate | **Concepts:** subtests, validation cause, field names.

## Your task — code first

1. Implement ValidateEmail for empty, missing @, invalid domain and excessive length.
2. Use table-driven tests with clear case names.
3. Assert both a wrapped sentinel and ValidationError fields.

## Failure cases and acceptance tests

- [ ] A valid address passes.
- [ ] All four invalid forms return inspectable field-specific errors.

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

