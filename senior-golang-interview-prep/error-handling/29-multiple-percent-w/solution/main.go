package main
import ("errors";"fmt")
var ErrA=errors.New("A")
var ErrB=errors.New("B")
func both()error{return fmt.Errorf("two failures: %w and %w",ErrA,ErrB)}
func main(){fmt.Println(both())}

