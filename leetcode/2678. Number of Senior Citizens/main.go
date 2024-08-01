package main

// You are given a 0-indexed array of strings details. Each element of details provides
// information about a given passenger compressed into a string of length 15. The system is such that:
// The first ten characters consist of the phone number of passengers.
// The next character denotes the gender of the person.
// The following two characters are used to indicate the age of the person.
// The last two characters determine the seat allotted to that person.
import (
	"fmt"
	"strconv"
)

func countSeniors(details []string) int {
	count := 0
	for i := range details {
		age, err := strconv.Atoi(details[i][11:13])
		if err != nil {
			return 0
		}
		if age > 60 {
			count++
		}
	}
	// fmt.Println(age)
	return count
}

func main() {
	details1 := []string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}
	details2 := []string{"1313579440F2036", "2921522980M5644"}
	out1 := countSeniors(details1)
	out2 := countSeniors(details2)
	fmt.Println("for details1 no. of elders are: ", out1)
	fmt.Println("for details2 no. of elders are: ", out2)
}

// Output:
// harish $ go run main.go
// for details1 no. of elders are:  2
// for details2 no. of elders are:  0
