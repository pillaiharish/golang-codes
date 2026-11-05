package main
import "testing"
func TestCustomError(t *testing.T) {
 var _ error = (*myError)(nil)
 err := fail()
 if err == nil || err.Error() != "example failed" { t.Fatalf("unexpected error: %v", err) }
}

