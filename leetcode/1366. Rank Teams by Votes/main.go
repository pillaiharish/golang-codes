package main

import (
	"fmt"
	"sort"
)

func rankTeams(votes []string) string {
	n := len(votes[0])           // number of teams
	rank := make(map[byte][]int) //make a map
	//initialize map with each letter and assign empty arry to it
	for i := 0; i < n; i++ {
		rank[votes[0][i]] = make([]int, n)
	} // this will give rank length 3 keys

	// count the votes for each team in each rank position
	for _, v := range votes {
		for i := 0; i < n; i++ {
			rank[v[i]][i]++
		}
	}
	fmt.Println(rank) //map[65:[5 0 0] 66:[0 2 3] 67:[0 3 2]]
	teams := []byte(votes[0])

	// sort the teams based on the rank counts, and lexicographically if tied
	sort.Slice(teams, func(i, j int) bool {
		// select two value and loop through its innermost slice to sort them
		team1, team2 := teams[i], teams[j]
		for k := 0; k < n; k++ {
			if rank[team1][k] > rank[team2][k] {
				return true
			} else if rank[team1][k] < rank[team2][k] {
				return false
			}
		}
		// after inner slices are sorted now sort the teams itself
		return team1 < team2
	})

	return string(teams)
}

func main() {
	votes := []string{"ABC", "ACB", "ABC", "ACB", "ACB"}
	fmt.Println(rankTeams(votes)) // Output: "ACB"
}

// Output:
// harish $ go run main.go
// map[65:[5 0 0] 66:[0 2 3] 67:[0 3 2]]
// ACB
