# Q26. Structure concise error messages across layers
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../25-contextual-config-read/README.md) · [Next](../27-join-validation-errors/README.md)

**Level:** Intermediate | **Concepts:** operations, composing messages, redundancy.

## Your task — code first

1. Write a lower-layer error with one concrete cause.
2. Wrap it at a boundary with one operation name.
3. Compare it against repetitive nested 'failed to' messages.

## Failure cases and acceptance tests

- [ ] Combined text includes meaningful context.
- [ ] Underling sentinel is retained.

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

