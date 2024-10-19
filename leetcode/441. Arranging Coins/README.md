# Problem: 441. Arranging Coins
The problem asks you to find how many complete rows of coins can be formed given n coins, where each row contains one more coin than the previous row (the first row contains 1 coin, the second contains 2, and so on).

## Approach:
This problem can be solved using mathematical reasoning. The sum of coins for k complete rows is given by the formula for the sum of the first k integers:


*sum := k*(k+1)/2*
 

We need to find the largest integer k such that the total number of coins required for k rows does not exceed n.

## Binary Search Solution:
We can use binary search to find the maximum number of complete rows that can be formed.

## Explanation:
### Binary Search: 
We perform binary search between 0 and n to find the number of complete rows that can be formed.


### Sum Formula: 
For each mid, the number of coins used to fill mid rows is calculated as mid * (mid + 1) / 2.


### Condition Check: 
If the coins used exactly equal n, we return mid. Otherwise, we adjust the search bounds.


### Termination: 
When the loop ends, right will contain the number of complete rows.


### Time Complexity:
O(log n): Due to the binary search.


### Space Complexity:
O(1): No extra space is used beyond a few variables.