package main
import ("errors";"strings";"testing")
func TestContextError(t *testing.T){err:=&ContextError{"backup",ErrDisk};if !errors.Is(err,ErrDisk)||!strings.Contains(err.Error(),"backup"){t.Fatalf("%v",err)}}

