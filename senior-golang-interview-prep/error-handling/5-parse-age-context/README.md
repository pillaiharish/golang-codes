# Q5. Add context to parsing failures

[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../4-new-vs-errorf/README.md) · [Next question](../6-do-not-ignore-errors/README.md)

**Level:** Starter | **Concepts:** strconv.Atoi, error chains, contextual messages.

## Your task — code first

1. Implement ParseAge(input string) (int,error) with strconv.Atoi.
2. Wrap the returned parsing error using %w.
3. Keep user input in the context without losing the original cause.

## Failure cases and acceptance tests

- [ ] 42 parses correctly.
- [ ] abc fails and errors.Is reaches strconv.ErrSyntax.

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

