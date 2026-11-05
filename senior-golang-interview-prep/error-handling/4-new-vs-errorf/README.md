# Q4. Compare errors.New and fmt.Errorf

[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../3-nil-error-zero-value/README.md) · [Next question](../5-parse-age-context/README.md)

**Level:** Starter | **Concepts:** static messages, dynamic format, %w vs %v.

## Your task — code first

1. Create errors using errors.New and fmt.Errorf.
2. Format a wrapped cause with both %v and %w.
3. Use errors.Is to prove which one preserves identity.

## Failure cases and acceptance tests

- [ ] errors.Is succeeds with %w.
- [ ] errors.Is fails with %v for a distinct wrapped error.

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

