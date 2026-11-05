# Q36. Wrap context errors without changing their policy
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../35-context-cancellation-errors/README.md) · [Next](../37-translate-third-party-errors/README.md)

**Level:** Upper intermediate | **Concepts:** cancellation, retry policy, context error identity.

## Your task — code first

1. Wrap ctx.Err() with operation context.
2. Test errors.Is against cancellation.
3. Explain why a caller cancellation should not be retried automatically.

## Failure cases and acceptance tests

- [ ] Wrapped canceled error matches context.Canceled.
- [ ] Successful context returns nil.

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

