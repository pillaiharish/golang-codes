# Q11. Return a sentinel error from FindUser
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../10-composable-error-messages/README.md) · [Next](../12-avoid-error-string-comparison/README.md)

**Level:** Beginner | **Concepts:** sentinels, error identity, exported names.

## Your task — code first

1. Define one package-level ErrNotFound.
2. Implement FindUser(id int) (User,error) over a small in-memory lookup.
3. Explain why the shared sentinel can be compared directly until wrapping is introduced.

## Failure cases and acceptance tests

- [ ] Known ID returns the user.
- [ ] Unknown ID matches ErrNotFound.

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

