package main
import ("errors";"testing")
func TestAs(t *testing.T){var v *ValidationError;if !errors.As(validate(),&v)||v.Field!="email"{t.Fatalf("wrong validation error: %+v",v)}}

