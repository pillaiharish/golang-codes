# Q24. Extract DBError through a chain with errors.As
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../23-traverse-error-chain/README.md) · [Next](../25-contextual-config-read/README.md)

**Level:** Intermediate | **Concepts:** typed wrappers, error metadata.

## Your task — code first

1. Implement DBError{Code,Err} with Unwrap.
2. Wrap it at two higher layers.
3. Extract Code with errors.As.

## Failure cases and acceptance tests

- [ ] DBError is found through two wrappers.
- [ ] Its nested sentinel remains inspectable.

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

