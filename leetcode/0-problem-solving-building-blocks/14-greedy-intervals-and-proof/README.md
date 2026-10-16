# 14 — Greedy choices and interval sorting

[← Foundations index](../README.md) · [← Problem index](../../README.md)

## Recognize the building block
A greedy algorithm commits to locally chosen positions. The hard part is proving the choice keeps a globally optimal solution possible.

## Operations you must be able to write
- For #757, sort intervals by right endpoint ascending; break ties by left endpoint descending.
- Track the last two chosen points `p1<p2`. Each interval needs at least two selected points.
- If interval start >p2, add two points at `end-1,end`. If start>p1, add one point at `end`.
- Choosing points as far right as possible preserves usefulness for later intervals sorted by end.
- Use a counter and a proof of exchange: moving selected points right cannot hurt earlier satisfied intervals given this ordering.

## Complete small Go example
```go
package main
import ("fmt";"sort")
func minSelected(intervals [][]int)int{
    sort.Slice(intervals,func(i,j int)bool{
        if intervals[i][1]!=intervals[j][1]{return intervals[i][1]<intervals[j][1]}
        return intervals[i][0]>intervals[j][0]
    })
    count,p1,p2:=0,-1,-1
    for _,v:=range intervals{
        if v[0]>p2{count+=2;p1,p2=v[1]-1,v[1]
        }else if v[0]>p1{count++;p1,p2=p2,v[1]}
    }
    return count
}
func main(){fmt.Println(minSelected([][]int{{1,3},{1,4},{2,5},{3,5}}))}
```

## Dry run
After sorting, select points 2,3 for [1,3]; [1,4] already has two; for [3,5], only 3 is present, add 5; [2,5] has 2,3,5. Answer 3.

## Tiny practice task
Explain why adding `end-1,end` is better than choosing the leftmost two points. Show a tie where start descending matters.

## Common mistakes
- Sort mutates the intervals slice; copy before sorting if the caller expects original order.
- A greedy implementation is not correct merely because it passes examples—state the exchange argument.
- LeetCode #757 assumes interval length at least 2; `end-1` must belong to the interval.

## Problem links from this repository
- [#757 Set Intersection Size At Least Two](../../757.%20Set%20Intersection%20Size%20At%20Least%20Two/main.go)
- [#2273 Find Resultant Array After Removing Anagrams](../../2273.%20Find%20Resultant%20Array%20After%20Removing%20Anagrams/main.go)

## Blank-screen check
Close this note. Restate the invariant, retype the operations, then solve the first linked problem without looking.
