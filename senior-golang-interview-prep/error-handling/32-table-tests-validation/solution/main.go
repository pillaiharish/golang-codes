package main
import ("errors";"fmt";"strings")
var ErrInvalid=errors.New("invalid email")
type ValidationError struct{Field,Reason string;Err error}
func(e *ValidationError)Error()string{return e.Field+": "+e.Reason}
func(e *ValidationError)Unwrap()error{return e.Err}
func ValidateEmail(s string)error{reason:="";switch{case s=="":reason="required";case len(s)>254:reason="too long";case !strings.Contains(s,"@"):reason="missing @";case !strings.Contains(strings.SplitN(s,"@",2)[1],"."):reason="invalid domain"}
 if reason!=""{return fmt.Errorf("validate email: %w",&ValidationError{"email",reason,ErrInvalid})};return nil}
func main(){fmt.Println(ValidateEmail("bad"))}

