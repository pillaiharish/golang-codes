package main
import "fmt"
type ValidationError struct{Field,Message string}
func(e *ValidationError)Error()string{return e.Field+": "+e.Message}
func ValidateAge(n int)error{if n<0{return &ValidationError{"age","must be non-negative"}};return nil}
func main(){fmt.Println(ValidateAge(-1))}

