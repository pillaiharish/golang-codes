package main
import ("errors";"testing")
func TestIdentity(t *testing.T){if !isMissing(lookup()){t.Fatal("wrapped sentinel missed")}
 if isMissing(errors.New("not found")){t.Fatal("different error value matched by text")}
}

