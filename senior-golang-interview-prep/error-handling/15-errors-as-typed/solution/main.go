package main
import ("errors";"fmt")
type ValidationError struct{Field string}
func(e *ValidationError)Error()string{return "invalid "+e.Field}
func validate()error{return fmt.Errorf("create user: %w",&ValidationError{"email"})}
func main(){var e *ValidationError;fmt.Println(errors.As(validate(),&e),e)}

