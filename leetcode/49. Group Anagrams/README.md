# 49 — Group Anagrams

## Before you solve: prerequisite basics

Study these in sequence and type the tiny Go example in each before reopening this solution:

1. [03 maps sets frequency and grouping](../0-problem-solving-building-blocks/03-maps-sets-frequency-and-grouping/README.md)
2. [04 strings bytes runes and parsing](../0-problem-solving-building-blocks/04-strings-bytes-runes-and-parsing/README.md)
3. [05 sorting and comparators](../0-problem-solving-building-blocks/05-sorting-and-comparators/README.md)
4. [13 tries and fixed alphabet indexing](../0-problem-solving-building-blocks/13-tries-and-fixed-alphabet-indexing/README.md)

[Full mapping for all 27 LeetCode problems](../0-problem-solving-building-blocks/README.md#b-problem--exact-prerequisites-all-27-problem-folders) · [Practice log](../REVISION-LOG.md)

[← Index](../README.md) · [Building block: canonical signatures](../0-problem-solving-building-blocks/README.md#3-frequency-counts-are-a-signature-for-anagrams) · [Original solutions](main.go) · [New frequency-key variant](variants.go) · [Tests](main_test.go) · [Revision log](../REVISION-LOG.md) · [LeetCode](https://leetcode.com/problems/group-anagrams/)

## Recognition

**Clue:** "Group words made from the same letters, ignoring order." Instead of comparing every pair of words, compute an identical **signature** for all anagrams and use a map to group the original words.

**Mental trigger:** **Normalize → group by signature → collect map values.**

## Start with the simplest mental picture

For `["eat","tea","tan","ate","nat","bat"]`, sorting each word gives:

```text
eat → aet       tea → aet       ate → aet
tan → ant       nat → ant       bat → abt
```

Use `map[string][]string` to append each **original word** under its **sorted key**.

## Invariant

After processing i words, each map bucket contains exactly the processed original words whose sorted-character signature matches the bucket's key. No two words with different sorted signatures can be grouped together.

## Preserve and compare the three forms

| Implementation | Present in repository | Explanation |
|---|---|---|
| `groupAnagrams` | Existing `main.go` | `strings.Split`, `sort.Slice`, `strings.Join`, map append; also prints debugging output |
| `groupAnagrams2` | Existing `main.go` | Helper `sortString` uses `sort.Strings`; otherwise same signature approach |
| Supplied new version | Equivalent to `groupAnagrams2` | Inlines `strings.Split → sort.Strings → strings.Join`; it is a clearer narrative, **not a different algorithm** |
| `groupAnagramsFreqKey` | Added `variants.go` | Array of 26 letter counts as a comparable map key, avoiding sorting under lowercase a-z input |

The original implementations are deliberately **kept intact**. The new submitted version is documented, not duplicated as an additional indistinguishable function.

## Reconstruct the sorted-key method

1. Create `groups := make(map[string][]string)`.
2. For every word s, split into characters.
3. Sort the character slice alphabetically.
4. Join it back into a canonical key.
5. Append the **unsorted original** s to `groups[key]`.
6. Iterate over map values to create `[][]string`.

### Dry run of the map

| Word | Sorted key | Groups after insertion |
|---|---|---|
| eat | aet | aet:[eat] |
| tea | aet | aet:[eat,tea] |
| tan | ant | aet:[eat,tea]; ant:[tan] |
| ate | aet | aet:[eat,tea,ate]; ant:[tan] |
| nat | ant | aet:[eat,tea,ate]; ant:[tan,nat] |
| bat | abt | plus abt:[bat] |

### Ordering caveat

**Go map iteration order is not guaranteed.** The outer order of the returned groups can change between runs. That is okay for LeetCode 49. Tests must compare normalized groups, not demand a particular outer ordering. Individual words within a bucket retain their input order because of append.

## Alternate signature — letter counts

For lowercase a-z only:

```go
var count [26]int
for i:=0; i<len(word); i++ { count[word[i]-'a']++ }
groups[count] = append(groups[count],word)
```

Arrays are comparable and can be map keys; slices cannot. Two words have the same `[26]int` exactly when every letter frequency matches.

## Complexity / limitations

Let n be number of words, k be their maximum length, and T the sum of their lengths.

- Sorting per word: O(n·k log k) upper-bound time, O(T) for stored outputs/signatures plus sort/split scratch space. More precisely sum O(length_i log length_i).
- Fixed frequency key: O(T + 26n) time (often written O(T) with fixed alphabet); O(T + 26g) space for original output strings and g distinct keys.
- The frequency-key variant assumes lowercase ASCII **a-z**. The sorted-string method can be adapted to broader character sets; Go `strings.Split(s,"")` operates by Unicode code points, not arbitrary grapheme clusters.
- Existing `groupAnagrams` prints debugging output; prefer `groupAnagrams2` or the new frequency variant in production.

## 30-second spoken answer

> I compute a canonical representation of each word: sort its characters so every anagram shares the same string key. A map from that key to a slice of original words collects groups, which I return as the values. For n words of length k it takes O(n k log k) time. If input is guaranteed lowercase a-z, I can replace sorting with a 26-count array key, making the pass linear in total input size.

## Cold recall

1. Why do we append the **original word** instead of the sorted key?
2. How do you create a key that all anagrams share?
3. Why is `sort.Strings` vs `sort.Slice` not a real algorithmic improvement here?
4. Why may outer group ordering differ between runs?
5. When is a frequency array key preferable, and what assumption does it require?
6. Solve [Valid Anagram #242](../242.%20Valid%20Anagram/README.md) and state how its decision version differs from grouping.

**Personal review:** Confidence `— / 5` · Last reviewed `—` · Next review `—`.
