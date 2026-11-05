package main
import ("errors";"testing")
func TestPolicyWrapper(t *testing.T){if atBoundary("save",nil)!=nil{t.Fatal("nil changed")};if !errors.Is(atBoundary("save",ErrConflict),ErrConflict){t.Fatal("lost cause")}}

