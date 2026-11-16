# Mastery card template — <concept>

[Mastery system](MASTERY-SYSTEM.md)

> For algorithms, mark data-structure-only items **N/A** instead of inventing operations. **Do not fill advanced sections prematurely.** Link working code; do not paste another answer to a solved problem.

## Quick recall (60 seconds)
- What is it, and when do I choose it?
- Core Go declaration / 5-line template:
- Invariant:
- O(…) time, O(…) space:
- Most common mistake:

## A. Identity and representation — Layer 2
Category (primitive/composite/ADT/concrete DS/algorithm/pattern), Go representation, standard library alternatives.

## B. Zero values and creation — Layer 2
Zero value, default element values, **is zero value usable?**, literals, make/new/constructor, validity of empty input.

## C. Full operations lifecycle — Layer 2
Access, traverse (order guarantee), add/insert, remove/delete, push/pop if applicable, search, sort, reverse, copy/clone, size/len/cap/empty/full, conversions. State **N/A** where appropriate; give a per-operation complexity table.

## D. Ownership and memory correctness — Layer 2/3
Assignment vs deep copy, aliasing/backing storage, comparable/map-key eligible, retention and cleanup, argument passing, race/concurrency safety. Distinguish `clear(s)` from dropping a reference; cleanup may require clearing pointer slots.

## E. Invariants, traps and edge cases — Layer 2/3
1–5 invariants; nil/empty, singleton, overflow, repeated states, invalid indices, ragged grids, cycles, broken vs fixed proof. Label intentionally failing code.

## F. Progressive implementations — Layer 2/4
**Preserve learner-original version**, simple baseline, alternative with rationale, complete runnable files, table tests and run commands.

## G. LeetCode recognition and transfer — Layer 4/6
Recognition clue → required state → algorithm → example related problem **links** (do not duplicate implementations), problem-specific traps.

## H. Go design — later
Encapsulation, zero-value API choice, invariants enforced by methods, constructors, error/ok contract, alternatives and when not to use it. Stub if not studied.

## I. Profiling and optimization — later
Only measured improvements: capacity planning, retention, allocations, cache locality, benchmarks, pprof/escape analysis. Stub until correct.

## J. Recall (all layers)
- 3 quick questions with answers hidden lower in doc or tests.
- Blank-file exercise, one edge-case challenge, one transfer exercise.
- Review on D1/D3/D7/D14/D30 (adjust after errors).
- Confidence/last review/next review: leave **blank** until an actual attempt.
