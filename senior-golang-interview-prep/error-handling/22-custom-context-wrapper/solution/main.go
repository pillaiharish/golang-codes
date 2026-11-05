package main
import ("errors";"fmt")
var ErrDisk=errors.New("disk failed")
type ContextError struct{Op string;Err error}
func(e *ContextError)Error()string{return e.Op+": "+e.Err.Error()}
func(e *ContextError)Unwrap()error{return e.Err}
func main(){fmt.Println(&ContextError{"backup",ErrDisk})}

