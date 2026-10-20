# Module 8 — Personal clarity and revision notes

[← Module 8](README.md) · [Preparation index](../README.md) · [Revision log](../REVISION-LOG.md)

**Not mastered by documentation alone.** Use this page to record predictions, failed experiments, and revised understanding.

## Current revision status

- Confidence (1–5): `—`
- Last reviewed (YYYY-MM-DD): `—`
- Next review (YYYY-MM-DD): `—`
- Current status: `Not yet self-assessed`

## Targeted foundations to test

| Concept | Predict first (without looking) | Evidence to record | Confidence (1–5) |
|---|---|---|---|
| State lifetime | Local map vs instance field across repeated calls | What survived and why? | — |
| Struct creation | `var`, `T{}`, `&T{}`, `new(T)`, `NewT()` | Types, zero values, initialization | — |
| Nil map | Read absent key vs write | Exact panic point | — |
| Constructor contract | What is bypassed with a raw literal? | Corrected construction code | — |
| Value/pointer receiver | Assign field vs mutate map entry | Original vs copy values | — |
| Struct copy | Copy initialized map vs two constructors | Alias observations | — |
| Method set | Concrete method call vs interface assignment | Which compiles and why? | — |
| Concurrency | Unsynchronized counter vs mutex | Tests and race-detector evidence | — |

## Recall history

| Review date | Missed explanation / syntax | Code or test evidence | What changed | Next review |
|---|---|---|---|---|
| _YYYY-MM-DD_ | | | | |

## My 60-second interview answer

> Write from memory: what a struct is, what a constructor does and does not do, why a map must be initialized before writes, how method calls preserve state, and how value/pointer receiver/copy semantics affect the result.

## Next independent exercise

Implement a `SessionStore` with `map[string]string`, `NewSessionStore()`, `Put`, `Get(key) (string, bool)`, and two independent instances. Then make the zero value usable and describe (but do not silently assume) a concurrency policy.
