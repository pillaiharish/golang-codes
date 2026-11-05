package main
import ("errors";"testing")
func TestCollect(t *testing.T){err:=collect();if !errors.Is(err,ErrA)||!errors.Is(err,ErrB){t.Fatalf("lost concurrent failures: %v",err)}}

