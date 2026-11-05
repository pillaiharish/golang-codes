package main
import ("errors";"testing")
func TestIs(t *testing.T){err:=get(2);if err==ErrNotFound||!errors.Is(err,ErrNotFound){t.Fatalf("chain broken: %v",err)}}

