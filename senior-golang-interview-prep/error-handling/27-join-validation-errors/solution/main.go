package main
import ("errors";"fmt")
type User struct{Name string;Age int}
type ValidationError struct{Field string}
func(e *ValidationError)Error()string{return e.Field+": invalid"}
func ValidateUser(u User)error{var errs []error;if u.Name==""{errs=append(errs,&ValidationError{"name"})};if u.Age<0{errs=append(errs,&ValidationError{"age"})};return errors.Join(errs...)}
func main(){fmt.Println(ValidateUser(User{}))}

