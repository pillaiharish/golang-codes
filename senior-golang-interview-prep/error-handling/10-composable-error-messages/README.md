# Q10. Write composable Go error messages

[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../9-return-vs-log-error/README.md) · [Next question](../11-sentinel-not-found/README.md)

**Level:** Starter | **Concepts:** lowercase, punctuation, operation context.

## Your task — code first

1. Compare sentence-style errors with contextual error chains.
2. Write concise lower-case messages that compose.
3. Test the structural cause rather than brittle exact text.

## Failure cases and acceptance tests

- [ ] Message starts lower-case and has no terminal period.
- [ ] errors.Is still reaches underlying cause.

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

