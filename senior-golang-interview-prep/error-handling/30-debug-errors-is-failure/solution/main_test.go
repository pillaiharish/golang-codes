package main
import ("errors";"testing")
func TestRepair(t *testing.T){if errors.Is(broken(),ErrPermission){t.Fatal("broken case unexpectedly matches")};if !errors.Is(fixed(),ErrPermission){t.Fatal("regression: cause lost")}}

