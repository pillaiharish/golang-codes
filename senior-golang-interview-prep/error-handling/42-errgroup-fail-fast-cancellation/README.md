# Q42. Cancel sibling tasks using errgroup
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../41-goroutine-error-propagation/README.md) · [Next](../43-concurrent-multierror-validation/README.md)

**Level:** Advanced | **Concepts:** golang.org/x/sync/errgroup, derived context.

## Your task — code first

1. Start two tasks using errgroup.WithContext.
2. Make the first task return an error and the other stop on ctx.Done().
3. Explain why tasks must accept the derived context.

## Failure cases and acceptance tests

- [ ] First error is returned.
- [ ] Sibling observes cancellation and exits.

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

