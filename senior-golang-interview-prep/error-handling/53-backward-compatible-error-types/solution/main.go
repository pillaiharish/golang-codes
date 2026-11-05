package main
import ("errors";"fmt")
var ErrNotFound=errors.New("not found")
type NotFoundError struct{Resource,ID string}
func(e *NotFoundError)Error()string{return fmt.Sprintf("%s %q not found",e.Resource,e.ID)}
func(e *NotFoundError)Is(target error)bool{return target==ErrNotFound}
func lookup()error{return &NotFoundError{"user","42"}}
func main(){fmt.Println(errors.Is(lookup(),ErrNotFound))}

