package main

import "fmt"

type Logger struct {
	logs map[string]int
}

func Constructor() Logger {
	return Logger{logs: make(map[string]int)}
}

func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if previousTimestamp, exists := this.logs[message]; exists {
		if timestamp-previousTimestamp < 10 {
			return false
		}
	}
	this.logs[message] = timestamp
	return true
}

/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */

func main() {
	logger := Constructor()

	// Test cases
	fmt.Println(logger.ShouldPrintMessage(1, "foo"))  // true
	fmt.Println(logger.ShouldPrintMessage(2, "bar"))  // true
	fmt.Println(logger.ShouldPrintMessage(3, "foo"))  // false
	fmt.Println(logger.ShouldPrintMessage(11, "foo")) // true
}

// Output:
// $ go run main.go
// true
// true
// false
// true
