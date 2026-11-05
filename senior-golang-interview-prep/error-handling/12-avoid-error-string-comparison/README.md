# Q12. Replace message matching with error identity
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../11-sentinel-not-found/README.md) · [Next](../13-typed-validation-error/README.md)

**Level:** Beginner | **Concepts:** string comparisons, sentinels.

## Your task — code first

1. Demonstrate fragile err.Error() text comparison.
2. Refactor to errors.Is with a shared sentinel.
3. Wrap the error and show why errors.Is still works.

## Failure cases and acceptance tests

- [ ] Wrapped sentinel matches.
- [ ] A freshly constructed error with identical text is distinct.

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

