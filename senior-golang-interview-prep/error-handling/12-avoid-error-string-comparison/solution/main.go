package main
import ("errors";"fmt")
var ErrMissing=errors.New("not found")
func lookup()error{return fmt.Errorf("lookup user: %w",ErrMissing)}
func isMissing(err error)bool{return errors.Is(err,ErrMissing)}
func main(){fmt.Println(isMissing(lookup()))}

