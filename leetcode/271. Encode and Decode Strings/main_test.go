package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestCodecRoundTrip(t *testing.T) {
	codec := Constructor()
	tests := []struct {
		name  string
		input []string
	}{
		{"common example", []string{"neet", "code", "love", "you"}},
		{"empty list", []string{}},
		{"single empty string", []string{""}},
		{"multiple empty strings", []string{"", "", ""}},
		{"delimiter inside data", []string{"a#b", "12#34", "#"}},
		{"digits and whitespace", []string{"123", "0", "hello world", "\n"}},
		{"unicode uses byte length", []string{"é", "😊", "ગુજરાતી"}},
		{"long multi digit length", []string{strings.Repeat("x", 120), "tail"}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			encoded := codec.Encode(tc.input)
			if got := codec.Decode(encoded); !reflect.DeepEqual(got, tc.input) {
				t.Fatalf("Decode(Encode(%q)) = %q; encoded %q", tc.input, got, encoded)
			}
		})
	}
}

func TestCodecFramingAndRepeatedCalls(t *testing.T) {
	codec := &Codec{} // The alternate user-provided construction works too.
	if got := codec.Encode([]string{"a#b", ""}); got != "3#a#b0#" {
		t.Fatalf("framing got %q, want %q", got, "3#a#b0#")
	}
	if got := codec.Encode([]string{"x"}); got != "1#x" {
		t.Fatalf("second Encode must not include previous call: %q", got)
	}
	if got := codec.Decode("0#"); !reflect.DeepEqual(got, []string{""}) {
		t.Fatalf("empty payload got %q", got)
	}
	if got := codec.Decode(""); !reflect.DeepEqual(got, []string{}) {
		t.Fatalf("empty input got %q", got)
	}
}
