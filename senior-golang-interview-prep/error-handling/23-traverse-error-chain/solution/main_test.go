package main
import ("errors";"testing")
func TestChain(t *testing.T){if !errors.Is(chain(false),base){t.Fatal("missing deep cause")};if errors.Is(chain(true),base){t.Fatal("broken chain matched")}}

