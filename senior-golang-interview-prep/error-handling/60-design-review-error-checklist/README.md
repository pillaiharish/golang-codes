# Q60. Review error policy before implementation
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../59-legacy-error-migration/README.md)

**Level:** Senior | **Concepts:** failure matrix, SLOs, partial failures, shutdown.

## Your task — code first

1. Design a checklist for categories, retryability, timeouts, cancellation, partial failure, logging and tests.
2. Implement a small validator to identify missing design fields.
3. Write a review for a hypothetical service with an unbounded retry loop.

## Failure cases and acceptance tests

- [ ] Fully specified design passes.
- [ ] Missing timeout and tests are reported.

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

