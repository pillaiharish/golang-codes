package main
import ("errors";"testing")
func TestJoinNil(t *testing.T){if collected(nil,nil)!=nil{t.Fatal("all nil must be nil")};if !errors.Is(collected(nil,ErrOne),ErrOne){t.Fatal("non-nil child lost")}}

