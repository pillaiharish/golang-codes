package main
import ("errors";"testing")
func TestLookup(t *testing.T){err:=lookup("abc");var e *LookupError
 if !errors.Is(err,ErrMissing)||!errors.As(err,&e)||e.Key!="abc"{t.Fatalf("lost contract: %v",err)}
}

