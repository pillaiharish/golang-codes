package main
import ("errors";"testing")
func TestWrapping(t *testing.T){if errors.Is(withV(),ErrRemote){t.Fatal("%v must not preserve identity")};if !errors.Is(withW(),ErrRemote){t.Fatal("%w should preserve identity")}}

