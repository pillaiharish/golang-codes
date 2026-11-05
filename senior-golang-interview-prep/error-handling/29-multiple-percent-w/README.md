# Q29. Wrap multiple errors with fmt.Errorf
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../28-join-nil-errors/README.md) · [Next](../30-debug-errors-is-failure/README.md)

**Level:** Intermediate | **Concepts:** Go 1.20, multi-error trees.

## Your task — code first

1. Use two %w verbs in one fmt.Errorf call.
2. Check both sentinels with errors.Is.
3. Inspect the multi-Unwrap return shape.

## Failure cases and acceptance tests

- [ ] Both children match.
- [ ] Multi-Unwrap has two causes.

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

