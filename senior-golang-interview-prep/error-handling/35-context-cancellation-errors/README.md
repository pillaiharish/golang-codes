# Q35. Return context cancellation and deadline errors
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../34-structured-slog-errors/README.md) · [Next](../36-wrap-context-errors/README.md)

**Level:** Upper intermediate | **Concepts:** context.Context, cancellation, deadlines.

## Your task — code first

1. Implement DoWork(ctx) with select on ctx.Done().
2. Test canceled and expired contexts deterministically.
3. Distinguish context.Canceled from context.DeadlineExceeded.

## Failure cases and acceptance tests

- [ ] Canceled context yields context.Canceled.
- [ ] Expired context yields context.DeadlineExceeded.

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

