# #11 Container With Most Water — cold-recall card

[Full guide](README.md) · [Optimized Go 1.21+](main.go) · [Go 1.20 legacy](variants_pre_go121.go) · [Brute force](variants.go) · [Tests](main_test.go) · [LeetCode index](../README.md)

**Trigger:** Maximum area between two vertical lines; area = `min(leftHeight,rightHeight) * (right-left)`.

**Insight:** Width only decreases. Moving the taller line while keeping the limiting shorter line cannot create a larger area. Therefore move the shorter line.

**Invariant:** All endpoint pairs discarded by pointer moves cannot beat the best area already examined.

| Step | Reconstruct | Purpose |
|---|---|---|
| 1 | `left=0; right=n-1; best=0` | Start with maximum width |
| 2 | `area = min(height[left],height[right])*(right-left)` | Compute capacity |
| 3 | `best = max(best,area)` | Remember greatest found |
| 4 | Move the shorter side inward | Discard provably unhelpful pairs |
| 5 | Stop when `left >= right` | Need two distinct lines |

**Trace:** `[1,8,6,2,5,4,8,3,7]` → pair `(1,8)` has area `7*7=49`.

**Pitfalls:** `right-left+1` is wrong; `max` height is wrong; moving taller wall is wrong; no sorting. For Go 1.21+ use `h := min(height[left],height[right])`—do not shadow the input slice `height`.

**Version switch:** Primary `maxArea` now directly uses `max(area, min(height[left],height[right])*(right-left))` (Go 1.21+). The older `main.go` is preserved as `maxAreaLegacy` in [variants_pre_go121.go](variants_pre_go121.go). To submit on Go 1.20 or older, copy that function and rename it to `maxArea`; the full folder includes newer-Go code and will not compile with an older toolchain.

**Complexity:** O(n) time and O(1) extra space, versus brute force O(n²)/O(1).

**Cold drill:** Write two-pointer solution without opening main.go, then explain the discard inequality in one sentence. Test `[1,1]`, `[4,3,2,1,4]` and `[1,2,1]`.

Confidence `—` · Last cold recall `—` · Next recall `—`
