package main
import ("errors";"testing")
func TestAdapter(t *testing.T){if !errors.Is(getUser(0),ErrNotFound){t.Fatal("wrong domain error")};var v *VendorError;if !errors.As(getUser(1),&v)||v.Code!=503{t.Fatalf("lost vendor metadata: %+v",v)}}

