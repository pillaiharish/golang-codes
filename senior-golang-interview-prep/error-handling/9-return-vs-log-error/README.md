# Q9. Return errors and log once at the boundary

[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../8-cli-error-exit/README.md) · [Next question](../10-composable-error-messages/README.md)

**Level:** Starter | **Concepts:** error ownership, log-and-return antipattern.

## Your task — code first

1. Write a function that simply returns a wrapped lower-level error.
2. Keep logging in main or an outer handler.
3. Explain why logging at every layer duplicates incidents.

## Failure cases and acceptance tests

- [ ] Returned error remains inspectable.
- [ ] Lower-level functions do not log.

## Explain before revealing

- What is the expected behavior on success and failure?
- Which concrete error value or type should a caller be able to inspect?
- Where should this error be handled, wrapped, logged, or retried?
- What is one implementation that looks correct but loses information?

## Practice workflow

1. Implement a Go solution without opening the answer.
2. Write a failure-case test first and then the success case.
3. Compare with [the worked solution](solution/README.md); refactor your attempt.
4. Run tests and static checks, then record evidence in [revision notes](../NOTES.md).

**Mastery is not implied by the reference code.** Mark this exercise complete only after recreating and explaining it.

