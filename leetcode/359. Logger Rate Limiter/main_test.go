package main

import "testing"

// All sequences respect the problem's non-decreasing timestamp guarantee.
// The "AlgoMonster six-call example" reproduces the documented call inputs
// and expected booleans. All other cases are additional boundary tests.
func TestLoggerShouldPrintMessage(t *testing.T) {
	type call struct {
		timestamp int
		message   string
		want      bool
	}

	tests := []struct {
		name  string
		calls []call
	}{
		{
			name: "AlgoMonster six-call example",
			calls: []call{
				{1, "foo", true},
				{2, "bar", true},
				{3, "foo", false},
				{8, "bar", false},
				{11, "foo", true},
				{12, "bar", true},
			},
		},
		{
			name: "9 seconds blocked, 10 seconds allowed",
			calls: []call{
				{1, "foo", true},
				{10, "foo", false}, // 9 seconds after accepted timestamp 1
				{11, "foo", true},  // exactly 10 seconds after
				{20, "foo", false},
				{21, "foo", true},
			},
		},
		{
			name: "same timestamp repeated message",
			calls: []call{
				{5, "ping", true},
				{5, "ping", false},
				{5, "pong", true},
				{5, "pong", false},
				{15, "ping", true},
				{15, "pong", true},
			},
		},
		{
			name: "rejections must not extend cooldown",
			calls: []call{
				{1, "ping", true},
				{8, "ping", false},
				{9, "ping", false},
				{10, "ping", false},
				{11, "ping", true},
				{12, "ping", false},
				{21, "ping", true},
			},
		},
		{
			name: "different keys keep independent windows",
			calls: []call{
				{1, "a", true},
				{2, "b", true},
				{3, "a", false},
				{10, "a", false},
				{11, "a", true},
				{11, "b", false},
				{12, "b", true},
				{13, "c", true},
			},
		},
		{
			name: "long idle period allows reprinting",
			calls: []call{
				{1, "hello", true},
				{100, "hello", true},
				{101, "hello", false},
				{110, "hello", true},
			},
		},
		{
			name: "first occurrence at timestamp zero",
			calls: []call{
				{0, "startup", true},
				{0, "startup", false},
				{9, "startup", false},
				{10, "startup", true},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			logger := Constructor() // each case starts with independent state
			for i, c := range tc.calls {
				got := logger.ShouldPrintMessage(c.timestamp, c.message)
				if got != c.want {
					t.Errorf("call %d: ShouldPrintMessage(%d, %q) = %t, want %t",
						i+1, c.timestamp, c.message, got, c.want)
				}
			}
		})
	}
}

func TestLoggerInstancesHaveIndependentState(t *testing.T) {
	first := Constructor()
	second := Constructor()

	if !first.ShouldPrintMessage(1, "foo") {
		t.Fatal("first logger should accept its first message")
	}
	if first.ShouldPrintMessage(1, "foo") {
		t.Fatal("first logger should suppress an immediate duplicate")
	}
	if !second.ShouldPrintMessage(1, "foo") {
		t.Fatal("another logger must have its own independent state")
	}
}
