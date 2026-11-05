# Q6. Show why ignoring errors hides failure

[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../5-parse-age-context/README.md) · [Next question](../7-defer-file-close/README.md)

**Level:** Starter | **Concepts:** blank identifier, error propagation.

## Your task — code first

1. Write an example that discards a parsing error.
2. Implement a safe function that validates the result before using it.
3. Explain the rare cases where explicitly discarding an error is justified.

## Failure cases and acceptance tests

- [ ] Valid text is parsed.
- [ ] Invalid text is rejected, never treated as zero.

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

