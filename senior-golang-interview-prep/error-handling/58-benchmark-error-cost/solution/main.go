package main
import ("errors";"fmt")
var ErrMissing=errors.New("not found")
func sentinel()error{return ErrMissing}
func contextual(id string)error{return fmt.Errorf("user %q: %w",id,ErrMissing)}
func main(){fmt.Println(contextual("a"))}

