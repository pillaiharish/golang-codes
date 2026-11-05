# Q30. Debug and repair a broken error chain
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../29-multiple-percent-w/README.md) · [Next](../31-test-errors-structurally/README.md)

**Level:** Intermediate | **Concepts:** regression tests, sentinel identity, %v bug.

## Your task — code first

1. Run the broken function and observe errors.Is failure.
2. Change the formatting verb to %w.
3. Add a regression test checking identity after wrapping.

## Failure cases and acceptance tests

- [ ] Broken function fails structural matching.
- [ ] Fixed function passes.

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

