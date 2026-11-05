package main
import ("errors";"fmt")
var ErrDenied=errors.New("permission denied")
func repository()error{return ErrDenied}
func service()error{return repository()} // no new context to add
func handler()error{return fmt.Errorf("save profile: %w",service())}
func main(){fmt.Println(handler())}

