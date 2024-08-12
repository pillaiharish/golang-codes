# 1366. Rank Teams by Votes

## Approach

### Counting Votes
- Use a map (or a dictionary) to store each team's rank position across all votes.
- For each position (1st, 2nd, etc.), count how many times each team appears in that position.

### Sorting
- After counting, sort the teams based on their rankings across all positions.
- If two teams have the same rank in all positions, sort them alphabetically.

## Explanation

### Initialization
- We first determine the number of teams (length of each vote string).
- We create a rank map where the key is a team (byte) and the value is an integer array that will hold the count of votes at each rank position.

### Counting Votes
- For each vote, we update the rank map by increasing the count of each team at its corresponding rank position.

### Sorting
- We create a slice `teams` that holds all team identifiers.
- We then sort the `teams` slice based on the following criteria:
  - For each position, compare the number of votes each team has received. The team with more votes at a higher position is ranked higher.
  - If two teams have the same number of votes across all positions, we sort them lexicographically (alphabetically).

### Returning the Result
- Finally, we convert the sorted slice of teams back to a string and return it.

## Complexity
- **Time Complexity**: \(O(N \times M \log M)\), where \(N\) is the number of votes and \(M\) is the number of teams. The \(\log M\) factor comes from sorting.
- **Space Complexity**: \(O(M^2)\) for storing the rank counts.

This approach efficiently ranks the teams by their positions in the votes and resolves ties lexicographically, as required by the problem.