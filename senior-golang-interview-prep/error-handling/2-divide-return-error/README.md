# Q2. Return and check an error from Divide

[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../1-error-interface/README.md) · [Next question](../3-nil-error-zero-value/README.md)

**Level:** Starter | **Concepts:** multiple return values, errors.New, zero denominator.

## Your task — code first

1. Implement Divide(a, b float64) (float64, error).
2. Reject a zero divisor without panicking.
3. Test both success and failure, including negative numbers.

## Failure cases and acceptance tests

- [ ] 10/2 returns 5 and nil.
- [ ] 10/0 returns an error and a zero result.

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

