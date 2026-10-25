package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Codec is stateless: each Encode/Decode call uses only its input.
type Codec struct{}

// Constructor follows the design-problem API and returns a Codec value.
// Codec{} is a value; &Codec{} would be a *Codec instead.
func Constructor() Codec {
	return Codec{}
}

// Encode frames each string as ASCII byte-length, '#', then its exact bytes.
// Example: []string{"a#b", ""} => "3#a#b0#".
func (c *Codec) Encode(strs []string) string {
	var builder strings.Builder
	for _, str := range strs {
		// Step 1: Write the byte count; len(string) counts UTF-8 bytes.
		builder.WriteString(strconv.Itoa(len(str)))
		// Step 2: Separate a variable-width count from the raw payload.
		builder.WriteByte('#')
		// Step 3: Write the entire string without escaping its contents.
		builder.WriteString(str)
	}
	return builder.String()
}

// Decode assumes s was created by Encode (the problem's valid-input contract).
// Byte positions intentionally match the byte counts written by Encode.
func (c *Codec) Decode(s string) []string {
	result := []string{} // Distinguishes [] from [""] for round trips.
	i := 0
	for i < len(s) {
		// Step 1: Find the FIRST '#' after the ASCII decimal length.
		j := i
		for s[j] != '#' {
			j++
		}
		// Step 2: Parse just the length prefix, not the payload.
		length, _ := strconv.Atoi(s[i:j]) // Valid Encode output ensures decimal digits.
		// Step 3: Slice exactly length bytes, even if the payload contains '#'.
		start := j + 1
		end := start + length // Go slice end is exclusive.
		result = append(result, s[start:end])
		// Step 4: Resume at the next record, not the next '#' in the payload.
		i = end
	}
	return result
}

func main() {
	codec := Constructor() // Same behavior as (&Codec{}): there is no stored state.
	input := []string{"neet", "code", "love", "you"}
	encoded := codec.Encode(input)
	decoded := codec.Decode(encoded)
	fmt.Println("Encoded:", encoded)
	fmt.Println("Decoded:", decoded)
}
