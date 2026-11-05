package main
import ("errors";"os";"testing")
func TestFileFailure(t *testing.T){_,err:=size("file-that-does-not-exist-987654");if !errors.Is(err,os.ErrNotExist){t.Fatalf("missing cause: %v",err)}}

