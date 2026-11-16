# Practice 03 — predict, run and fix slice memory traps

[Layer 3 trap notes](../TRAPS-AND-MEMORY.md) · [Slice mastery](../SLICE-MASTERY.md) · [Recall drills](../RECALL-DRILLS.md)

## Predict before running

1. Appending to a subslice with spare capacity: can it overwrite original elements?
2. How do `slices.Clone` and `base[:2:2]` differ? Is an existing element still shared?
3. What is returned by copying into a slice with length 0 but capacity 3?
4. How do nil and empty slices marshal to JSON?
5. Does changing a ranged struct variable change the slice element?
6. Is cloning the outer `[][]int` slice enough to clone every row?
7. Why zero a pointer slot before shrinking the visible length?

## Observable results

```text
SharedAppend       -> base [1 2 9 4], part [1 2 9]
IsolatedAppend     -> base [1 2 3 4], part [1 2 9]
CapacityLimited    -> base [1 2 3 4], part [1 2 9]
ZeroLengthCopy     -> 0
NilEmptyJSON       -> null  []
```

The first case shares spare backing capacity. The clone provides independent int storage. The full-slice expression **only restricts capacity**, forcing reallocation on subsequent append when full; it does not copy existing elements. Shrinking doesn't automatically release pointers in backing storage.

## Run and test

```sh
GO111MODULE=off go run .
GO111MODULE=off go test -race -count=1 -v .
GO111MODULE=off go vet .
```

[Source](main.go) · [Tests](main_test.go). All Go files are compilable; intentionally invalid equality is documented **only** in the [trap notes](../TRAPS-AND-MEMORY.md).

## Cold recall

Write the three append variants from memory and explain each difference in backing storage. Write a safe pointer-slice deletion that clears obsolete tail. Then return to existing problems rather than copying a second #238 or #1293.
