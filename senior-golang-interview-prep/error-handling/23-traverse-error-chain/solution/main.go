package main
import ("errors";"fmt")
var base=errors.New("base")
func chain(broken bool)error{var mid error;if broken{mid=fmt.Errorf("middle: %v",base)}else{mid=fmt.Errorf("middle: %w",base)};return fmt.Errorf("top: %w",mid)}
func main(){fmt.Println(errors.Is(chain(false),base))}

