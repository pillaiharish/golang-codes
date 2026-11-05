package main
import "testing"
func TestNilError(t *testing.T){var err error; if !isNil(err)||!isNil(successful()){t.Fatal("zero error should be nil")}}

