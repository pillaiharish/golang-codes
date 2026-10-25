package main

import "fmt"

type Logger struct {
	// Persist state across method calls; a local map would be recreated on
	// every call. The key is a message; the value is its last accepted time.
	logs map[string]int
}

func Constructor() Logger {
	// Step 1: Allocate a writable map for this logger instance.
	// Why: the zero value of a map is nil; assignment into a nil map panics.
	return Logger{logs: make(map[string]int)}
}

func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	// Step 2: Check this message's prior accepted time, if any.
	if previousTimestamp, exists := this.logs[message]; exists {
		// Step 3: Reject while fewer than 10 seconds have elapsed.
		// Why: a rejected arrival must NOT restart the cooldown.
		if timestamp-previousTimestamp < 10 {
			return false
		}
	}
	// Step 4: First occurrence or expired cooldown: accept and persist time.
	this.logs[message] = timestamp
	// Step 5: Report eligibility; printing is the caller's responsibility.
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
