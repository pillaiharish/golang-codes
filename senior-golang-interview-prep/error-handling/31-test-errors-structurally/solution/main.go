package main
import ("errors";"fmt")
var ErrMissing=errors.New("missing")
type LookupError struct{Key string;Err error}
func(e *LookupError)Error()string{return fmt.Sprintf("lookup %s: %v",e.Key,e.Err)}
func(e *LookupError)Unwrap()error{return e.Err}
func lookup(k string)error{return &LookupError{k,ErrMissing}}
func main(){fmt.Println(lookup("abc"))}

