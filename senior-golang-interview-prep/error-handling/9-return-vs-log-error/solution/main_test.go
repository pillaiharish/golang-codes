package main
import ("errors";"testing")
func TestLoadPreservesCause(t *testing.T){if !errors.Is(load(),ErrStorage){t.Fatal("lost storage failure")}}

