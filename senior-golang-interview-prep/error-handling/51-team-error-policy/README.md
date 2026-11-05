# Q51. Define and enforce a team error policy
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../50-observable-error-metrics/README.md) · [Next](../52-commerce-error-taxonomy/README.md)

**Level:** Senior | **Concepts:** return, wrap, log, panic, public vs internal.

## Your task — code first

1. Write a one-page policy covering creation, propagation, classification, public messages, tests and tooling.
2. Implement a boundary helper that wraps without logging.
3. Identify which rules belong in CI and which need code review.

## Failure cases and acceptance tests

- [ ] An error is preserved at a boundary.
- [ ] A nil error stays nil.

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

