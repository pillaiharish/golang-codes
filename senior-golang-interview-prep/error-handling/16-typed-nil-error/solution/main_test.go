package main
import "testing"
func TestTypedNil(t *testing.T){if broken()==nil{t.Fatal("expected typed nil interface to be non-nil")};if fixed()!=nil{t.Fatal("fixed must return nil")}}

