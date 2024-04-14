package main

import "fmt"

func isValid(s string) bool {

	stack := []rune{}
	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		opening, found := bracketMap[char]
		if found {
			fmt.Printf("%c\n", opening)
			// check if stack is empty or top of stack is not an opening bracket
			if len(stack) == 0 || stack[len(stack)-1] != opening {
				return false
			}
			// top of the stack is opening bracket
			// pop the stack
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}

func main() {
	// Test cases
	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("(]"))     // false
	fmt.Println(isValid("([)]"))   // false
	fmt.Println(isValid("{[]}"))   // true
}
