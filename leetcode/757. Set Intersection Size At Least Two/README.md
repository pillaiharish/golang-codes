# Minimum Number of Arrows to Burst Balloons

## Approach
The greedy approach to solve this problem is as follows:

1. **Sort the Intervals**:
   - Sort intervals based on their end points. If two intervals have the same end, sort by the start point in descending order.

2. **Greedy Selection**:
   - Use a greedy approach by iterating over the sorted intervals and greedily choosing two points (let's say `p1` and `p2`) to be included in the solution set `S`.
   - Ensure that these two points cover as many intervals as possible without violating the condition that every interval must have at least two points from `S`.

## Explanation

### Sorting
The intervals are sorted by their end points. If two intervals have the same end, they are sorted by their start points in descending order to maximize overlap.

### Greedy Choice
The main idea is to ensure that the chosen points (`p1` and `p2`) will cover the current interval while also being as close as possible to the end of the interval to cover subsequent intervals.

- If neither `p1` nor `p2` falls within the current interval, we need to add two new points from the current interval.
- If only one of `p1` or `p2` falls within the current interval, we need to add one more point.

## Complexity
- The time complexity of this solution is \(O(n \log n)\) due to the sorting step, where \(n\) is the number of intervals.
- The greedy process itself runs in \(O(n)\) time.