# Q40. Use linters to catch error mistakes
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../39-classify-app-errors/README.md) · [Next](../41-goroutine-error-propagation/README.md)

**Level:** Upper intermediate | **Concepts:** go vet, staticcheck, errcheck, nilness.

## Your task — code first

1. Run go vet, staticcheck, and errcheck on the solution.
2. Explain which tool finds unchecked return values.
3. Identify why a deliberately ignored error should be documented and tested.

## Failure cases and acceptance tests

- [ ] Error returned from ReadFile is handled.
- [ ] No unchecked relevant I/O error paths.

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

