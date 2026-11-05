# Q22. Implement Error and Unwrap methods
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../21-percent-v-vs-percent-w/README.md) · [Next](../23-traverse-error-chain/README.md)

**Level:** Intermediate | **Concepts:** custom wrapper, Unwrap, nil handling.

## Your task — code first

1. Implement ContextError with Op and Err fields.
2. Add Error() and Unwrap() methods.
3. Use errors.Is to find a sentinel beneath your wrapper.

## Failure cases and acceptance tests

- [ ] Message has operation context.
- [ ] Underlying sentinel is reachable.

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

