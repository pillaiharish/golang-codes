package main
import ("errors";"fmt")
var ErrRemote = errors.New("remote unavailable")
func withV() error {return fmt.Errorf("fetch: %v",ErrRemote)}
func withW() error {return fmt.Errorf("fetch: %w",ErrRemote)}
func main(){fmt.Println(errors.Is(withV(),ErrRemote),errors.Is(withW(),ErrRemote))}

